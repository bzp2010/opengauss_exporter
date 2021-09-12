package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"opengauss_exporter/internal/utils"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information of Exporter",
	Run: func(cmd *cobra.Command, args []string) {
		printVersionInfo()
	},
}

func printVersionInfo()  {
	fmt.Printf("openGauss Exporter\nVersion: %s\nBuild Time: %s\nGitHash: %s", utils.Version, utils.BuildTime, utils.GitHash)
}