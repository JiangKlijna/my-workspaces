package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/tealeg/xlsx/v3"
)

// 可以解析xlsx和csv两种文件,xlsx为多个Sheet,csv为一张Sheet
type Sheet interface {
	SheetName() string
	FirstLine() []string
	Foreach(func(*[]string))
	Length() int
}

// CsvSheet 二维字符串数组即为一张表
type CsvSheet struct {
	Name string     // 表名
	Data [][]string // 核心数据
}

// Name 获得表头名
func (s *CsvSheet) SheetName() string {
	return s.Name
}

// FirstLine 获得表头列表
func (s *CsvSheet) FirstLine() []string {
	return s.Data[0]
}

// Foreach 遍历[1:n]
func (s *CsvSheet) Foreach(f func(*[]string)) {
	for i, n := 1, len(s.Data); i < n; i++ {
		f(&s.Data[i])
	}
}

// Length size-1
func (s *CsvSheet) Length() int {
	return len(s.Data) - 1
}

// NewCsvSheet new(Sheet)
func NewCsvSheet(Name string, Data [][]string) Sheet {
	return &CsvSheet{Name, Data}
}

type XlsxSheet struct {
	*xlsx.Sheet
}

// Name 获得表头名
func (s *XlsxSheet) SheetName() string {
	return s.Name
}

// FirstLine 获得表头列表
func (s *XlsxSheet) FirstLine() []string {
	h := make([]string, 0)
	row, _ := s.Row(0)
	row.ForEachCell(func(c *xlsx.Cell) error {
		h = append(h, c.Value)
		return nil
	})
	return h
}

// Foreach 遍历[1:n]
func (s *XlsxSheet) Foreach(f func(*[]string)) {
	for i, n := 1, s.MaxRow; i < n; i++ {
		h := make([]string, 0)
		row, _ := s.Row(i)
		row.ForEachCell(func(c *xlsx.Cell) error {
			h = append(h, c.Value)
			return nil
		})
		f(&h)
	}
}

// Length size-1
func (s *XlsxSheet) Length() int {
	return s.MaxRow - 1
}

// NewCsvSheet new(Sheet)
func NewXlsxSheet(sheet *xlsx.Sheet) Sheet {
	return &XlsxSheet{sheet}
}

// InitSheets 初始化 命令行参数和excelize
func InitSheets() []Sheet {
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
		return []Sheet{NewCsvSheet(strings.TrimSuffix(filepath.Base(table.Name()), ".csv"), data)}
	}
	if strings.HasSuffix(os.Args[1], ".xlsx") {
		xlsx, err := xlsx.OpenFile(os.Args[1])
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
		ss := make([]Sheet, 0)

		for _, sheet := range xlsx.Sheets {
			ss = append(ss, NewXlsxSheet(sheet))
		}
		return ss
	}
	println("file ext is not [.csv, .xlsx]")
	os.Exit(1)
	return nil
}

// SheetsToSQL 整个excel文件处理,并返回共有几张表
func SheetsToSQL(done chan bool, sheets []Sheet) {
	for _, s := range sheets {
		go func(sheet Sheet) {
			time.Sleep(time.Second)
			if sheet.Length() < 1 {
				println(sheet.SheetName(), "is empty")
				done <- false
				return
			}
			sqlF, err := os.Create(sheet.SheetName() + ".sql")
			if err != nil {
				println(sheet.SheetName()+".sql", "create error", err.Error())
				done <- false
				return
			}
			SheetToSQL(sheet, sqlF)
			done <- true
		}(s)
	}
}

// SheetToSQL 针对一张表的sql文件写入
func SheetToSQL(sheet Sheet, file *os.File) {
	header, template := SplitHeader(sheet.FirstLine())
	h := []byte(HeaderToSQL(sheet.SheetName(), header))
	out := file.WriteString
	sheet.Foreach(func(row *[]string) {
		file.Write(h)
		out("VALUES (")
		n := len(*row)
		for i, col := range *row {
			if len(col) == 0 {
				out("NULL")
			} else {
				if "" == template[i] {
					out("'")
					out(col)
					out("'")
				} else {
					out(strings.ReplaceAll(template[i], header[i], col))
				}
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
		out("\"")
		out(col)
		out("\"")
		if i+1 != n {
			out(",")
		}
	}
	out(") ")
	return buf.String()
}

func SplitHeader(oneline []string) ([]string, []string) {
	header := make([]string, len(oneline))
	template := make([]string, len(oneline))
	for i, v := range oneline {
		arr := strings.Split(v, "#")
		if len(arr) == 1 {
			header[i] = v
			template[i] = ""
		} else {
			header[i] = arr[0]
			template[i] = arr[1]
		}
	}
	return header, template
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
