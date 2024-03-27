package cmd

import (
	excelUtils "MFRCli/pkg/excelutils"
	"fmt"

	excelTemplates "MFRCli/pkg/excelutils/excel_templates"

	parser "MFRCli/internal/http/parser"

	pReader "MFRCli/pkg"

	"github.com/spf13/cobra"
)

var sreqCmd = &cobra.Command{
	Use:   "sreq",
	Short: "Receive ServiceRequests based on specified T-numbers and enter content in Excel Sheet",
	Run:   cmdRun,
}

func init() {
	var exp bool
	var ter bool

	GetRootCmd().AddCommand(sreqCmd)
	sreqCmd.PersistentFlags().String("d", "", "excel path")
	sreqCmd.PersistentFlags().BoolVar(&exp, "exp", false, "Export ServiceRequests")
	sreqCmd.PersistentFlags().BoolVar(&ter, "ter", false, "Export ServiceRequests")
}

func cmdRun(cmd *cobra.Command, args []string) {
	excelPath, _ := cmd.Flags().GetString("d")
	exp, _ := cmd.Flags().GetBool("exp")
	ter, _ := cmd.Flags().GetBool("ter")
	tNumbers := excelUtils.GetTNumbers(excelPath)

	if exp {
		exportServiceRequests(excelPath, tNumbers)
	} else if ter {
		exportServiceRequestsAddress(excelPath, tNumbers)
	}
}

func exportServiceRequests(excelPath string, tNumbers []string) {
	for i := range tNumbers {
		excelServiceRequest, err := parser.GetExcelModel(tNumbers[i])

		if err == nil {
			excelTemplates.WriteToExcel(excelPath, excelServiceRequest)
		} else {
			fmt.Println(fmt.Sprintf("* %s %s %s", pReader.GetProperty("serviceRequestExport"), tNumbers[i], err.Error()))
		}
	}

}

func exportServiceRequestsAddress(excelPath string, tNumbers []string) {
	for i := range tNumbers {
		excelServiceRequestAddress := parser.GetExcelAddressModel(tNumbers[i])
		for j := range excelServiceRequestAddress {
			excelTemplates.WriteToAddressExcel(excelPath, excelServiceRequestAddress[j], tNumbers[i])
		}
	}
}
