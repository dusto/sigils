package main

import (
	_ "embed"
	"log/slog"

	"github.com/go-chi/httplog/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"

	"github.com/dusto/sigils/cmd"
	"github.com/dusto/sigils/internal/version"
)

//go:embed schema.sql
var backendDDL string

func main() {
	registry := prometheus.NewRegistry()
	prometheus.NewRegistry()
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
			"version": version.Version,
			"env":     "prod",
		},
	})

	cli := cmd.NewCmd(logger, registry, backendDDL)
	cli.Cli.Execute()
}
