package cmd

import (
	excelUtils "MFRCli/pkg/excelutils"

	excelTemplates "MFRCli/pkg/excelutils/excel_templates"

	parser "MFRCli/internal/http/parser"

	"github.com/spf13/cobra"
)

var sreqCmd = &cobra.Command{
	Use:   "sreq",
	Short: "Receive ServiceRequests based on specified T-numbers and enter content in Excel Sheet",
	Run:   cmdRun,
}

func init() {
	var exp bool

	GetRootCmd().AddCommand(sreqCmd)
	sreqCmd.PersistentFlags().String("d", "", "excel path")
	sreqCmd.PersistentFlags().BoolVar(&exp, "exp", false, "Export ServiceRequests")
}

func cmdRun(cmd *cobra.Command, args []string) {
	excelPath, _ := cmd.Flags().GetString("d")
	exp, _ := cmd.Flags().GetBool("exp")
	tNumbers := excelUtils.GetTNumbers(excelPath)

	if exp {
		exportServiceRequests(excelPath, tNumbers)
	}
}

func exportServiceRequests(excelPath string, tNumbers []string) {
	for i := range tNumbers {
		excelTemplates.WriteToExcel(excelPath, parser.GetExcelModel(tNumbers[i]))
	}
}
