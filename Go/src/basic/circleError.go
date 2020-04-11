package main

import "fmt"
import "math"

func main(){
    radius := float64(-10)
    s,err := getCircleArea(radius)
    if err != nil{
        fmt.Println(err)
    }else{
        fmt.Println("面积是",s)
    }


}

type circleAreaError struct{//自定义错误类型
    msg string //两个字段,一个错误信息，一个半径
    radius float64
}

// 实现error接口返回一个字符串
func (c *circleAreaError)Error()string{
    return fmt.Sprintf("发生错误 半径:%f,%s",c.radius,c.msg)//生成格式化字符串
}

//返回错误接口,error是指针类型的
func getCircleArea(radius float64)(float64,error){
    if radius < 0{
        return 0,&circleAreaError{msg:"半径是非法的",radius:radius}
    }else{
        return math.Pi*radius*radius,nil
    }
}


