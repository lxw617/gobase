package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func main() {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("标签页1")
	if err != nil {
		fmt.Println(err.Error())
	}
	//设置表格标题
	title := sheet.AddRow()
	titleRow := title.AddCell()
	titleRow.HMerge = 11 //向右合并列数，不包括自身列
	titleRow.Value = "this is excel title"

	//新建2行
	header1 := sheet.AddRow()
	header2 := sheet.AddRow()

	span1 := header1.AddCell()

	// 注意 这里非常重要
	// 空白格，用于被合并，否则后续的单元格会出现合并格式错误
	header2.AddCell()
	span1.VMerge = 1 // span1 向下合并1格
	span1.Value = "向下合并的单元格"

	span2 := header1.AddCell()
	header2.AddCell()
	span2.VMerge = 1
	span2.Value = "向下合并的单元格"

	span3 := header1.AddCell()

	// span3 向右合并一格
	// 所在行需要增加空白格用于合并
	header1.AddCell()
	span3.HMerge = 1
	span3.Value = "向右合并的单元格"
	span4 := header2.AddCell()
	span4.Value = "子单元格-1"
	span5 := header2.AddCell()
	span5.Value = "子单元格-2"

	span6 := header1.AddCell()
	span7 := header2.AddCell()
	span6.Value = "单元格6"
	span7.Value = "单元格7"

	file.Save("./temp/excel1.xlsx")

}
