package main

import (
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/mattn/getwild"
)

func mains(args []string) error {
	items, err := readXmlFiles(args)
	if err != nil {
		return err
	}
	xls, err := excelize.OpenFile("template.xlsx")
	if err != nil {
		return err
	}
	// style, err := xls.NewStyle(`{ "alignment": { "wrap_text":true } }`)

	const sheet = "01.コマンド名"
	for i, item1 := range items {
		xls.SetCellValue(sheet, fmt.Sprintf("B%d", i+3), item1.Case)
		xls.SetCellValue(sheet, fmt.Sprintf("C%d", i+3), item1.Operation)
		xls.SetCellValue(sheet, fmt.Sprintf("E%d", i+3), item1.Status)
	}
	xls.SaveAs("output.xlsx")
	return nil
}

func main() {
	if err := mains(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
