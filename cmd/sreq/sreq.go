package cmd

import (
	rootCmd "github.com/A-Oez/go-mfr/cmd"
	"github.com/A-Oez/go-mfr/internal/http"
	"github.com/A-Oez/go-mfr/internal/service"
	"github.com/spf13/cobra"
)

var sreqCmd = &cobra.Command{
	Use:   "sreq",
	Short: "Receive ServiceRequests based on specified T-numbers and enter content in Excel Sheet",
	Run:   cmdRun,
}

func init() {
	rootCmd.GetRootCmd().AddCommand(sreqCmd)
	sreqCmd.PersistentFlags().String("d", "", "excel path")
}

func cmdRun(cmd *cobra.Command, args []string) {
	excelPath, _ := cmd.Flags().GetString("d")
	service.HandleServiceRequestExport(excelPath, &http.HttpHandler{})
}