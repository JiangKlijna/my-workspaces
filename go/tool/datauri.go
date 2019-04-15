package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/base64"
)

const headerString string = "data:image/png;base64,"

func convert(filenama string) string {
    bytes, err := ioutil.ReadFile(filenama)
    if err != nil {
        fmt.Print(err)
    }
	encodeString := base64.StdEncoding.EncodeToString(bytes)
	return headerString + encodeString
}

func main() {
	if len(os.Args) < 2 {
		println("file is null")
		os.Exit(1)
	}
	datauri := convert(os.Args[1]);
	fmt.Println(datauri)
}