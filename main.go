package main

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humaecho"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/labstack/echo/v4"

	_ "github.com/mattn/go-sqlite3"

	"github.com/dusto/sigils/pkg/handler"
)

//go:embed schema.sql
var backendDDL string

type Options struct {
	Port int `help:"Port to listen on " default:"8888"`
}

func main() {

	ctx := context.Background()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	if _, err := db.ExecContext(ctx, backendDDL); err != nil {
		panic(err)
	}

	cli := humacli.New(func(hooks humacli.Hooks, opts *Options) {
		router := echo.New()

		api := humaecho.New(router, huma.DefaultConfig("Test API", "0.0.1"))

		handler.RegisterApi(api)

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
