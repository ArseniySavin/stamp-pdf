package commands

import "github.com/spf13/cobra"

var (
	inFile        string
	outFile       string
	stampFile     string
	page          string
	stampPosition string

	stamp         bool
	stampText     string
	stampExecutor string

	rootCmd = &cobra.Command{
		Use:   "stamp-pdf",
		Short: "Setting stamp as watermark into pdf",
		Long:  `The application is setting watermark stamp using the png file`,
	}
)
