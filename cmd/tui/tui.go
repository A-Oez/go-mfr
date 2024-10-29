package cmd

import (
	rootCmd "github.com/A-Oez/go-mfr/cmd"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "TUI - A GoLang CLI to communicate with the REST interface of mfr",
	Run:   cmdRun,
}

func init() {
	rootCmd.GetRootCmd().AddCommand(tuiCmd)
}

func cmdRun(cmd *cobra.Command, args []string) {
	
}