package main

import (
    "fmt"
    "errors"
)

func main(){
    err := errors.New("Test error")
    fmt.Println(err)

    err2 := fmt.Errorf("error: %d\n",404)
    fmt.Println(err2)
}

