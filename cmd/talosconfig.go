package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/dusto/sigils/sdk"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func (c *Cmd) talosconfig() *cobra.Command {
	var url string
	var clusterUUID string
	var talosconfig = &cobra.Command{
		Use:   "talosconfig",
		Short: "Get the talosconfig for specified cluster",
		Run: func(cmd *cobra.Command, args []string) {
			if clusterUUID == "" {
				fmt.Println("Must pass an UUID")
				os.Exit(1)
			}
			if _, err := uuid.Parse(clusterUUID); err != nil {
				fmt.Println("Could not parse uuid:", err)
				os.Exit(1)
			}

			ctx := context.Background()
			client := getClient(ctx, url)
			resp, err := client.GetClusterWithResponse(ctx, clusterUUID)
			if err != nil {
				fmt.Println("Could not get cluster:", err)
				os.Exit(1)
			}

			for _, cluster := range *resp.JSON200 {
				for _, config := range *cluster.Configs {
					if config.Configtype == sdk.ClusterConfigConfigtypeTalosctl {
						fmt.Println(config.Config)
					}
				}
			}

		},
	}
	talosconfig.PersistentFlags().StringVarP(&clusterUUID, "uuid", "i", "", "Cluster UUID")
	talosconfig.PersistentFlags().StringVarP(&url, "url", "u", "http://localhost:8888", "URL for sigils api")

	return talosconfig
}
