package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"stamp-pdf/pkg"

	"github.com/spf13/cobra"
)

var stampCmd = &cobra.Command{
	Use:   "stamp",
	Short: "Set stamp",
	Long:  `Command has injected stamp in pdf file`,
	RunE: func(cmd *cobra.Command, args []string) error {
		page = cmd.PersistentFlags().Lookup("page").Value.String()
		stampPosition = cmd.PersistentFlags().Lookup("pos").Value.String()

		if page != "1-" {
			file, err := pkg.ExtractPageFromFile(inFile, page)
			if err != nil {
				return err
			}
			inFile = file
			defer os.Remove(inFile)
		}

		if stamp {

			stampApplicant = cmd.PersistentFlags().Lookup("applicant").Value.String()
			stampText = cmd.PersistentFlags().Lookup("text").Value.String()
			stampExecutor = cmd.PersistentFlags().Lookup("executor").Value.String()

			dir, err := os.Getwd()
			if err != nil {
				return err
			}

			template := fmt.Sprintf("%s/%s", dir, "stamp_template.png")

			img, err := pkg.AddTextOnImg(template, stampText, stampExecutor, stampApplicant)
			if err != nil {
				return err
			}

			stampFile = fmt.Sprintf("%s/stamp_out_%d.png", filepath.Dir(outFile), time.Now().Nanosecond())
			defer os.Remove(stampFile)

			err = pkg.Save(img, stampFile)
			if err != nil {
				return err
			}
		}

		err := pkg.AddStamp(inFile, outFile, page, stampFile, stampPosition)
		if err != nil {
			return err
		}
		return nil
	},
}
