package main

import (
    "fmt"
)

func main() {
    var i interface{}
    i = 20
    v, ok := i.(string)
    fmt.Println(v, ok) // false 因为 i 为int
}
