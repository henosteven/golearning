package main

import (
    "fmt"
    "reflect"
)

func main() {
    funcs := map[string]interface{} {
        "foo" : foo,
        "bar" : bar,
    }

    call(funcs, "foo", 1)
    call(funcs, "bar", 1, 2)
}

func call(funcs map[string]interface{}, funcname string, params ... interface{}) {
    f := reflect.ValueOf(funcs[funcname])
    if len(params) != f.Type().NumIn() {
        panic("invalid params")
    }

    in := make([]reflect.Value, len(params))
    for k, param := range params {
        in[k] = reflect.ValueOf(param)
    }
    f.Call(in)
}

func foo(a int) {
    fmt.Println(a)
}

func bar(a, b int) {
    fmt.Println(a, b)
}
