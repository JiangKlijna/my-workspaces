//go build -ldflags "-s -w" recmd.go && recmd.exe && go clean
package main

import (
	"io"
	"os"
	"fmt"
	"time"
	"bufio"
	"strings"
	"strconv"
	"os/exec"
)

const remark = `repeat execute shell:
    recmd [shell] [time(h|m|s)]

Example:
    recmd "python Test.py" 10m
    recmd xxx.exe 24h
    recmd ls 3s`

func invoke(sh []string) {
	//CombinedOutput
	cmd := exec.Command(sh[0], sh[1:]...)
	fmt.Println(sh)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Print(line)
	}
	fmt.Println()
	cmd.Wait()
}

func parseTime(s string) time.Duration {
	n := len(s) - 1
	i, _ := strconv.ParseInt(s[0:n], 10, 64)
	switch s[n] {
	case 'h':
		return time.Duration(i) * time.Hour
	case 'm':
		return time.Duration(i) * time.Minute
	case 's':
		return time.Duration(i) * time.Second
	default:
		fmt.Println(remark)
		os.Exit(-1)
		return time.Duration(0)
	}
}

func initParameter() ([]string, time.Duration) {
	args := os.Args
	if len(args) < 3 {
		fmt.Println(remark)
		os.Exit(-1)
	}
	return strings.Split(args[1], " "), parseTime(args[2])
}

func main() {
	cmd, dur := initParameter()
	for i := 0; true; i++ {
		go invoke(cmd)
		time.Sleep(dur)
	}
}
