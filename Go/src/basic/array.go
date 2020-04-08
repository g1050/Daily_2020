package main

import "fmt"

func main(){
    var arr[5] int

    for i := 0;i<4;i++{
        arr[i] = i
    }

        
    for i := 0;i<4;i++{
        fmt.Println(arr[i])
    }

    fmt.Println("len: ",len(arr))
    fmt.Println("capacity: ",cap(arr))

    for index,value := range arr{
        fmt.Printf("下标是%d,值是:%d\n",index,value)
    }

    var arr2[3][4] int 
    
    for i := 0;i<len(arr2);i++{
        for j := 0;j<len(arr2[i]);j++{
            arr2[i][j] = i*j
        }
    }

    fmt.Println(len(arr2),len(arr2[0]))

    for _,arr := range arr2{
        for _,value := range arr{
            fmt.Print(value)
        }
        fmt.Println()
    } 
        
}

