package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// 可以解析xlsx和csv两种文件,xlsx为多个Sheet,csv为一张Sheet

// Sheet 二维字符串数组即为一张表
type Sheet struct {
	Name string     // 表名
	Data [][]string // 核心数据
}

// Header 获得表头列表
func (s *Sheet) Header() []string {
	return s.Data[0]
}

// Foreach 遍历[1:n]
func (s *Sheet) Foreach(f func(*[]string)) {
	for i, n := 1, len(s.Data); i < n; i++ {
		f(&s.Data[i])
	}
}

// Length size-1
func (s *Sheet) Length() int {
	return len(s.Data) - 1
}

// NewSheet new(Sheet)
func NewSheet(Name string, Data [][]string) *Sheet {
	return &Sheet{Name, Data}
}

// InitSheets 初始化 命令行参数和excelize
func InitSheets() []*Sheet {
	if len(os.Args) <= 1 {
		os.Exit(1)
	}
	if strings.HasSuffix(os.Args[1], ".csv") {
		table, err := os.Open(os.Args[1])
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
		defer table.Close()
		reader := csv.NewReader(table)
		data := make([][]string, 0)
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				println(err.Error())
				os.Exit(1)
			}
			data = append(data, record)
		}
		return []*Sheet{NewSheet(strings.TrimSuffix(filepath.Base(table.Name()), ".csv"), data)}
	}
	if strings.HasSuffix(os.Args[1], ".xlsx") {
		xlsx, err := excelize.OpenFile(os.Args[1])
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
		ss := make([]*Sheet, 0)
		sheetmap := xlsx.GetSheetMap()
		for _, name := range sheetmap {
			ss = append(ss, NewSheet(name, xlsx.GetRows(name)))
		}
		return ss
	}
	println("file ext is not [.csv, .xlsx]")
	os.Exit(1)
	return nil
}

// SheetsToSQL 整个excel文件处理,并返回共有几张表
func SheetsToSQL(done chan bool, sheets []*Sheet) {
	for _, s := range sheets {
		go func(sheet *Sheet) {
			time.Sleep(time.Second)
			if sheet.Length() < 1 {
				println(sheet.Name, "is empty")
				done <- false
				return
			}
			sqlF, err := os.Create(sheet.Name + ".sql")
			if err != nil {
				println(sheet.Name+".sql", "create error", err.Error())
				done <- false
				return
			}
			SheetToSQL(sheet, sqlF)
			done <- true
		}(s)
	}
}

// SheetToSQL 针对一张表的sql文件写入
func SheetToSQL(sheet *Sheet, file *os.File) {
	h := []byte(HeaderToSQL(sheet.Name, sheet.Header()))
	out := file.WriteString
	sheet.Foreach(func(row *[]string) {
		file.Write(h)
		out("VALUES (")
		n := len(*row)
		for i, col := range *row {
			if len(col) == 0 {
				out("NULL")
			} else {
				out("'")
				out(col)
				out("'")
			}
			if i+1 != n {
				out(",")
			}
		}
		out(");\n")
	})
}

// HeaderToSQL insert into [dbname] () values ()\n
func HeaderToSQL(name string, header []string) string {
	buf := bytes.Buffer{}
	out := buf.WriteString
	out("INSERT INTO ")
	out(name)
	out(" (")
	n := len(header)
	for i, col := range header {
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
	ss := InitSheets()
	n := len(ss)
	SheetsToSQL(done, ss)
	for i := 0; i < n; i++ {
		<-done
	}
}
