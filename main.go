package main

import (
	"fmt"
	"time"
	"zhanhuo/adb"
	"zhanhuo/entity"
	"zhanhuo/process"
)

func main() {
	//fmt.Println("脚本开始......")
	//阻止主goroutine 退出
	//signalChan := make(chan os.Signal, 1)
	//捕捉 Ctrl+c 和 kill 信号，写入signalChan
	//signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	// 此处执行处理逻辑
	//process.Create()
	for i := 0; i < 5; i++ {
		process.Process()
		<-time.After(time.Duration(1) * time.Hour)
	}

	// signalChan阻塞进程
	//<-signalChan

	// 捕捉信号后在Exit函数中处理信息，例如内存持久化等信息防止丢失
	//nsqd.Exit()
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
	shotImg := adb.Screenshot(devices[0])
	img := entity.JiangLiTaskImg(devices[0])
	process.CutImage(shotImg, img)

	//text1 := process.GetText1(devices[0], P.Img7(devices[0]))
	//fmt.Println(text1)
	/*if process.Compare1(entity.ShiBingImg(devices[0]), devices[0]) {
		fmt.Println("匹配成功")
	}*/

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
