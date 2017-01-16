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

var remark =
`repeat execute shell:
    recmd [shell] [time(h|m|s)]

Example:
    recmd "python Test.py" 10m
    recmd xxx.exe 24h
    recmd ls 3s`

var sh []string
var tm time.Duration

func cmd()  {
    /*re, err := exec.Command(sh[0], sh[1:]...).CombinedOutput()
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println(string(re))
    }*/
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
    n := len(s)-1
    i, _ := strconv.ParseInt(s[0:n], 10, 64)
    switch s[n] {
    case "h"[0]:
        return time.Duration(i)*time.Hour
    case "m"[0]:
        return time.Duration(i)*time.Minute
    case "s"[0]:
        return time.Duration(i)*time.Second
    default:
        fmt.Println(remark)
        os.Exit(-1)
        return time.Duration(0)
    }
}

func init()  {
    args := os.Args
    if len(args) < 3 {
        fmt.Println(remark)
        os.Exit(-1)
    }
    sh = strings.Split(args[1], " ")
    tm = parseTime(args[2])
}

func main()  {
    for i := 0; true; i++ {
        go cmd()
        time.Sleep(tm)
    }
}
