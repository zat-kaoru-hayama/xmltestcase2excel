package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	_ "github.com/mattn/getwild"
)

var (
	flagTemplate = flag.String("template", "template.xlsx", "Template xlsx file")
	flagOutput   = flag.String("o", "output.xlsx", "Output xlsx file")
	flagStartRow = flag.Int("startrow", 3, "Start RowLine")
)

func mains(args []string) error {
	items, err := readXmlFiles(args)
	if err != nil {
		return err
	}
	xls, err := excelize.OpenFile(*flagTemplate)
	if err != nil {
		return err
	}
	// style, err := xls.NewStyle(`{ "alignment": { "wrap_text":true } }`)

	sheetList := xls.GetSheetList()
	sheet := sheetList[len(sheetList)-1]
	startRow := *flagStartRow

	for i, item1 := range items {
		xls.SetCellValue(sheet, fmt.Sprintf("B%d", startRow+i), item1.Case)
		xls.SetCellValue(sheet, fmt.Sprintf("C%d", startRow+i), item1.Operation)
		xls.SetCellValue(sheet, fmt.Sprintf("E%d", startRow+i), item1.Status)
	}
	xls.SaveAs(*flagOutput)
	return nil
}

func main() {
	flag.Parse()
	if err := mains(flag.Args()); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
