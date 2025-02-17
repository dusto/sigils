package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/dusto/sigils/sdk"
	"github.com/spf13/cobra"
)

type profileOpts struct {
	Name    string      `yaml:"name"`
	Patches []patchOpts `yaml:"patches"`
}

type patchOpts struct {
	Fqdn     string `yaml:"fqdn,omitempty"`
	NodeType string `yaml:"nodetype,omitempty"`
	Patch    string `yaml:"patch"`
}

func (c *Cmd) profile() *cobra.Command {

	var opts = &CliOptions{}
	var popts = &profileOpts{}
	var profile = &cobra.Command{
		Use:   "profile",
		Short: "Profile API commands",
	}
	var inputFile string
	var proId int64
	profile.PersistentFlags().StringVarP(&opts.Url, "url", "u", "http://localhost:8888", "URL for sigils api")

	addProfile := &cobra.Command{
		Use:   "add",
		Short: "Add new profile",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			client := getClient(ctx, opts.Url)
			if inputFile == "" {
				fmt.Println("Must pass yaml file")
				os.Exit(1)
			}
			parseYamlFile(inputFile, popts)
			in := sdk.Profile{
				Name:    &popts.Name,
				Patches: &[]sdk.Patch{},
			}
			for _, patch := range popts.Patches {
				p := sdk.Patch{}
				if patch.Patch == "" {
					// If no patch do nothing
					continue
				} else {
					p.Patch = &patch.Patch
				}
				if patch.Fqdn != "" {
					p.Fqdn = &patch.Fqdn
				}
				if patch.NodeType != "" {
					var nodetype sdk.PatchNodetype
					switch patch.NodeType {
					case "all":
						nodetype = sdk.PatchNodetypeAll
					case "controlplane":
						nodetype = sdk.PatchNodetypeControlplane
					case "worker":
						nodetype = sdk.PatchNodetypeWorker
					default:
						fmt.Printf("Invalid nodetype %s", patch.NodeType)
						os.Exit(1)
					}
					p.Nodetype = &nodetype
				}
				*in.Patches = append(*in.Patches, p)
			}
			resp, err := client.PostProfiles(ctx, []sdk.Profile{in})
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(resp.StatusCode)
		},
	}
	addProfile.Flags().StringVarP(&inputFile, "file", "f", "", "Profile Yaml")

	profile.AddCommand(addProfile)

	deleteProfile := &cobra.Command{
		Use:   "delete",
		Short: "Delete an existing profile",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			client := getClient(ctx, opts.Url)
			if proId == 0 {
				fmt.Println("Must pass a valid ID")
				os.Exit(1)
			}
			params := sdk.DeleteProfileParams{}
			resp, err := client.DeleteProfile(ctx, proId, &params)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(resp.StatusCode)
		},
	}
	deleteProfile.Flags().Int64VarP(&proId, "id", "i", 0, "Profile ID to update")
	profile.AddCommand(deleteProfile)

	return profile
}
