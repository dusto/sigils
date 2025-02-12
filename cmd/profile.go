package cmd

import "github.com/spf13/cobra"

func (c *Cmd) profile() *cobra.Command {

	var opts = &CliOptions{}
	var host = &cobra.Command{
		Use:   "profile",
		Short: "Profile API commands",
	}
	host.PersistentFlags().StringVarP(&opts.Url, "url", "u", "http://localhost:8888", "URL for sigils api")

	host.AddCommand(&cobra.Command{
		Use:   "add",
		Short: "Add new profile",
		Run: func(cmd *cobra.Command, args []string) {
		},
	})

	host.AddCommand(&cobra.Command{
		Use:   "update",
		Short: "Update an existing profile or add new profile",
		Run: func(cmd *cobra.Command, args []string) {
		},
	})

	return host
}
