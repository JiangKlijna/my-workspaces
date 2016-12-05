//go build -ldflags "-s -w" compile.go
package main

import (
  "os"
  "fmt"
  "strings"
	"net/http"
  "io/ioutil"
  "encoding/json"
)

var info = `
  compile [language] [filename]

  Example:
    compile java Test.java
    compile swift hello.swift

  Supported language:
    php
    pascal
    lua
    node.js
    go
    swift
    vb.net
    perl
    objectiveC
    scala
    c#
    ruby
    c++
    c
    java
    python
`

var langs_str = `{"php":"3","pascal":"18","lua":"17","node.js":"4","go":"6","swift":"16","vb.net":"9","perl":"14","objectiveC":"12","scala":"5","c#":"10","ruby":"1","c++":"7","c":"7","java":"8","python":"15"}`
var langs map[string]string

type Compile struct {
  lang string
}

//
func NewCompile() *Compile {
  lang := test_args()
  return &Compile{lang}
}

func (self *Compile) Params() (string, error) {
  filename := os.Args[2]
  f, err := os.Open(filename)
  if err != nil {
    fmt.Println("File not found : " + filename)
    return filename, err
  }
  data, err := ioutil.ReadAll(f)
  if err != nil {
    fmt.Println("File read failed : " + filename)
    return filename, err
  }
  return "language=" + self.lang + "&code=" + string(data), nil
}

func (self *Compile) Respose(params string) (*http.Response, error) {
  req, _ := http.NewRequest("POST", "http://tool.runoob.com/compile.php", strings.NewReader(params))
  req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko")
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
  req.Header.Add("Host", "tool.runoob.com")
  req.Header.Add("Origin", "http://c.runoob.com")
  req.Header.Add("Referer", "http://c.runoob.com/compile/" + self.lang)

  resp, err := (&http.Client{}).Do(req)
  if err != nil {
    fmt.Println("Network connection is failed")
  }
  return resp, err
}

func (self *Compile) Analysi(resp *http.Response) (*map[string]string, error) {
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println("Server error 500")
    return nil, err
  }
  // fmt.Println(string(body))

  var res map[string]string
  json.Unmarshal(body, &res)

  return &res, nil
}

func (self *Compile) Run() int {
  params, err := self.Params()
  if err != nil {
    return 1
  }
  reps, err := self.Respose(params)
  if err != nil {
    return 2
  }
  res, err := self.Analysi(reps)
  if err != nil {
    return 3
  }
  if (*res)["output"] == "" {
    fmt.Print((*res)["errors"])
  } else {
    fmt.Print((*res)["output"])
  }
  return 0
}

func init() {
  json.Unmarshal([]byte(langs_str), &langs)
}

func main() {
  os.Exit(NewCompile().Run())
}

func test_args() string {
  if len(os.Args) < 3 {
    fmt.Print(info)
    os.Exit(-1)
  }
  lang, isNil := langs[os.Args[1]]
  if !isNil {
    fmt.Print(info)
    os.Exit(-2)
  }
  return lang
}
