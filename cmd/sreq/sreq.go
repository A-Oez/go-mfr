package cmd

import (
	rootCmd "github.com/A-Oez/go-mfr/cmd"
	"github.com/A-Oez/go-mfr/internal/http"
	"github.com/A-Oez/go-mfr/internal/service"
	"github.com/A-Oez/go-mfr/pkg"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var sreqCmd = &cobra.Command{
	Use:   "sreq",
	Short: "Import/Export von ServiceRequests anhand von T-Nummern in eine Excel-Datei.",
	Run:   cmdRun,
}

func init() {
	rootCmd.GetRootCmd().AddCommand(sreqCmd)
}

func cmdRun(cmd *cobra.Command, args []string) {
	err := service.HandleServiceRequestExport(pkg.GetProperty("excel_path"), &http.HttpHandler{})
	if err != nil{
		pterm.Error.Println(err)
	} else {
		pterm.Success.Println("Export erfolgreich!")
	}
}