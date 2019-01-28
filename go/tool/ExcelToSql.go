package main

import (
	"bytes"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"time"
)

// 二维字符串数组即为一张表
type Sheet [][]string

// 获得表头列表
func (s *Sheet) header() *[]string {
	return &(*s)[0]
}

// 遍历[1:n]
func (s *Sheet) foreach(f func(*[]string)) {
	for i, n := 1, len(*s); i < n; i++ {
		f(&(*s)[i])
	}
}

// size-1
func (s *Sheet) length() int {
	return len(*s) - 1
}

// 初始化 命令行参数和excelize
func InitExcel() *excelize.File {
	if len(os.Args) <= 1 {
		os.Exit(1)
	}
	xlsx, err := excelize.OpenFile(os.Args[1])
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	return xlsx
}

// 整个excel文件处理,并返回共有几张表
func ExcelToSql(done chan bool, xlsx *excelize.File) int {
	sheetmap := xlsx.GetSheetMap()
	n := len(sheetmap)
	for _, name := range sheetmap {
		go func(sheetname string) {
			time.Sleep(time.Second)
			s := Sheet(xlsx.GetRows(sheetname))
			if s.length() < 1 {
				println(sheetname, "is empty")
				done <- false
				return
			}
			f, err := os.Create(sheetname + ".sql")
			if err != nil {
				println(sheetname+".sql", "create error", err.Error())
				done <- false
				return
			}
			SheetToSql(f, &sheetname, &s)
			done <- true
		}(name)
	}
	return n
}

// 针对一张表的sql文件写入
func SheetToSql(file *os.File, name *string, sheet *Sheet) {
	h := []byte(HeaderToSql(name, sheet.header()))
	out := file.WriteString
	sheet.foreach(func(row *[]string) {
		file.Write(h)
		out("VALUES (")
		n := len(*row)
		for i, col := range *row {
			out("\"")
			out(col)
			out("\"")
			if i+1 != n {
				out(",")
			}
		}
		out(");\n")
	})
}

// insert into [dbname] () values ()\n
func HeaderToSql(name *string, header *[]string) string {
	buf := bytes.Buffer{}
	out := buf.WriteString
	out("INSERT INTO ")
	out(*name)
	out(" (")
	n := len(*header)
	for i, col := range *header {
		out(col)
		if i+1 != n {
			out(",")
		}
	}
	out(") ")
	return buf.String()
}

func main() {
	done := make(chan bool)
	n := ExcelToSql(done, InitExcel())
	for i := 0; i < n; i++ {
		<-done
	}
}
