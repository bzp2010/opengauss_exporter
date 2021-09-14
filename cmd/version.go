package cmd

import (
	"fmt"

	"github.com/prometheus/common/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information of Exporter",
	Run: func(cmd *cobra.Command, args []string) {
		printVersionInfo()
	},
}

func printVersionInfo()  {
	fmt.Printf("openGauss Exporter\nVersion: %s\nBuild Time: %s\nGitHash: %s", version.Version, version.BuildDate, version.Branch)
}