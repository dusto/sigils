package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/dusto/sigils/sdk"
	"github.com/google/uuid"
	"gopkg.in/yaml.v3"
)

func getClient(ctx context.Context, url string) *sdk.ClientWithResponses {
	client, err := sdk.NewClientWithResponses(url)
	if err != nil {
		fmt.Printf("failed to connect", err)
		os.Exit(1)
	}

	return client
}

func parseYamlFile[o any](file string, opts *o) {
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("File %s does not exist", file)
		os.Exit(1)
	}
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Could not read %s", file)
		os.Exit(1)
	}
	err = yaml.Unmarshal(data, opts)
	if err != nil {
		fmt.Printf("Could not unmarshal %s:\n %w", file, err)
		os.Exit(1)
	}
}

func validUUID(m string) bool {
	_, err := uuid.Parse(m)
	if err != nil {
		fmt.Println("Could not parse uuid", err)
		return false
	}
	return true
}
