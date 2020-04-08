package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime"
	"os"
	"path/filepath"
)

const DataUriRemark = `datauri execute shell:
	datauri [file]

Example:
	datauri 1.png`

func getFileName() string {
	if len(os.Args) < 2 {
		println(DataUriRemark)
		os.Exit(1)
	}
	return os.Args[1]
}

func getBase64ByFilename(filenama string) string {
	bytes, err := ioutil.ReadFile(filenama)
	if err != nil {
		print(err.Error())
	}
	return base64.StdEncoding.EncodeToString(bytes)
}

func getContentTypeByFilename(filenama string) string {
	ct := mime.TypeByExtension(filepath.Ext(filenama))
	if ct == "" {
		return "text/plain"
	}
	return ct
}

func main() {
	fn := getFileName()
	fmt.Printf("data:%s;base64,%s\n", getContentTypeByFilename(fn), getBase64ByFilename(fn))
}
