package cmd

import (
	"context"
	"fmt"
	"os"
	"slices"

	"github.com/dusto/sigils/sdk"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

type hostOpts struct {
	Fqdn       string  `yaml:"fqdn"`
	Uuid       string  `yaml:"uuid"`
	Mac        string  `yaml:"mac,omitempty"`
	Nodetype   string  `yaml:"nodetype"`
	ProfileIds []int64 `yaml:"profileids,omitempty"`
}

type listOpts struct {
	Search string
}

func commonHostFlags(cmd *cobra.Command, hopts *hostOpts, inputfile *string) {
	defaultIds := make([]int64, 0)
	cmd.Flags().StringVarP(&hopts.Uuid, "uuid", "i", "", "Host UUID")
	cmd.Flags().StringVarP(&hopts.Fqdn, "fqdn", "f", "", "Host FQDN")
	cmd.Flags().StringVarP(&hopts.Nodetype, "type", "t", "", "Host NodeType (controlplane,worker)")
	cmd.Flags().StringVarP(&hopts.Mac, "mac", "m", "", "Host Mac Address")
	cmd.Flags().StringVar(inputfile, "file", "", "Yaml file to use as input instead of specifying arguments on cli")
	cmd.Flags().Int64SliceVarP(&hopts.ProfileIds, "profiles", "p", defaultIds, "Profile Ids to associate to Host")
}

func (c *Cmd) host() *cobra.Command {

	var inputFile string
	var hostUuid string
	var opts = &CliOptions{}
	var hopts = &hostOpts{}
	var lopts = &listOpts{}

	var host = &cobra.Command{
		Use:   "host",
		Short: "Host API commands",
	}
	host.PersistentFlags().StringVarP(&opts.Url, "url", "u", "http://localhost:8888", "URL for sigils api")

	getHosts := &cobra.Command{
		Use:   "get",
		Short: "Get host(s)",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			client := getClient(ctx, opts.Url)

			params := sdk.ListHostsParams{
				Search: &hopts.Uuid,
			}
			res, err := client.ListHosts(ctx, &params)
			if err != nil {
				fmt.Println(err)
			}
			rep, err := sdk.ParseListHostsResponse(res)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%s", string(rep.Body))

		},
	}
	getHosts.Flags().StringVarP(&lopts.Search, "search", "s", "", "Search for specific UUID")
	host.AddCommand(getHosts)

	delHosts := &cobra.Command{
		Use:   "delete",
		Short: "Delete host",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			client := getClient(ctx, opts.Url)

			if !validUUID(hostUuid) {
				fmt.Printf("UUID: %s is not valid", hostUuid)
			}
			_, err := client.DeleteHost(ctx, uuid.MustParse(hostUuid))
			if err != nil {
				fmt.Println("Could not remove existing Host", err)
			}
			fmt.Printf("Host %s deleted", hostUuid)

		},
	}
	delHosts.Flags().StringVarP(&hostUuid, "uuid", "i", "", "UUID of Host")
	host.AddCommand(delHosts)

	addHost := &cobra.Command{
		Use:   "add",
		Short: "Add new host",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			client := getClient(ctx, opts.Url)

			if inputFile != "" {
				// Try to read file
				parseYamlFile(inputFile, hopts)
			}
			if !validUUID(hopts.Uuid) {
				os.Exit(1)
			}
			in := sdk.HostInput{
				Fqdn:       hopts.Fqdn,
				Uuid:       uuid.MustParse(hopts.Uuid),
				Nodetype:   sdk.HostInputNodetype(hopts.Nodetype),
				Mac:        &hopts.Mac,
				Profileids: &hopts.ProfileIds,
			}

			res, err := client.PostHosts(ctx, []sdk.HostInput{in})
			if err != nil {
				fmt.Println(err)
			}

			if res.StatusCode != 201 {
				fmt.Println("Could not add host")
				os.Exit(1)

			}
			fmt.Println("Added host:", res.StatusCode)
		},
	}
	commonHostFlags(addHost, hopts, &inputFile)
	host.AddCommand(addHost)

	updateHost := &cobra.Command{
		Use:   "update",
		Short: "Update an existing Host or add new host",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			client := getClient(ctx, opts.Url)

			if inputFile != "" {
				// Try to read file
				parseYamlFile(inputFile, hopts)
			}

			if !validUUID(hopts.Uuid) {
				os.Exit(1)
			}
			rep, err := client.GetHost(ctx, uuid.MustParse(hopts.Uuid))
			if err != nil {
				fmt.Printf("Failed to get host", err)
				os.Exit(1)
			}

			in := sdk.HostInput{}

			if rep.StatusCode == 200 {
				hostRep, err := sdk.ParseGetHostResponse(rep)
				if err != nil {
					fmt.Println("Could not parse host response", err)
					os.Exit(1)
				}

				// Only should get one
				hostS := (*hostRep.JSON200)[0]
				in.Uuid = hostS.Uuid
				in.Fqdn = hostS.Fqdn
				in.Mac = hostS.Mac
				in.Nodetype = sdk.HostInputNodetype(hostS.Nodetype)
				if hostS.Profiles != nil {
					for _, profile := range *hostS.Profiles {
						if profile.Id != nil {
							*in.Profileids = append(*in.Profileids, *profile.Id)
						}
					}
				}
				_, err = client.DeleteHost(ctx, in.Uuid)
				if err != nil {
					fmt.Println("Could not remove existing Host", err)
				}
			}

			//Override values from input
			in.Uuid = uuid.MustParse(hopts.Uuid)
			if hopts.Fqdn != "" {
				in.Fqdn = hopts.Fqdn
			}
			if hopts.Nodetype != "" {
				in.Nodetype = sdk.HostInputNodetype(hopts.Nodetype)

			}
			if hopts.Mac != "" {
				in.Mac = &hopts.Mac
			}
			if in.Profileids != nil || len(hopts.ProfileIds) > 0 {
				*in.Profileids = append(*in.Profileids, hopts.ProfileIds...)
				*in.Profileids = slices.Compact(*in.Profileids)
			}

			// If host override changes from opts
			res, err := client.PostHosts(ctx, []sdk.HostInput{in})
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Updated/Added host: Status: ", res.StatusCode)
		},
	}
	commonHostFlags(updateHost, hopts, &inputFile)
	host.AddCommand(updateHost)

	attachProfile := &cobra.Command{
		Use:   "attach-profile",
		Short: "Attach a profile(s) to a Host",
		Run: func(cmd *cobra.Command, args []string) {
			//ctx := context.Background()
			//client := getClient(ctx, opts.Url)

		},
	}
	host.AddCommand(attachProfile)

	attachCluster := &cobra.Command{
		Use:   "attach-cluster",
		Short: "Attach a Host to a Cluster",
		Run: func(cmd *cobra.Command, args []string) {
			//ctx := context.Background()
			//client := getClient(ctx, opts.Url)

		},
	}

	host.AddCommand(attachCluster)

	return host
}
