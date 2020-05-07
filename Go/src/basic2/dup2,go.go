package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//命令行参数加文件名字
	counts := make(map[string]int)
	files := os.Args[1:]
	fmt.Println(files)

	//没有文件名字，就从标准输入中读取数据
	if len(files) == 0{
		countLines(os.Stdin,counts)
	}else {
		for index,arg := range files {
			fmt.Println("第",index,"个文件名字是",arg)
			f,err := os.Open(arg)
			defer f.Close()
			if err != nil {
				fmt.Println(err)
				continue
			}
			//countLines(f,counts)
			fmt.Println("------------------")
			countLines2(f,counts)
		}
	}

	//因为是引用传递所以此处可以便利输出
	for line,n := range counts {
		if n >= 1 {
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}

//版本3采用split和readfile
func countLines2(f *os.File,counts map[string]int) {
	//借助ioutil读取到buf中

	buf,err:= ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	//切分成数组，然后遍历
	arr := strings.Split(string(buf),"\n")
	for index,line:= range arr {
		counts[line] = index+1
	}
}
//map是引用传递
func countLines(f * os.File,counts map[string]int) {
	input := bufio.NewScanner(f)
	num := 1
	for input.Scan() {
		counts[input.Text()] = num
		num++
	}

}
