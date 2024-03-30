package cmd

import (
	"fmt"

	excelUtils "github.com/A-Oez/MFRCli/pkg/excel_utils"

	"github.com/A-Oez/MFRCli/internal/service/excel_handler"

	pReader "github.com/A-Oez/MFRCli/pkg"

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
		exportGeneralSREQ(excelPath, tNumbers)
	} else if ter {
		exportAddressSREQ(excelPath, tNumbers)
	}
}

func exportGeneralSREQ(excelPath string, tNumbers []string) {
	var excelHandler = &excel_handler.SREQGeneral{}

	for i := range tNumbers {
		SREQGeneral, err := excelHandler.GetExcelModel(tNumbers[i])

		if err == nil {
			excelHandler.WriteExcel(excelPath, SREQGeneral)
		} else {
			fmt.Println(fmt.Sprintf("* %s %s %s", pReader.GetProperty("serviceRequestExport"), tNumbers[i], err.Error()))
		}

	}

}

func exportAddressSREQ(excelPath string, tNumbers []string) {
	var excelHandler = &excel_handler.SREQAddress{}

	for i := range tNumbers {
		SREQAddressArr, err := excelHandler.GetExcelModel(tNumbers[i])

		if err == nil {
			for j := range SREQAddressArr {
				excelHandler.WriteExcel(excelPath, SREQAddressArr[j])
			}
		} else {
			fmt.Println(fmt.Sprintf("* %s %s %s", pReader.GetProperty("serviceRequestAddress"), tNumbers[i], err.Error()))
		}

	}
}
