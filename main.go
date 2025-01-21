package main

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/mattn/go-sqlite3"

	"github.com/dusto/sigils/pkg/repository"
	"github.com/dusto/sigils/pkg/route"
)

//go:embed schema.sql
var backendDDL string

type Options struct {
	StorePath string `help:"Path to database file" default:"sigils.db"`
	Port      int    `help:"Port to listen on " default:"8888"`
}

func main() {

	cli := humacli.New(func(hooks humacli.Hooks, opts *Options) {
		ctx := context.Background()

		dbUri := fmt.Sprintf("file:%s?_foreign_keys", opts.StorePath)
		db, err := sql.Open("sqlite3", dbUri)
		if err != nil {
			panic(err)
		}

		if _, err := db.ExecContext(ctx, backendDDL); err != nil {
			panic(err)
		}

		queries := repository.New(db)
		router := chi.NewRouter()
		router.Use(middleware.RequestID)
		router.Use(middleware.RealIP)
		router.Use(middleware.Logger)
		router.Use(middleware.Recoverer)

		api := humachi.New(router, huma.DefaultConfig("Sigils", "0.0.1"))

		handle := route.NewHandler(api, queries)
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
			Addr:    fmt.Sprintf("%s:%d", "localhost", opts.Port),
			Handler: router,
		}

		hooks.OnStart(func() {
			log.Printf("Server is running with: host:%v port:%v\n", "localhost", opts.Port)

			err := srv.ListenAndServe()
			if err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		})

	})

	cli.Run()
}
