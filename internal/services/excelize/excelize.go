package excelizeservice

import (
	"fmt"
	"log"

	"bookingrooms/internal/models"
	"github.com/xuri/excelize/v2"
)

type service struct {
}

func NewExcelizeService() service {
	return service{}
}

type ExcelizeService interface {
	WriteRoomsToNewSheetExcel() error
}

func (s service) WriteRoomsToNewSheetExcel() error {
	rooms := []models.Rooms{
		{ID: 1,
			Name:          "a",
			RoomName:      "M1,",
			MaximumPerson: 10},
		{ID: 2,
			Name:          "b",
			RoomName:      "M2,",
			MaximumPerson: 10}}

	// open file
	f, err := excelize.OpenFile("book1.xlsx")
	if err != nil {
		return err
	}
	defer func() {
		// Close the spreadsheet.
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	index := f.NewSheet("Rooms")
	categories := map[string]string{
		"A1": "ID", "B1": "Room Name", "C1": "Maximum Person", "D1": "Name"}
	for i, c := range categories {
		err = f.SetCellValue("Rooms", i, c)
		if err != nil {
			return err
		}
	}

	for i, room := range rooms {
		err = f.SetCellValue("Rooms", fmt.Sprintf("A%d", i+2), room.ID)
		err = f.SetCellValue("Rooms", fmt.Sprintf("B%d", i+2), room.RoomName)
		err = f.SetCellValue("Rooms", fmt.Sprintf("C%d", i+2), room.MaximumPerson)
		err = f.SetCellValue("Rooms", fmt.Sprintf("D%d", i+2), room.Name)
	}

	f.SetActiveSheet(index)
	return f.Save()
}

func (s service) CreatedExcel() error {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	err := f.SetCellValue("Sheet2", "A2", "Hello world.")
	if err != nil {
		return err
	}
	err = f.SetCellValue("Sheet1", "B2", 100)
	if err != nil {
		return err
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	err = f.SaveAs("book1.xlsx")
	if err != nil {
		return err
	}
	return nil
}

func (s service) ReadExcel() ([][]string, error) {
	f, err := excelize.OpenFile("book1.xlsx")
	if err != nil {
		return [][]string{}, err
	}
	defer func() {
		// Close the spreadsheet.
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	// Get value from cell by given worksheet name and cell reference.
	// Get data 1 cell
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		return [][]string{}, err
	}
	log.Println(cell)

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return [][]string{}, err
	}

	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
	return rows, nil
}

func (s service) AddChartToSpreadsheetFile() error {
	categories := map[string]string{
		"A2": "Small", "A3": "Normal", "A4": "Large",
		"B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{
		"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	f := excelize.NewFile()
	for k, v := range categories {
		err := f.SetCellValue("Sheet1", k, v)
		if err != nil {
			return err
		}
	}
	for k, v := range values {
		err := f.SetCellValue("Sheet1", k, v)
		if err != nil {
			return err
		}
	}
	if err := f.AddChart("Sheet1", "E1", `{
        "type": "col3DClustered",
        "series": [
        {
            "name": "Sheet1!$A$2",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$2:$D$2"
        },
        {
            "name": "Sheet1!$A$3",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$3:$D$3"
        },
        {
            "name": "Sheet1!$A$4",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$4:$D$4"
        }],
        "title":
        {
            "name": "Fruit 3D Clustered Column Chart"
        }
    }`); err != nil {
		fmt.Println(err)
		return err
	}
	// Save spreadsheet by the given path.
	if err := f.SaveAs("book1.xlsx"); err != nil {
		fmt.Println(err)
	}
	return nil
}
