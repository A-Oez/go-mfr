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
		excelServiceRequest, err := parser.GetExcelModel(tNumbers[i])

		if err == nil {
			excelTemplates.WriteToExcel(excelPath, excelServiceRequest)
			exportServiceRequestsAddress(excelPath, excelServiceRequest.Description, excelServiceRequest.Auftragsname, tNumbers[i])
		} else {
			fmt.Println(fmt.Sprintf("* %s %s %s", pReader.GetProperty("serviceRequestExport"), tNumbers[i], err.Error()))
		}
	}

}

func exportServiceRequestsAddress(excelPath string, description string, orderName string, tNumber string) {
	excelServiceRequestAddress := parser.GetExcelAddressModel(description, orderName)
	for j := range excelServiceRequestAddress {
		excelTemplates.WriteToAddressExcel(excelPath, excelServiceRequestAddress[j], tNumber)
	}
}
