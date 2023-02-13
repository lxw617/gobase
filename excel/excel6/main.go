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
	if err := f.SetRowHeight("Sheet1", 1, 120); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.MergeCell("Sheet1", "A1", "M1"); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SetCellRichText("Sheet1", "A1", []excelize.RichTextRun{
		{
			Text: "备注：此表格可用于批量发货、批量取消标识的回传模板使用。",
			Font: &excelize.Font{
				Bold: true,
			},
		},
		{
			Text: "【如需回传表格，请按规范填写以下表格内容，若输入错误词条，无法回传系统。】\r\n1. ",
		},
		{
			Text: "订单号：",
			Font: &excelize.Font{
				Bold:  true,
				Color: "2354e8",
			},
		},
		{
			Text: "为重要的匹配信息点，每个订单有唯一的订单号，此信息",
		},
		{
			Text: "不可更改。",
			Font: &excelize.Font{
				Bold:  true,
				Color: "e83723",
			},
		},
		{
			Text: "\r\n2. 物流公司和物流单号：按系统词条填写快递公司名称，快递单号在发货后根据物流提供的单号实际填写即可。\r\n3. 发货备注：可手动填写发货需备注的信息点。如：本次不发货的原因；11.30发货…\r\n4. 发货结果、失败原因、最新物流状态、最新物流说明：",
		},
		{
			Text: "无需填写，系统自动识别；",
			Font: &excelize.Font{
				Bold:  true,
				Color: "e83723",
			},
		},
		{
			Text: "此4项内容不可删改！\r\n5. ",
		},
		{
			Text: "取消原因：",
			Font: &excelize.Font{
				Bold:  true,
				Color: "2354e8",
			},
		},
		{
			Text: "批量取消标识时必填，批量发货不填写此项，【此项",
		},
		{
			Text: "仅可填写为：疫情等原因、商品补货中、商品原因、其他原因",
			Font: &excelize.Font{
				Bold:  true,
				Color: "e83723",
			},
		},
		{
			Text: "】；",
		},
		{
			Text: "取消备注：",
			Font: &excelize.Font{
				Bold:  true,
				Color: "2354e8",
			},
		},
		{
			Text: "批量取消标识时填写，对取消原因作补充说明，此项可不填。\r\n6. 取消结果、取消失败说明：",
		},
		{
			Text: "无需填写，系统自动识别；",
			Font: &excelize.Font{
				Bold:  true,
				Color: "e83723",
			},
		},
		{
			Text: "此2项内容不可删改！",
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
	if err := f.SaveAs("./temp/Book2.xlsx"); err != nil {
		fmt.Println(err)
	}
}
