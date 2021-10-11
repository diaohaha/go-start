package main

import (
	"errors"
	"fmt"
	"path"

	"github.com/tealeg/xlsx"
)

// 输入xlsx或者文件流 解析出list 每行为一个map

func ExcelParse(filePath string, fileBytes []byte) (rows []map[string]string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("解析失败:%r", r))
		}
	}()
	var xlFile *xlsx.File
	if filePath != "" {
		if path.Ext(filePath) != ".xlsx" {
			return nil, errors.New("文件格式错误")
		}
		xlFile, err = xlsx.OpenFile(filePath)
		if err != nil {
			return nil, err
		}

	} else {
		xlFile, err = xlsx.OpenBinary(fileBytes)
		if err != nil {
			return nil, err
		}

	}

	headMap := map[int]string{}

	emptyFlag := true
	for _, sheet := range xlFile.Sheets {
		for rowIndex, row := range sheet.Rows {
			if rowIndex == 0 {
				for i, cell := range row.Cells {
					text := cell.String()
					headMap[i] = text
				}
				continue
			}
			// 单列可空 全列不可
			emptyFlag = true
			rowInfo := map[string]string{}
			for cellIndex, cell := range row.Cells {
				text := cell.String()
				if text != "" {
					emptyFlag = false
					rowInfo[headMap[cellIndex]] = text
				} else {
					rowInfo[headMap[cellIndex]] = text
				}
			}
			if !emptyFlag {
				rows = append(rows, rowInfo)
			}
		}
	}
	return
}

