package main

import (
	"context"
	_ "embed"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/dusto/sigils/internal/repository"
	"github.com/dusto/sigils/internal/route"
)

//go:embed schema.sql
var backendDDL string

type Options struct {
	StorePath    string `help:"Path to database file" default:"."`
	DatabaseFile string `help:"Filename for database" default:"sigils.db"`
	Port         int    `help:"Port to listen on " default:"8888"`
	MetricsPort  int    `help:"Port to serve Prometheus metrics" default:"9001"`
	AutoAdd      bool   `help:"Enable/Disable Auto adding hosts that tried to get a machineconfig" default:"true"`
}

func main() {
	registry := prometheus.NewRegistry()
	registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{
			Namespace:    "sigils",
			ReportErrors: true,
		}))

	logger := httplog.NewLogger("sigils", httplog.Options{
		JSON:           true,
		LogLevel:       slog.LevelDebug,
		RequestHeaders: true,
		Tags: map[string]string{
			"version": "v1.0.3",
			"env":     "prod",
		},
	})

	cli := humacli.New(func(hooks humacli.Hooks, opts *Options) {
		ctx := context.Background()

		dbPath := filepath.Join(opts.StorePath, opts.DatabaseFile)
		db := &repository.MultiSqliteDB{}
		err := db.SetupMultiSqliteDB(dbPath, repository.DefaultConnectionParams())
		if err != nil {
			logger.Error("Could not open database %s")
			panic(err)
		}
		registry.MustRegister(
			db.CollectorReadDB(),
			db.CollectorWriteDB(),
		)

		if _, err := db.ExecContext(ctx, backendDDL); err != nil {
			panic(err)
		}

		queries := repository.New(db)
		router := chi.NewRouter()
		router.Use(middleware.RequestID)
		router.Use(middleware.RealIP)
		router.Use(httplog.RequestLogger(logger))
		router.Use(middleware.Recoverer)

		api := humachi.New(router, huma.DefaultConfig("Sigils", "0.0.1"))

		handleOpts := &route.HandlerOpts{
			AutoAdd: opts.AutoAdd,
		}

		handle := route.NewHandler(api, db, queries, logger, handleOpts)
		handle.Register()

		// One off define style for docs
		router.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`<!doctype html>
        <html>
          <head>
            <title>API Reference</title>
            <meta charset="utf-8" />
            <meta
              name="viewport"
              content="width=device-width, initial-scale=1" />
          </head>
          <body>
            <script
              id="api-reference"
              data-url="/openapi.json"></script>
            <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
          </body>
        </html>`))
		})

		srv := &http.Server{
			Addr:    fmt.Sprintf(":%d", opts.Port),
			Handler: router,
		}

		metricssrv := &http.Server{
			Addr:    fmt.Sprintf(":%d", opts.MetricsPort),
			Handler: promhttp.HandlerFor(registry, promhttp.HandlerOpts{}),
		}

		hooks.OnStart(func() {

			go func() {
				logger.Info("Server is running", "port", opts.Port)
				err := srv.ListenAndServe()
				if err != nil && err != http.ErrServerClosed {
					logger.Error("Failed to start api", err)
					os.Exit(1)
				}
			}()

			go func() {
				logger.Info("Metrics is running", "port", opts.MetricsPort)
				err := metricssrv.ListenAndServe()
				if err != nil && err != http.ErrServerClosed {
					logger.Error("Failed to start metrics", err)
					os.Exit(1)
				}

			}()
			sigC := make(chan os.Signal, 1)
			signal.Notify(sigC, syscall.SIGTERM, syscall.SIGINT)
			<-sigC
		})

		hooks.OnStop(func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			srv.Shutdown(ctx)
		})

	})

	cli.Run()
}
