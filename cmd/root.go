package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "MFRCli",
	Long: "A GoLang CLI to communicate with the OData interface of mfr",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func GetRootCmd() *cobra.Command {
	return rootCmd
}
