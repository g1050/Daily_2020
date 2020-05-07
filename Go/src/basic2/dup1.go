package main

import (
	"bufio"
	"fmt"
	"os"
)

func  main()  {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	num := 1
	for input.Scan() {
		text := input.Text()
		if text == "end" {
			break
		}
		counts[text] = num
		num++
	}

	fmt.Println(counts)
	//遍历
	for line,n := range counts {
		if n >= 1 {
			fmt.Printf("%s %d\n",line,n)
		}

	}

}
