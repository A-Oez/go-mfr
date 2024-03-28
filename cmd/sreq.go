package cmd

import (
	"fmt"

	excelUtils "github.com/A-Oez/MFRCli/pkg/excel_utils"

	excelHandler "github.com/A-Oez/MFRCli/internal/service/excel_handler"

	pReader "github.com/A-Oez/MFRCli/pkg"

	excelModel "github.com/A-Oez/MFRCli/internal/model/excel_model"

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
		exportGeneralSREQ(excelPath, tNumbers)
	} else if ter {
		exportAddressSREQ(excelPath, tNumbers)
	}
}

func exportGeneralSREQ(excelPath string, tNumbers []string) {
	var excelParser interfaces.ExcelParser = &excelHandler.SREQGeneral{}
	var excelWriter interfaces.ExcelWriter = &excelHandler.SREQGeneral{}

	for i := range tNumbers {
		model, err := excelParser.GetExcelModel(tNumbers[i])

		if SREQGeneral, ok := model.(excelModel.SREQGeneral); ok {
			if err == nil {
				excelWriter.WriteExcel(excelPath, SREQGeneral)
			} else {
				fmt.Println(fmt.Sprintf("* %s %s %s", pReader.GetProperty("serviceRequestExport"), tNumbers[i], err.Error()))
			}
		}
	}

}

func exportAddressSREQ(excelPath string, tNumbers []string) {
	var excelParser interfaces.ExcelParser = &excelHandler.SREQAddress{}
	var excelWriter interfaces.ExcelWriter = &excelHandler.SREQAddress{}

	for i := range tNumbers {
		model, err := excelParser.GetExcelModel(tNumbers[i])

		if SREQAddressArr, ok := model.([]excelModel.SREQAddress); ok {
			if err == nil {
				for j := range SREQAddressArr {
					excelWriter.WriteExcel(excelPath, SREQAddressArr[j])
				}
			} else {
				fmt.Println(fmt.Sprintf("* %s %s %s", pReader.GetProperty("serviceRequestAddress"), tNumbers[i], err.Error()))
			}
		}

	}
}
