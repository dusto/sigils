package cmd

type CliOptions struct {
	Url string `help:"URL for sigils instance" default:"http://localhost:8888"`
}
