package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Sheet to insert sql
type Sheet interface {
	SheetName() string
	FileName() string
	FirstLine() []string
	Foreach(func(*[]string))
	Length() int
}

// JsonSheet 二维字符串数组即为一张表
type JsonSheet struct {
	TableName string                   // 表名
	fileName  string                   // 表名
	SqlHead   []string                 // sql表头
	JsonHead  []string                 // json表头
	Data      []map[string]interface{} // 核心数据
}

// SheetName 获得表头名 [table].json
func (s *JsonSheet) SheetName() string {
	return s.TableName
}

// FileName 获得表头名 [table].json
func (s *JsonSheet) FileName() string {
	return s.fileName
}

// FirstLine 获得表头列表
func (s *JsonSheet) FirstLine() []string {
	return s.SqlHead
}

// Foreach 遍历[1:n]
func (s *JsonSheet) Foreach(f func(*[]string)) {
	for i, n := 0, len(s.Data); i < n; i++ {
		line := s.Data[i]
		data := make([]string, len(s.JsonHead))
		for j, H := range s.JsonHead {
			var v interface{} = line
			hs := strings.Split(H, ".")
			for _, h := range hs {
				_h, err := strconv.Atoi(h)
				if err == nil { // h is number
					v = v.([]interface{})[_h]
				} else {
					v = v.(map[string]interface{})[h]
				}
			}
			data[j] = fmt.Sprint(v)
		}
		f(&data)
	}
}

// Length size-1
func (s *JsonSheet) Length() int {
	return len(s.Data) - 1
}

// NewJsonSheet new(Sheet)
func NewJsonSheet(TableName, FileName string, Head []map[string]string, Data []map[string]interface{}) Sheet {
	SqlHead := make([]string, len(Head))
	JsonHead := make([]string, len(Head))
	for i, H := range Head {
		for json, sql := range H {
			SqlHead[i] = sql
			JsonHead[i] = json
		}
	}
	return &JsonSheet{TableName, FileName, SqlHead, JsonHead, Data}
}

func parseJson(file string, v interface{}) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// InitSheets 初始化 命令行参数
func InitSheets() []Sheet {
	if len(os.Args) <= 1 {
		os.Exit(1)
	}
	header := make([]map[string]string, 0)
	err := parseJson(os.Args[1], &header)
	if err != nil {
		println("第", 1, "个文件必须是Array<String>格式，作为表头")
		os.Exit(1)
	}
	sheets := make([]Sheet, 0)
	tableName := strings.TrimSuffix(filepath.Base(os.Args[1]), filepath.Ext(os.Args[1]))
	for i := 2; i < len(os.Args); i++ {
		content := make([]map[string]interface{}, 0)
		err := parseJson(os.Args[i], &content)
		if err != nil {
			println("第", i-1, "个文件必须是Map<String, Any>格式，作为表头")
			os.Exit(1)
		}
		fileName := strings.TrimSuffix(filepath.Base(os.Args[i]), filepath.Ext(os.Args[i]))
		sheets = append(sheets, NewJsonSheet(tableName, fileName, header, content))
	}
	return sheets
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
			sqlF, err := os.Create(sheet.FileName() + ".sql")
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
