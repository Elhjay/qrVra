package excelService

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/tealeg/xlsx"
)

// Define a struct to represent your data structure
type Person struct {
	Pid        string `json:"pid"`
	Name       string `json:"name"`
	Cell       string `json:"cell"`
	Work       string `json:"work"`
	Email      string `json:"email"`
	Department string `json:"department"`
	Section    string `json:"section"`
}

func DataFromExcel(excelFileName string) []Person {
	//excelFileName := "vCardsheet.xlsx"

	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("Error opening Excel file:", err)
		return []Person{}
	}

	sheet := xlFile.Sheets[0]

	var people []Person

	for rowIndex, row := range sheet.Rows {
		if rowIndex == 0 {
			continue
		}

		person := Person{}

		for cellIndex, cell := range row.Cells {
			cellValue := cell.String()
			person.Pid = uuid.NewString()

			switch cellIndex {
			case 0:
				person.Name = cellValue
			case 1:
				person.Cell = cellValue
			case 2:
				person.Work = cellValue
			case 3:
				person.Email = cellValue
			case 4:
				person.Department = cellValue
			case 5:
				person.Section = cellValue
			}
		}

		people = append(people, person)
	}

	return people

}
