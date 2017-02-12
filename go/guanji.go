package main

import (
    "fmt"
    "os/exec"
    "strconv"
)

func main()  {
    var tm_s string
    var tm int

    for {
        fmt.Println("请输入时间(分钟数):")
        fmt.Scanf("%s", &tm_s)
        i, err := strconv.ParseInt(tm_s, 10, 64)
        if err != nil {
            fmt.Println("您输入时间有误，请重新输入")
        } else {
            tm = int(i)
            break
        }
    }
    sh := []string{"-s", "-t", strconv.Itoa(tm * 60)}
    re, err := exec.Command("shutdown", sh...).CombinedOutput()
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println(string(re))
    }
    fmt.Printf("计算机将在%s分钟后关机", tm_s)

}
