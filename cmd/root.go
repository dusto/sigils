package cmd

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/httplog/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
)

var (
	URI string

	cliRoot = &cobra.Command{
		Use:   "sigils",
		Short: "Sigils cli",
		Long:  "Commands for running/interacting with sigils",
	}
)

type Cmd struct {
	httplogger *httplog.Logger
	promreg    *prometheus.Registry
	Cli        *cobra.Command
	humacli    humacli.CLI
	api        huma.API
	backendDDL string
}

func NewCmd(logger *httplog.Logger, registry *prometheus.Registry, ddl string) *Cmd {
	c := &Cmd{
		httplogger: logger,
		promreg:    registry,
		Cli:        cliRoot,
		backendDDL: ddl,
	}

	c.Cli.AddCommand(c.server())
	c.Cli.AddCommand(c.host())
	c.Cli.AddCommand(c.profile())
	c.Cli.AddCommand(c.talosconfig())

	return c
}
