//go build -ldflags "-s -w" compile.go
package main

import (
	"net/url"
	"os"
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const info = `run a single code file
compile [language] [filename]

Example:
	compile java Test.java
	compile swift hello.swift

Language:
	php
	python
	java
	c
	c++
	ruby
	c#
	scala
	objc
	perl
	bash
	vb.net
	swift
	go
	node
	lua
	pascal
	kotlin
`

type Result struct {
	Output string `json:"output"`
	Langid int    `json:"langid"`
	Code   string `json:"code"`
	Errors string `json:"errors"`
	Time   string `json:"time"`
}

func (self *Result) String() string {
	str, err := json.Marshal(*self)
	if err != nil {
		return err.Error()
	} else {
		return string(str)
	}
}

type Compile struct {
	langStr string
	langNum string
}

var compiles = map[string]*Compile{
	"php":    &Compile{"php", "3"},
	"python": &Compile{"python3", "15"},
	"java":   &Compile{"java", "8"},
	"c":      &Compile{"c", "7"},
	"c++":    &Compile{"c++", "7"},
	"ruby":   &Compile{"ruby", "1"},
	"c#":     &Compile{"c#", "10"},
	"scala":  &Compile{"scala", "5"},
	"objc":   &Compile{"objc", "12"},
	"perl":   &Compile{"perl", "14"},
	"bash":   &Compile{"bash", "11"},
	"vb.net": &Compile{"vb.net", "9"},
	"swift":  &Compile{"swift", "16"},
	"go":     &Compile{"go", "6"},
	"node":   &Compile{"node", "4"},
	"lua":    &Compile{"lua", "17"},
	"pascal": &Compile{"pascal", "18"},
	"kotlin": &Compile{"kotlin", "19"},
}

func (self *Compile) Params() ([]byte, error) {
	filename := os.Args[2]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("File not found :", filename)
		return []byte{}, err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("File read failed :", filename)
		return data, err
	}
	buf := bytes.Buffer{}
	buf.WriteString("language=")
	buf.WriteString(self.langNum)
	buf.WriteString("&code=")
	buf.WriteString(url.QueryEscape(string(data)))
	return buf.Bytes(), nil
}

func (self *Compile) Request(params []byte) (*http.Response, error) {
	req, _ := http.NewRequest("POST", "https://m.runoob.com/api/compile.php", bytes.NewReader(params))
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Host", "m.runoob.com")
	req.Header.Add("Origin", "https://c.runoob.com")
	req.Header.Add("Referer", "https://c.runoob.com/compile/"+self.langNum)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("request error", err.Error())
	}
	return resp, err
}

func (self *Compile) Analyse(resp *http.Response) (*Result, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("server error", err.Error())
		return nil, err
	}
	var res Result
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("result error", err.Error())
		return nil, err
	}
	return &res, nil
}

func (self *Compile) Run() int {
	params, err := self.Params()
	if err != nil {
		return 1
	}
	reps, err := self.Request(params)
	if err != nil {
		return 2
	}
	res, err := self.Analyse(reps)
	if err != nil {
		return 3
	}
	if res.Output == "" {
		fmt.Print(res.Errors)
	} else {
		fmt.Print(res.Output)
	}
	return 0
}

func getCompile() *Compile {
	n := len(os.Args)
	if n != 3 {
		fmt.Print(info)
		os.Exit(-1)
	}
	compile, isNil := compiles[os.Args[1]]
	if !isNil {
		fmt.Print(info)
		os.Exit(-2)
	}
	return compile
}

func main() {
	compile := getCompile()
	os.Exit(compile.Run())
}
