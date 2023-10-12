package commands

import (
	"fmt"
	"os"
)

func init() {
	rootCmd.PersistentFlags().StringVar(&inFile, "file", "", "the path for file document")
	rootCmd.PersistentFlags().StringVarP(&outFile, "out", "o", "./", "the path for result pdf document")
	rootCmd.PersistentFlags().StringVar(&stampFile, "stamp-file", "", "the path for stamp png")

	stampCmd.PersistentFlags().StringP("page", "p", "1-", "the page number for stamp (default 1- all pages)")
	stampCmd.PersistentFlags().StringP("pos", "", "bl", "the position for stamp setting on the page. Default bl (bottom left)")

	stampCmd.PersistentFlags().BoolVar(&stamp, "stamp", false, "the stamp set over for one action")
	stampCmd.PersistentFlags().StringP("applicant", "", "", "the applicant for setting on stamp")
	stampCmd.PersistentFlags().StringP("executor", "", "", "the executor for setting on stamp")
	stampCmd.PersistentFlags().StringP("text", "", "", "the text for setting on stamp")
	stampCmd.MarkFlagsRequiredTogether("stamp", "applicant", "text", "executor")

	stampInitCmd.PersistentFlags().StringP("applicant", "", "", "the applicant for setting on stamp")
	stampInitCmd.PersistentFlags().StringP("text", "", "", "the text for setting on stamp")
	stampInitCmd.PersistentFlags().StringP("executor", "", "", "the executor for setting on stamp")
	stampInitCmd.MarkFlagsRequiredTogether("applicant", "text", "executor")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(stampCmd)
	rootCmd.AddCommand(stampInitCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Remove(inFile)
		os.Remove(outFile)
		os.Remove(stampFile)
		os.Exit(1)
	}
}
