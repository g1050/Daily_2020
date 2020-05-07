package main

import (
    "fmt"
    "github.com/astaxie/beego/httplib"
    "io/ioutil"
    "log"

    // "encoding/json"
)

func main()  {
     var obj interface{}

     req:=httplib.Post("http://192.168.0.115:8080/")
     // req.PostFile("file","username")//注意不是全路径
    bt,err:=ioutil.ReadFile("gg.txt")
    if err!=nil{
        log.Fatal("read file err:",err)
    }
    fmt.Printf("%s",bt)
    req.Body(bt)

     req.Param("output","json")
     req.Param("scene","default")
     req.Param("path","")
    req.ToJSON(&obj)
    fmt.Println(obj)
    //fmt.Println(obj.(map[string]interface{})["url"])
    fmt.Println(req.String())


}
