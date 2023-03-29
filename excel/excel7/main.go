package main

import (
	"fmt"
	"os"

	"github.com/tealeg/xlsx"
	"github.com/xuri/excelize/v2"
)

var (
	outFile = "./temp/excel7.xlsx"
)

func main() {
	tableName := "ceshi"
	fileName := "out_excel7"
	tableHead := []interface{}{111, 222}
	sheet, file, fileUrl, fileName, _ := GetPointExcelHeader(fileName, tableHead, tableName)
	err := file.Save(outFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = GetPointExcel(fileName, tableName, 1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(sheet)
	fmt.Println(fileUrl)
	fmt.Println("\n\nexport success")
}
func GetExcel(fileName string, tableName string) (sheet *xlsx.Sheet, file *xlsx.File, url string, name string, err error) {
	root, _ := os.Getwd()
	path := root + "/temp"
	name = fileName + ".xlsx"
	url = path + "/" + name
	file = xlsx.NewFile()
	sheet, err = file.AddSheet(tableName)
	return
}

func GetPointExcelHeader(fileName string, values []interface{}, tableName string) (sheet *xlsx.Sheet, file *xlsx.File, url string, name string, err error) {
	sheet, file, url, name, err = GetExcel(fileName, tableName)
	if err != nil {
		return nil, nil, "", "", err
	}
	sheet.AddRow()
	sheet.AddRow().WriteSlice(&values, -1)
	return
}
func GetPointExcel(fileName string, tableName string, excel int) error {
	root, _ := os.Getwd()
	path := root + "/temp"
	url := path + "/" + fileName
	excelizefFile, err := excelize.OpenFile(url)
	if err != nil {
		return err
	}
	GetPointRemark(excelizefFile, url, tableName, excel)
	return nil
}

func GetPointRemark(f *excelize.File, fileName, tableName string, excel int) {
	blueFont := &excelize.Font{
		Bold:  true,
		Color: "2354e8",
		Size:  12,
	}
	redFont := &excelize.Font{
		Bold:  true,
		Color: "e83723",
		Size:  12,
	}
	boldFont := &excelize.Font{
		Bold: true,
		Size: 12,
	}
	defaultFont := &excelize.Font{
		Size: 12,
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}()
	if err := f.SetRowHeight(tableName, 1, 120); err != nil {
		fmt.Println(err)
		return
	}
	if excel == 1 {
		if err := f.MergeCell(tableName, "A1", "M1"); err != nil {
			fmt.Println(err)
			return
		}
	} else {
		if err := f.MergeCell(tableName, "A1", "Y1"); err != nil {
			fmt.Println(err)
			return
		}
	}
	if excel == 1 {
		if err := f.SetCellRichText(tableName, "A1", []excelize.RichTextRun{
			richText("备注：此表格可用于批量补发的回传模板使用。", boldFont),
			richText("【如需回传表格，请按规范填写以下表格内容，若输入错误词条，无法回传系统。】\r\n1. ", defaultFont),
			richText("订单号：", blueFont),
			richText("为重要的匹配信息点，每个订单有唯一的订单号。", defaultFont),
			richText("商品名称：", blueFont),
			richText("需补发商品的名称。】\r\n2. 收货人、手机号、收货地址：根据用户提供收件信息填写（地址需填写收件的", defaultFont),
			richText("详细地址", redFont),
			richText("）。\r\n3. 发货方式：【", defaultFont),
			richText("仅可填写为：无需物流、物流发货", redFont),
			richText("】；\r\n4. 物流公司、物流单号：按系统词条填写快递公司名称，快递单号在发货后根据物流提供的单号实际填写即可。\r\n5. 补发类型：【", defaultFont),
			richText("仅可填写为：货损补发、退回再发", redFont),
			richText("】；此项用于对库存的计算，库存无变化都可以填写【退回再发】，如：揽件失败、少发或漏发...\r\n6. ", defaultFont),
			richText("补发原因：", blueFont),
			richText("【", defaultFont),
			richText("仅可填写为：原快递物品损坏、少发或漏发、物流退回、其他原因", redFont),
			richText("】；", defaultFont),
			richText("补发备注：", blueFont),
			richText("对补发信息补充说明，此项可不填", defaultFont),
		}); err != nil {
			fmt.Println(err)
			return
		}
	} else {
		if err := f.SetCellRichText(tableName, "A1", []excelize.RichTextRun{
			richText("备注：此表格可用于批量发货、批量取消标识的回传模板使用。", boldFont),
			richText("【如需回传表格，请按规范填写以下表格内容，若输入错误词条，无法回传系统。】\r\n1. ", defaultFont),
			richText("订单号：", blueFont),
			richText("为重要的匹配信息点，每个订单有唯一的订单号，此信息", defaultFont),
			richText("不可更改。", redFont),
			richText("\r\n2. 物流公司和物流单号：按系统词条填写快递公司名称，快递单号在发货后根据物流提供的单号实际填写即可。\r\n3. 发货备注：可手动填写发货需备注的信息点。如：本次不发货的原因；11.30发货…\r\n4. 发货结果、失败原因、最新物流状态、最新物流说明：", defaultFont),
			richText("无需填写，系统自动识别；", redFont),
			richText("此4项内容不可删改！\n5. ", defaultFont),
			richText("取消原因：", blueFont),
			richText("批量取消标识时必填，批量发货不填写此项，【此项", defaultFont),
			richText("仅可填写为：疫情等原因、商品补货中、商品原因、其他原因", redFont),
			richText("】；", defaultFont),
			richText("取消备注：", blueFont),
			richText("批量取消标识时填写，对取消原因作补充说明，此项可不填。\n6. 取消结果、取消失败说明：", defaultFont),
			richText("无需填写，系统自动识别；", redFont),
			richText("此2项内容不可删改！", defaultFont),
		}); err != nil {
			fmt.Println(err)
			return
		}
	}
	style, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			WrapText: true,
			Vertical: "center",
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SetCellStyle(tableName, "A1", "A1", style); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
		return
	}

}

func richText(test string, font *excelize.Font) excelize.RichTextRun {
	richTextRun := excelize.RichTextRun{
		Text: test,
		Font: font,
	}
	return richTextRun
}
