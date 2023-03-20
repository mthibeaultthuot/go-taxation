package main

import (
	"github.com/D3xt3rrrr/verificationTax-service/data"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
)

func CreateExcelFile(tax data.Tax) (*os.File, error) {
	header := [6]string{"IsPstValid", "IsQstValid", "PstNumber", "QstNumber", "Enterprise", "Date"}
	file := excelize.NewFile()
	sheet := file.NewSheet("search")

	err := file.SetCellValue("search", "A1", "hwllorodl")

	index := 'A'
	for v := range header {
		file.SetCellValue("search", string(index)+"1", header[v])
		index = index + 1
	}

	file.SetCellValue("search", "A2", tax.IsPstValid)
	file.SetCellValue("search", "B2", tax.IsQstValid)
	file.SetCellValue("search", "C2", tax.PstNumber)
	file.SetCellValue("search", "D2", tax.QstNumber)
	file.SetCellValue("search", "E2", tax.Enterprise)
	file.SetCellValue("search", "F2", tax.Date)

	file.SetActiveSheet(sheet)
	err = file.SaveAs("search.xlsx")
	if err != nil {
		log.Panic(err)
	}

	return os.Open("search.xlsx")
}
