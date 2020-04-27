package main

import (
    "fmt"
    "github.com/astaxie/beego/httplib"
    // "encoding/json"
)

func main()  {
    var obj interface{}
    // mp := make(map[string]interface{})

    req:=httplib.Post("http://192.168.0.115:8080/upload")
    req.PostFile("file","linjie.go")//注意不是全路径
    req.Param("output","json")
    req.Param("scene","default")
    req.Param("path","")

    req.ToJSON(&obj)
    // json.Unmarshal(obj.([]byte), &mp) 

    fmt.Println(obj)
    fmt.Println(obj.(map[string]interface{})["url"])


}
