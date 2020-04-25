package main

import "fmt"

type User struct{
    password string
}
func main(){
    mp := make(map[string]interface{})
    user := User{}
    
    mp["password"] = "123456"
    user.password = mp["password"].(string)
    fmt.Printf("%T\n",mp["password"])
    fmt.Println(user)
}

