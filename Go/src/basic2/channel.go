package main

import "fmt"

func main(){
    var a chan int 
    fmt.Printf("%T %v\n",a,a)
    a = make(chan int)
    fmt.Printf("%T %v\n",a,a)//通道是一个引用类型的数据
    
    testchan(a)
}


func testchan(ch chan int){
    fmt.Printf("%T %v\n",ch,ch)//通道是一个引用类型的数据
}

