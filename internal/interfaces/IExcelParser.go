package interfaces

type ExcelParser interface {
	GetExcelModel(tNumber string) (interface{}, error)
}
