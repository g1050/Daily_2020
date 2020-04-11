package main

import "fmt"
import "math"

func main(){
    circle := Circle{10}
    rec := Rec{3,4}
    
    testInterface(circle)
    testInterface(rec)
    
    testInterface2(circle)
    testInterface2(rec)
}

func testInterface2(shape Shape){
    switch ins := shape.(type){
    case Circle:
        fmt.Printf("It is a circle,it's radius is %f\n",ins.radius)
    case Rec:
        fmt.Printf("It is a rectangle,is's length is %f,it's width is %f\n",ins.a,ins.b)
    }
}

func testInterface(shape Shape){

    fmt.Printf("Area is %f\n",shape.getS())
    fmt.Printf("Peri is %f\n",shape.getC())

    if ins,ok := shape.(Circle);ok{
        fmt.Printf("It is a circle,it's radius is %f\n",ins.radius)
    }else if ins,ok := shape.(Rec);ok{
        fmt.Printf("It is a rectangle,is's length is %f,it's width is %f\n",ins.a,ins.b)
    }else{
        fmt.Println("啥都不是\n")
    }

}

//定义一个接口
type Shape interface{
    getC() float64
    getS() float64
}

type Circle struct{
    radius float64
}

type Rec struct{
    a float64
    b float64
}

func (circle Circle)getC()float64{
    return 2*math.Pi*circle.radius
}

func (circle Circle)getS()float64{
    return math.Pi*circle.radius*circle.radius
}


func (rec Rec)getC()float64{
    return 2*(rec.a+rec.b)
}

func (rec Rec)getS()float64{
    return rec.a*rec.b
}

