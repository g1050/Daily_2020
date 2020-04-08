package main

import "fmt"

func main(){
    map1 := make(map[string]int)
    
    map1["Jack"] = 98
    map1["Mary"] = 99
    map1["Tom"] = 97

    fmt.Println(map1)

    // delete(map1,"Jack")
    value,ok := map1["Jack"]
    if ok{
        fmt.Println("存在值,值是",value)
    }else{
        fmt.Println("不存在")
    }

    for key,value := range map1{
        fmt.Println(key,"----->",value)
    }

}

