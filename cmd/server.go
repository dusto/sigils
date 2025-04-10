package cmd

import (
	"context"
	_ "embed"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/dusto/sigils/internal/repository"
	"github.com/dusto/sigils/internal/route"
)

type ServerOptions struct {
	StorePath    string `help:"Path to database file" default:"."`
	DatabaseFile string `help:"Filename for database" default:"sigils.db"`
	Port         int    `help:"Port to listen on " default:"8888"`
	MetricsPort  int    `help:"Port to serve Prometheus metrics" default:"9001"`
	AutoAdd      bool   `help:"Enable/Disable Auto adding hosts that tried to get a machineconfig" default:"true"`
	OpenAPI      bool   `help:"Print openapi spec to stdout" default:"false"`
}

func (c *Cmd) server() *cobra.Command {

	var opts = &ServerOptions{}

	var server = &cobra.Command{
		Use:   "server",
		Short: "Run sigils server",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()

			dbPath := filepath.Join(opts.StorePath, opts.DatabaseFile)
			db := &repository.MultiSqliteDB{}
			err := db.SetupMultiSqliteDB(dbPath, repository.DefaultConnectionParams())
			if err != nil {
				c.httplogger.Error("Could not open database %s")
				panic(err)
			}
			c.promreg.MustRegister(
				db.CollectorReadDB(),
				db.CollectorWriteDB(),
			)

			if _, err := db.ExecContext(ctx, c.backendDDL); err != nil {
				panic(err)
			}

			queries := repository.New(db)
			router := chi.NewRouter()
			router.Use(middleware.RequestID)
			router.Use(middleware.RealIP)
			router.Use(httplog.RequestLogger(c.httplogger))
			router.Use(middleware.Recoverer)

			c.api = humachi.New(router, huma.DefaultConfig("Sigils", "0.0.1"))

			handleOpts := &route.HandlerOpts{
				AutoAdd: opts.AutoAdd,
			}

			handle := route.NewHandler(c.api, db, queries, c.httplogger, handleOpts)
			handle.Register()

			if opts.OpenAPI {
				b, _ := yaml.Marshal(c.api.OpenAPI())
				fmt.Println(string(b))
				return
			}

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
				Handler: promhttp.HandlerFor(c.promreg, promhttp.HandlerOpts{}),
			}

			done := make(chan struct{}, 1)
			go func() {

				go func() {
					c.httplogger.Info("Server is running", "port", opts.Port)
					err := srv.ListenAndServe()
					if err != nil && err != http.ErrServerClosed {
						c.httplogger.Error("Failed to start api", err)
						os.Exit(1)
					}
					done <- struct{}{}
				}()

				go func() {
					c.httplogger.Info("Metrics is running", "port", opts.MetricsPort)
					err := metricssrv.ListenAndServe()
					if err != nil && err != http.ErrServerClosed {
						c.httplogger.Error("Failed to start metrics", err)
						os.Exit(1)
					}
					done <- struct{}{}
				}()
			}()

			sigC := make(chan os.Signal, 1)
			signal.Notify(sigC, syscall.SIGTERM, syscall.SIGINT)

			select {
			case <-done:
				// done quit
			case <-sigC:
				fmt.Println(os.Stderr, "Gracefully shutting down server...")
				go func() {
					go func() {
						ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
						defer cancel()
						srv.Shutdown(ctx)
					}()
					go func() {
						ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
						defer cancel()
						metricssrv.Shutdown(ctx)
					}()

				}()
			}
		},
	}

	server.Flags().StringVarP(&opts.StorePath, "storepath", "s", ".", "Path to database files")
	server.Flags().StringVarP(&opts.DatabaseFile, "dbname", "d", "sigils.db", "Filename for database")
	server.Flags().IntVarP(&opts.Port, "port", "p", 8888, "Port for Rest API to listen on")
	server.Flags().IntVarP(&opts.MetricsPort, "metrics_port", "m", 9001, "Port for Metrics endpoint")
	server.Flags().BoolVarP(&opts.AutoAdd, "auto_add", "a", true, "Enable/Disable auto adding hosts that tried to get a machine config")
	server.Flags().BoolVarP(&opts.OpenAPI, "openapi", "o", false, "Print OpenAPI docs to stdout")
	return server
}
