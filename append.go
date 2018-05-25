package main

import (
    "fmt"
)

func main() {
    s := []byte("")
    fmt.Println(cap(s), len(s)) //32, 0
    s1 := append(s, 'a')
    s2 := append(s, 'b')
    //fmt.Println(s1, "==========", s2)
    fmt.Println(string(s1), "==========", string(s2))
    // 出现个让我理解不了的现象, 注释时候输出是 b ========== b
    // 取消注释输出是 [97] ========== [98] a ========== b
}

