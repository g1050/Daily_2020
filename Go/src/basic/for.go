package main

import "fmt"

func main() {
	for i := 58; i >= 23; i-- {
		fmt.Println(i)
	}

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}

	for i := 1; i <= 9; i++ { //jetbrains-license-server
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %2d ", j, i, i*j)
		}
		fmt.Println("")
	}
}
