package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "App version",
	Long:  "Application version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1.0.0")
	},
}
