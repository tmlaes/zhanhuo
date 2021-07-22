package process

import (
	"fmt"
	"time"
	"zhanhuo/adb"
	P "zhanhuo/entity"
)

func ZhaoMu(account *P.Account) {
	device := account.Name
	//队列1
	var status [4]int
	for i := 0; i < 4; i++ {
		shotImg := adb.Screenshot(device)
		img := P.Task1Img(device)
		img.Y = img.Y + 180 + 50*i
		text := GetText(shotImg, img)
		if i == 0 && text == "" {
			fmt.Println(time.Now(), device, "未读取到兵营状态", i)
			i--
			continue
		}
		fmt.Println(time.Now(), device, "兵营状态", text)
		if text == "空闲中" {
			status[i] = 1
		}
	}
	//关闭任务列表
	time.Sleep(time.Duration(1) * time.Second)
	adb.Click(P.Reset.X, P.Reset.Y, 1, device)
	for i, s := range status {
		if s == 0 {
			continue
		}
		if i == 0 {
			adb.ScreenTwo(device)
		} else {
			adb.ScreenOne(device)
		}
		var pp P.Point
		if i == 0 {
			pp = P.Bingyin4
		} else if i == 1 {
			pp = P.Bingyin3
		} else if i == 2 {
			pp = P.Bingyin1
		} else {
			pp = P.Bingyin2
		}
		time.Sleep(time.Duration(2) * time.Second)
		fmt.Println(time.Now(), device, "点击兵营")
		adb.Click(pp.X, pp.Y, 2, device)
		adb.Click(pp.X, pp.Y, 2, device)
		adb.Click(pp.X, pp.Y, 2, device)
		if i == 0 {
			adb.Click(P.FaShi.X, P.FaShi.Y, 2, device)
		} else {
			adb.Click(P.QiBin.X, P.QiBin.Y, 2, device)
		}
		//x+110 y+100
		fmt.Println(time.Now(), device, "点击招募按钮")
		adb.Click(pp.X+110, pp.Y+100, 3, device)
		adb.Click(P.Shengji.X, P.Shengji.Y, 2, device)
	}
}
