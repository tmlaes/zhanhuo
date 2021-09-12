package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"zhanhuo/adb"
	"zhanhuo/entity"
	"zhanhuo/process"
)

func main() {
	fmt.Println("类型：1：创建账号 \t 2：升级账号")
	in := bufio.NewReader(os.Stdin)
	str, _, err := in.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	types, _ := strconv.Atoi(strings.TrimSpace(string(str)))
	if types == 1 {
		process.Create()
	} else {
		for i := 0; i < 5; i++ {
			process.Process()
			<-time.After(time.Duration(1) * time.Hour)
		}
	}
	//test()
}

func test() {
	devices := adb.Devices()
	//emulator-5556
	fmt.Println(devices)
	//adb.Shake("1")
	//for {
	//	adb.ScreenSeven(devices[0])
	//adb.ScreenTwo(devices[0])
	//adb.ScreenShot("ScreenOne.png", devices[0])
	//shotImg := adb.Screenshot(devices[0])
	//img := entity.TaskGoImg(devices[0])
	//adb.CutImage(shotImg, img)
	//adb.ClickPoint(entity.Close2,0,devices[0])
	//text1 := process.GetText1(devices[0], P.Img7(devices[0]))
	//fmt.Println(text1)
	if adb.Compare1(entity.ShiBingImg(devices[0]), devices[0]) {
		fmt.Println("匹配成功")
	}

	//shotImg := adb.Screenshot(devices[0])
	//text := process.GetText(shotImg, entity.Task1Img(devices[0]))
	//fmt.Println(text)

	//img := &entity.OcrImg{
	//	Name: devices[0] + "task.png",
	//	Path:"",
	//	X:220, Y:550, W:80, H:30,
	//	RGB: [3][2]uint8{{28,40},{40,50},{40,55}},
	//}
	//
	//text := process.GetText1(devices[0], img)
	//text := process.GetText1(devices[0], entity.Task1Img(devices[0]))
	//fmt.Println(text)

}
