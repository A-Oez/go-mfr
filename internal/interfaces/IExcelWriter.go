package interfaces

type ExcelWriter interface {
	WriteExcel(filePath string, model interface{})
}
