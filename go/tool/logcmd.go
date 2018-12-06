//go build -ldflags "-s -w" logcmd.go && logcmd.exe && go clean
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const LogcmdRemark = `record the execution result of the command by time
	logcmd execute shell:
    logcmd [shell] [output] [time(y|M|d|h|m|s)]

Example:
    logcmd "python Test.py" py.log M
    logcmd xxx.exe exe.log d
    logcmd ls ls.log m`

type LogOuter struct {
	TimeLayout  string
	LogFileName string
}

func (o *LogOuter) print(line string) {
}

func invoke(cmd *exec.Cmd) {
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
		fmt.Println(LogcmdRemark)
		os.Exit(-1)
		return time.Duration(0)
	}
}

func initParameter() (*exec.Cmd, *LogOuter) {
	args := os.Args
	if len(args) < 3 {
		fmt.Println(LogcmdRemark)
		os.Exit(-1)
	}
	sh := strings.Split(args[1], " ")
	return exec.Command(sh[0], sh[1:]...), nil
}

func main() {
	cmd, out := initParameter()
	invoke(cmd)
	print(out)
}
