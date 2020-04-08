package main

import "fmt"

func main(){
    add(1,2)
    add2(1,2,3)
    var n int
    fmt.Scanf("%d",&n)
    fmt.Println(sum(n))

    c,s := rec(2,3)
    fmt.Println(c,s)

    fmt.Println(getSun(n))
}

func sum(n int) int{

    var num int
    for i := 1;i<=n;i++{
        num += i;
    }
    return num
}

func add(a, b int){
    fmt.Println(a+b)
}

func add2(nums...int){
    for _,value := range nums{
        fmt.Println(value)
    }
}

func rec(a,b int)(int,int){
    return (a+b)*2,a*b
}


func getSun(i int) (int){
    if(i == 0){
        return i
    }else{
        return i+getSun(i-1)
    }
}
