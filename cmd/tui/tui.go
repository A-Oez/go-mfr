package cmd

import (
	"fmt"
	"os"

	rootCmd "github.com/A-Oez/go-mfr/cmd"
	"github.com/A-Oez/go-mfr/internal/http"
	"github.com/A-Oez/go-mfr/internal/service"
	"github.com/A-Oez/go-mfr/pkg"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "TUI - A GoLang CLI to communicate with the REST interface of mfr",
	Run:   cmdRun,
}

var (
	options = []string{"Auftragsexport"} 
)

func init() {
	rootCmd.GetRootCmd().AddCommand(tuiCmd)
}

func cmdRun(cmd *cobra.Command, args []string) {
	setupHeader()
	
	o := selectOptions();
	pterm.Info.Println(o)

	if(o == "Auftragsexport"){
		err := service.HandleServiceRequestExport(pkg.GetProperty("excel_path"), &http.HttpHandler{})
		if err != nil{
			pterm.Error.Println(err)
		} else {
			pterm.Success.Println("Export erfolgreich!")
		}
	}
}

func setupHeader(){
	pterm.DefaultBigText.WithLetters(putils.LettersFromStringWithStyle(pkg.GetProperty("tui_header"), pterm.FgLightMagenta.ToStyle())).Render()
	pterm.Println()
}

func selectOptions() string{
	selectedOption, err := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
	if err != nil {
		err := fmt.Errorf("service konnte nicht ausgewählt werden: %v", err)
		pterm.Error.Println(err)
		os.Exit(1)
	}
	pterm.Println()
	return selectedOption	
}