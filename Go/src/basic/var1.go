package main

import "fmt"

func main(){
    var num1 int
    num1 = 30
    fmt.Printf("%d",num1);

    var num2 int = 15
    fmt.Printf("%d\n",num2);

    //类型推断
    var name = "wanghuahua"
    fmt.Printf("类型是%T,内容是%s\n",name,name);

    //简短声明,省略var,采用:
    num := 100
    fmt.Printf("%d\n",num);

    var m,n int = 100,200
    var p,q = 100,'a'
    fmt.Printf("%d %d\n",m,n)
    fmt.Printf("%d %c\n",p,q)

    var (
        studentname = "Jack"
        studentage = 22
    )
    fmt.Printf("%s %d",studentname,studentage)

    var num3 = 100
    fmt.Printf("%d %p\nn",num3,&num3)
    // fmt.Println("Hello, World!")
}

