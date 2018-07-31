package main

import (
	"testing"
	"github.com/360EntSecGroup-Skylar/excelize"
	"fmt"
)

func Test_EXCEl(t *testing.T) {
	xlsx := excelize.NewFile()
	// 创建一个工作表
	//index := xlsx.NewSheet("Sheet1")
	xlsx.AddTable("Sheet1", "A1", "H3", `{"table_name":"table","table_style":"TableStyleMedium2", "show_first_column":true,"show_last_column":true,"show_row_stripes":false,"show_column_stripes":true}`)
	xlsx.SetCellValue("Sheet1","A1","Database")
	xlsx.InsertRow("Sheet1", 2)
	xlsx.InsertRow("Sheet1", 2)
	//// 设置单元格的值
	//	xlsx.SetCellValue("Sheet2", "A2", "Hello world.")
	//xlsx.SetCellValue("Sheet1", "B2", 100)
	// 设置工作簿的默认工作表
	xlsx.SetActiveSheet(0)
	// 根据指定路径保存文件
	err := xlsx.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}

}
