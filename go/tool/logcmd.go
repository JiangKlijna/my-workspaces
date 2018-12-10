//go build -ldflags "-s -w" logcmd.go && logcmd.exe && go clean
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

const LogcmdRemark = `record the execution result of the command by time
	logcmd execute shell:
    logcmd [shell] [output] [time(y|M|d|h|m|s)]

Example:
    logcmd "python Test.py" file M
    logcmd xxx.exe fn d
    logcmd ls f m`

var timeLayouts = map[string]string{
	"y": "2006",
	"M": "2006-01",
	"d": "2006-01-02",
	"h": "2006-01-02+15",
	"m": "2006-01-02+15:04",
	"s": "2006-01-02+15:04:05",
}

type LogOuter struct {
	TimeLayout  string
	LogFileName string
}

func (o *LogOuter) print(line string, time time.Time) {
	filename := fmt.Sprintf("%s.%s.log", o.LogFileName, time.Format(o.TimeLayout))
	err := ioutil.WriteFile(filename, []byte(line), os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}
}

func (o *LogOuter) invoke(cmd *exec.Cmd) {
	//CombinedOutput
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	for {
		reader.ReadLine()
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		go o.print(line, time.Now())
	}
	cmd.Wait()
}

func initParameter() (*LogOuter, *exec.Cmd) {
	args := os.Args
	if len(args) < 4 {
		fmt.Println(LogcmdRemark)
		os.Exit(1)
	}
	tl, ok := timeLayouts[args[3]]
	if !ok {
		fmt.Println(LogcmdRemark)
		os.Exit(2)
	}
	sh := strings.Split(args[1], " ")
	return &LogOuter{tl, args[2]}, exec.Command(sh[0], sh[1:]...)
}

func main() {
	out, cmd := initParameter()
	out.invoke(cmd)
}
