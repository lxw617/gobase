package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	if err := f.SetRowHeight("Sheet1", 1, 35); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SetColWidth("Sheet1", "A", "A", 44); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SetCellRichText("Sheet1", "A1", []excelize.RichTextRun{
		{
			Text: "bold",
			Font: &excelize.Font{
				Bold:   true,
				Color:  "2354e8",
				Family: "Times New Roman",
			},
		},
		{
			Text: " and ",
			Font: &excelize.Font{
				Family: "Times New Roman",
			},
		},
		{
			Text: "italic ",
			Font: &excelize.Font{
				Bold:   true,
				Color:  "e83723",
				Italic: true,
				Family: "Times New Roman",
			},
		},
		{
			Text: "text with color and font-family,",
			Font: &excelize.Font{
				Bold:   true,
				Color:  "2354e8",
				Family: "Times New Roman",
			},
		},
		{
			Text: "\r\nlarge text with ",
			Font: &excelize.Font{
				Size:  14,
				Color: "ad23e8",
			},
		},
		{
			Text: "strike",
			Font: &excelize.Font{
				Color:  "e89923",
				Strike: true,
			},
		},
		{
			Text: " superscript",
			Font: &excelize.Font{
				Color:     "dbc21f",
				VertAlign: "superscript",
			},
		},
		{
			Text: " and ",
			Font: &excelize.Font{
				Size:      14,
				Color:     "ad23e8",
				VertAlign: "baseline",
			},
		},
		{
			Text: "underline",
			Font: &excelize.Font{
				Color:     "23e833",
				Underline: "single",
			},
		},
		{
			Text: " subscript.",
			Font: &excelize.Font{
				Color:     "017505",
				VertAlign: "subscript",
			},
		},
	}); err != nil {
		fmt.Println(err)
		return
	}
	style, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			WrapText: true,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SetCellStyle("Sheet1", "A1", "A1", style); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SaveAs("./temp/excel5.xlsx"); err != nil {
		fmt.Println(err)
	}
}
