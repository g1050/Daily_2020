package main

import "fmt"

func main(){

    mouse := Mouse{name:"罗技"}
    fmt.Println(mouse.name)
    testUSB(mouse)

    flashdisk := FlashDisk{name:"金士顿"}
    fmt.Println(flashdisk.name)
    testUSB(flashdisk)
        
}

type USB interface{
    start(second int)
    end()int
}

type Mouse struct{
    name string
}

type FlashDisk struct{
    name string
}

func (mouse Mouse)start(second int){
    fmt.Printf("%d seconds later mouse will begin to work.\n",second)
}

func (mouse Mouse)end()int{
    hour:= 5
    fmt.Printf("Mouse has been worked for %d hours.\n",hour)
    return hour
}

func (flashdisk FlashDisk)start(second int){
    fmt.Printf("%d seconds later FlashDisk will begin to work.\n",second)
}

func (flashdisk FlashDisk)end()int{
    hour := 10
    fmt.Printf("FlashDisk has been worked for %d hours.\n",hour)
    return hour
}

func testUSB(usb USB){
    usb.start(100)

    res := usb.end()
    fmt.Println(res)

}
