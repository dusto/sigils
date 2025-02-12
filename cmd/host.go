package cmd

import "github.com/spf13/cobra"

func (c *Cmd) host() *cobra.Command {

	var opts = &CliOptions{}
	var host = &cobra.Command{
		Use:   "host",
		Short: "Host API commands",
	}
	host.PersistentFlags().StringVarP(&opts.Url, "url", "u", "http://localhost:8888", "URL for sigils api")

	host.AddCommand(&cobra.Command{
		Use:   "add",
		Short: "Add new host",
		Run: func(cmd *cobra.Command, args []string) {
		},
	})

	host.AddCommand(&cobra.Command{
		Use:   "update",
		Short: "Update an existing Host or add new host",
		Run: func(cmd *cobra.Command, args []string) {
		},
	})

	return host
}
