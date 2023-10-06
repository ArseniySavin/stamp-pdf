package commands

import (
	"fmt"

	"stamp-pdf/pkg"

	"github.com/spf13/cobra"
)

var stampInitCmd = &cobra.Command{
	Use:   "stamp-init",
	Short: "Init stamp PNG",
	Long:  `The command is initialing stamp on png file`,
	RunE: func(cmd *cobra.Command, args []string) error {
		stampText = cmd.PersistentFlags().Lookup("text").Value.String()
		stampExecutor = cmd.PersistentFlags().Lookup("executor").Value.String()

		img, err := pkg.AddTextOnImg(inFile, stampText, stampExecutor)
		if err != nil {
			return fmt.Errorf("problem with setting text on image %w", err)
		}

		pkg.Save(img, outFile)
		if err != nil {
			return fmt.Errorf("problem with saving png file %w", err)
		}
		return nil
	},
}
