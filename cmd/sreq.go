package cmd

import (
	"fmt"

	excelUtils "github.com/A-Oez/MFRCli/pkg/excelutils"

	parser "github.com/A-Oez/MFRCli/internal/parser"

	pReader "github.com/A-Oez/MFRCli/pkg"

	"github.com/spf13/cobra"

	interfaces "github.com/A-Oez/MFRCli/internal/interfaces"
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
	var excelWriter interfaces.ExcelWriter

	for i := range tNumbers {
		excelServiceRequest, err := parser.GetExcelModel(tNumbers[i])
		if err == nil {
			excelWriter = &excelServiceRequest
			excelWriter.WriteExcel(excelPath, excelServiceRequest)
		} else {
			fmt.Println(fmt.Sprintf("* %s %s %s", pReader.GetProperty("serviceRequestExport"), tNumbers[i], err.Error()))
		}
	}

}

func exportServiceRequestsAddress(excelPath string, tNumbers []string) {
	var excelWriter interfaces.ExcelWriter

	for i := range tNumbers {
		excelServiceRequestAddressArr, err := parser.GetExcelAddressModel(tNumbers[i])

		if err == nil {
			for j := range excelServiceRequestAddressArr {
				excelServiceRequestAddress := excelServiceRequestAddressArr[j]

				excelWriter = &excelServiceRequestAddress
				excelWriter.WriteExcel(excelPath, excelServiceRequestAddressArr[j])
			}
		} else {
			fmt.Println(fmt.Sprintf("* %s %s %s", pReader.GetProperty("serviceRequestAddress"), tNumbers[i], err.Error()))
		}
	}
}
