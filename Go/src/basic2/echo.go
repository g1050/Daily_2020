package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	var s,sep string;
	for i := 1;i<len(os.Args);i++ {
		s += sep + os.Args[i] //0是命令本身
		sep = " " //空格隔开
	}
	fmt.Println(s)

	i := 5
	for i > 0 {
		fmt.Println(i)
		i--
	}

	//for range 遍历
	for index,arg := range os.Args[1:] {
		fmt.Println(index,"  ",arg)
	}

	//string.join函数
	fmt.Println(os.Args[1:])
	fmt.Println(strings.Join(os.Args[1:]," "))
}
