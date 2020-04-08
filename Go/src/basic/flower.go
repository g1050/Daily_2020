package main

import "fmt"
import "math"

func main(){
    for i := 100;i<=999;i++{
        a := i%10
        c := i/100
        b := i/10%10
        if float64(i) == math.Pow(float64(a),3) + math.Pow(float64(b),3) + math.Pow(float64(c),3){
           fmt.Println(i) 
        }
    }
}


