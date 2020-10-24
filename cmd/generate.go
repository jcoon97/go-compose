package cmd

import (
	"github.com/spf13/cobra"
)

var outputPath string

var generateCmd = &cobra.Command{
	Use:              "generate <mysql|mongo> [OPTIONS]...",
	Aliases:          []string{"gen"},
	Short:            "Generate a Docker Compose file for a specific database system",
	TraverseChildren: true,
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.PersistentFlags().StringVarP(&outputPath, "output", "o", "", "Specifies the output path for the docker-compose.yml file (default: Desktop)")
}
