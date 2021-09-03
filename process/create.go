package process

import (
	"fmt"
	"sync"
	"time"
	"zhanhuo/adb"
	"zhanhuo/entity"
	"zhanhuo/utils"
)

var wg sync.WaitGroup

func Create() {
	fmt.Println("******************创建账号******************")
	deviceList := utils.ReadDevices()
	for _, list := range deviceList {
		//启动设备
		devices := adb.Devices_(list)
		//读取配置
		for device, index := range devices {
			wg.Add(1)
			go createOne(device, index)
		}
		wg.Wait()
	}

}

func createOne(device string, index string) {
	adb.ClickPoint(entity.Game, 3, device)
	adb.ClickMore(entity.P2, 6, 2, device)
	ticker := time.NewTicker(3 * time.Second)
	for range ticker.C {
		startText := utils.GetText(device, entity.StartImg(device))
		fmt.Println(utils.Now(), startText)
		if startText == "贝拉" {
			fmt.Println(utils.Now(), "已进入开始界面")
			ticker.Stop()
			break
		}
	}
	createTwo(index, device)
	if statuMap[device] {
		return
	}
	createThree(index, device)
	if statuMap[device] {
		return
	}
	createFour(index, device)
	if statuMap[device] {
		return
	}
	createFive(device)
	createSix(device)
	account := createAccount(device, index)
	YeGuai(account)
	Caiji(account)
	updateLastTime(account)
	adb.Quit(account.Id)
	utils.WriteJons(account)
	wg.Done()
}

func createTwo(id, device string) {
	delay(2)
	adb.ClickMore(entity.P1, 4, 2, device)
	delay(4)
	adb.ClickMore(entity.P1, 2, 2, device)
	adb.ClickPoint(entity.P3, 3, device)
	adb.ClickPoint(entity.P4, 3, device)
	adb.ClickPoint(entity.P5, 3, device)
	adb.ClickPoint(entity.P6, 3, device)
	adb.ClickPoint(entity.P7, 3, device)
	adb.ClickPoint(entity.P8, 5, device)
	adb.ClickMore(entity.P1, 2, 2, device)
	delay(11)
	adb.ClickMore(entity.P1, 2, 2, device)
	delay(3)
	adb.ClickPoint(entity.P9, 2, device)
	adb.ClickPoint(entity.P10, 3, device)
	adb.ClickPoint(entity.P11, 5, device)
	adb.ClickMore(entity.P1, 3, 2, device)
	delay(1)
	adb.ClickPoint(entity.P12, 2, device)
	adb.ClickPoint(entity.P10, 3, device)
	adb.ClickPoint(entity.P13, 2, device)
	adb.ClickPoint(entity.P14, 3, device)
	adb.ClickPoint(entity.P15, 3, device)
	adb.ClickPoint(entity.P16, 5, device)
	adb.ClickMore(entity.P1, 4, 2, device)
	delay(2)
	adb.ClickPoint(entity.P17, 2, device)
	adb.ClickPoint(entity.P18, 1, device)
	delay(6)
	adb.ClickMore(entity.P19, 3, 2, device)
	adb.ClickPoint(entity.P20, 2, device)
	adb.ClickPoint(entity.P21, 3, device)
	adb.ClickPoint(entity.P22, 1, device)
	adb.ClickPoint(entity.P23, 3, device)
	adb.ClickPoint(entity.P24, 1, device)
	adb.ClickPoint(entity.P25, 18, device)
	adb.ClickPoint(entity.P26, 8, device)
	adb.ClickMore(entity.P1, 3, 2, device)
	delay(3)
	adb.ClickPoint(entity.P27, 2, device)
	adb.ClickPoint(entity.P28, 3, device)
	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到城堡升级按钮，退出")
			return
		}
		img := entity.ChengBaoImg(device)
		if utils.Compare1(img, device) {
			break
		}
		adb.ClickPoint(entity.P27, 2, device)
		adb.ClickPoint(entity.P28, 3, device)
	}

	adb.ClickPoint(entity.P29, 5, device)
	adb.ClickMore(entity.P1, 3, 2, device)
	delay(6)
	adb.ClickMore(entity.P1, 3, 2, device)
	delay(3)
	adb.ClickPoint(entity.P30, 3, device)
	adb.ClickPoint(entity.P31, 3, device)
	adb.ClickPoint(entity.P32, 3, device)
	adb.ClickPoint(entity.P33, 3, device)
	adb.ClickPoint(entity.P29, 5, device)
	adb.ClickPoint(entity.P30, 3, device)
	adb.ClickPoint(entity.P31, 3, device)
	adb.ClickPoint(entity.P34, 3, device)
	adb.ClickPoint(entity.P31, 3, device)
	adb.ClickMore(entity.P1, 3, 2, device)
	delay(1)
	adb.ClickPoint(entity.P27, 2, device)
	adb.ClickPoint(entity.P35, 3, device)
	adb.ClickPoint(entity.P36, 3, device)
	adb.ClickPoint(entity.P34, 3, device)
	adb.ClickPoint(entity.P37, 2, device)
	adb.ClickPoint(entity.P36, 2, device)
	adb.ClickPoint(entity.P37, 2, device)
	adb.ClickPoint(entity.P36, 3, device)
	adb.ClickPoint(entity.P34, 3, device)
	adb.ClickPoint(entity.P37, 2, device)
	adb.ClickPoint(entity.P36, 4, device)
	check1(id, device)
	if statuMap[device] {
		return
	}
	adb.ClickPoint(entity.P37, 2, device)
	adb.ClickPoint(entity.P29, 3, device)
	adb.ClickPoint(entity.P30, 3, device)
	adb.ClickPoint(entity.P31, 3, device)
	adb.ClickPoint(entity.P34, 3, device)
	adb.ClickPoint(entity.P31, 3, device)
	res(device)
	adb.ClickPoint(entity.P40, 1, device)
	adb.ClickPoint(entity.P34, 3, device)
	adb.ScreenTwo(device)
	adb.ClickPoint(entity.CC, 2, device)
	adb.ClickPoint_(entity.CC, 3, device)
	adb.ClickPoint(entity.P29, 3, device)
}

func createThree(id, device string) {
	adb.ClickPoint(entity.Bingyin1, 3, device)
	adb.ClickPoint_(entity.Bingyin1, 3, device)
	adb.ClickMore(entity.P36, 3, 1, device)
	adb.ClickPoint(entity.Back, 3, device)
	adb.ScreenOne(device)
	adb.ClickPoint(entity.Bingyin2, 3, device)
	for i := 0; i < 6; i++ {
		if utils.Compare1(entity.BingYing0Img(device), device) {
			break
		}
		adb.ClickPoint(entity.Left, 2, device)
	}
	adb.ClickPoint(entity.P36, 1, device)
	adb.ClickPoint(entity.Bingyin2, 3, device)
	adb.ClickPoint_(entity.Bingyin2, 3, device)
	adb.ClickMore(entity.P36, 3, 1, device)
	adb.ClickPoint(entity.Back, 3, device)
	taskJiangLi(device)
	adb.ClickPoint(entity.P30, 3, device)
	adb.ClickMore(entity.P47, 6, 1, device)
	adb.ClickPoint(entity.P31, 2, device)
	adb.ClickPoint(entity.P34, 2, device)
	adb.ClickPoint(entity.P31, 2, device)
	adb.ClickPoint(entity.P34, 2, device)
	adb.ClickPoint(entity.Back, 3, device)
	zhaoMuProcess(entity.Bingyin1, 2, device, 0)
	zhaoMuProcess(entity.Bingyin2, 1, device, 0)
	adb.ClickPoint(entity.Cangku, 2, device)
	adb.ClickPoint_(entity.Cangku, 3, device)
	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到仓库升级按钮，退出")
			return
		}
		img := entity.CangKuImg(device)
		if utils.Compare1(img, device) {
			break
		}
		adb.ScreenOne(device)
		adb.ClickPoint(entity.Cangku, 2, device)
		adb.ClickPoint_(entity.Cangku, 3, device)
	}
	adb.ClickMore(entity.P36, 3, 1, device)
	adb.ClickPoint(entity.Back, 3, device)
	adb.ClickPoint(entity.ZhanZhen, 2, device)
	adb.ClickPoint_(entity.ZhanZhen, 3, device)
	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到战争广场升级按钮，退出")
			return
		}
		img := entity.ZhanZhengImg(device)
		if utils.Compare1(img, device) {
			break
		}
		adb.ScreenOne(device)
		adb.ClickPoint(entity.ZhanZhen, 2, device)
		adb.ClickPoint_(entity.ZhanZhen, 3, device)
	}
	adb.ClickMore(entity.P36, 3, 2, device)
	adb.ClickPoint(entity.Back, 3, device)
	adb.ClickPoint(entity.P48, 3, device)
	adb.ClickMore(entity.P49, 2, 2, device)
	adb.ClickPoint(entity.P50, 10, device)
	adb.ClickPoint(entity.Back, 3, device)
}

func createFour(id, device string) {
	adb.ScreenTwo(device)
	adb.ClickPoint(entity.YiYuan, 2, device)
	adb.ClickPoint_(entity.YiYuan, 2, device)

	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到医院升级按钮，退出")
			return
		}
		img := entity.YiYuanImg(device)
		if utils.Compare1(img, device) {
			break
		}
		adb.ScreenTwo(device)
		adb.ClickPoint(entity.YiYuan, 2, device)
		adb.ClickPoint_(entity.YiYuan, 2, device)
	}
	adb.ClickMore(entity.P36, 3, 1, device)
	adb.ClickPoint(entity.Back, 3, device)
	adb.ClickPoint(entity.XueYuan, 3, device)
	adb.ClickPoint(entity.P36, 3, device)
	adb.ClickPoint(entity.XueYuan, 2, device)
	adb.ClickPoint_(entity.XueYuan, 3, device)
	adb.ClickMore(entity.P36, 3, 1, device)
	adb.ClickPoint(entity.Back, 3, device)
	adb.ClickPoint(entity.ShiGuan, 3, device)
	adb.ClickPoint(entity.P36, 3, device)
	adb.ClickPoint(entity.ShiGuan, 2, device)
	adb.ClickPoint_(entity.ShiGuan, 3, device)
	adb.ClickMore(entity.P36, 3, 1, device)
	adb.ClickPoint(entity.Back, 3, device)
	adb.ClickPoint(entity.RongLian, 3, device)
	adb.ClickPoint(entity.P36, 3, device)
	adb.ClickPoint(entity.RongLian, 2, device)
	adb.ClickPoint_(entity.RongLian, 3, device)
	adb.ClickMore(entity.P36, 3, 1, device)
	adb.ClickPoint(entity.Back, 3, device)
	adb.ClickPoint(entity.P40, 1, device)
	adb.ClickPoint(entity.P34, 3, device)
	updateChengbao(id, device)
	taskJiangLi(device)
	adb.ScreenSeven(device)
	adb.ScreenSeven(device)
	adb.ClickPoint(entity.ZY18, 3, device)
	for i := 0; i < 5; i++ {
		img := entity.ShiBingImg(device)
		if utils.Compare1(img, device) {
			break
		}
		adb.ClickPoint(entity.Left, 2, device)
	}
	adb.ClickPoint(entity.P36, 2, device)
	adb.ClickPoint(entity.ZY18, 2, device)
	adb.ClickPoint_(entity.ZY18, 3, device)
	adb.ClickMore(entity.P36, 4, 1, device)
	adb.ClickPoint(entity.Back, 3, device)
	adb.ScreenFour(device)
	adb.ClickPoint(entity.ZhenCha, 2, device)
	adb.ClickPoint_(entity.ZhenCha, 3, device)
	adb.ClickMore(entity.P36, 3, 1, device)
	adb.ClickPoint(entity.Back, 3, device)
	adb.ScreenThree(device)
	adb.ClickPoint(entity.CL, 3, device)
	adb.ClickMore(entity.P65, 2, 2, device)
	//x+140    y+140
	for i := 0; i < 4; i++ {
		p := entity.Point{X: 70 + 140*i, Y: 910}
		fmt.Println(utils.Now(), device, "点击材料", p)
		adb.ClickPoint(p, 2, device)
	}
	adb.ClickPoint(entity.Back, 3, device)
}

func createFive(device string) {
	adb.ScreenFive(device)
	var arr [12]entity.Point
	arr[0] = entity.ZY
	arr[1] = entity.ZY2
	arr[2] = entity.ZY3
	arr[3] = entity.ZY4
	arr[4] = entity.ZY5
	arr[5] = entity.ZY6
	arr[6] = entity.ZY7
	arr[7] = entity.ZY8
	arr[8] = entity.ZY9
	arr[9] = entity.ZY10
	arr[10] = entity.ZY12
	arr[11] = entity.ZY11
	for i := 0; i < len(arr); i++ {
		adb.ClickPoint(arr[i], 2, device)
		adb.ClickPoint_(arr[i], 2, device)
		adb.ClickPoint(entity.P36, 1, device)
		adb.ClickPoint(entity.Back, 2, device)
	}
	adb.ScreenFive(device)
	var arr1 [5]entity.Point
	arr1[0] = entity.ZY14
	arr1[1] = entity.ZY16
	arr1[2] = entity.ZY17
	arr1[3] = entity.ZY13
	arr1[4] = entity.ZY15
	for i := 0; i < len(arr1); i++ {
		if i == 4 {
			adb.ScreenFive(device)
		}
		adb.ClickPoint(arr1[i], 2, device)
		adb.ClickPoint_(arr1[i], 2, device)
		adb.ClickPoint(entity.P36, 1, device)
		adb.ClickPoint(entity.Back, 2, device)
	}
	adb.ScreenFive(device)
	adb.ClickPoint(entity.ZYZ2, 2, device)
	adb.ClickPoint(entity.P38, 2, device)
	adb.ClickPoint(entity.ZY19, 2, device)
	adb.ClickPoint(entity.P36, 2, device)
	adb.ClickPoint(entity.ZY19, 2, device)
	adb.ClickPoint_(entity.ZY19, 2, device)
	adb.ClickMore(entity.P36, 4, 1, device)
	adb.ClickPoint(entity.Back, 3, device)
}

func createSix(device string) {
	adb.ClickPoint(entity.Reset, 5, device)
	adb.ClickPoint(entity.P55, 2, device)
	adb.ClickPoint(entity.Search, 2, device)
	var arr [5]entity.Point
	arr[0] = entity.NC
	arr[1] = entity.MC
	arr[2] = entity.SC
	arr[3] = entity.TC
	arr[4] = entity.ZC
	for _, p := range arr {
		adb.ClickPoint(p, 1, device)
		adb.ClickPoint(entity.ZYL, 2, device)
		adb.Key("67", device)
		adb.Text("4", device)
		adb.Key("4", device)
		adb.ClickPoint(entity.GoTO, 2, device)
	}
	adb.ClickPoint(entity.P55, 2, device)
	adb.ClickPoint(entity.Reset, 6, device)

}

func check1(id, device string) {
	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到用户协议，退出")
			return
		}
		img := entity.XieYiImg(device)
		if utils.Compare1(img, device) {
			adb.ClickPoint(entity.P39, 2, device)
			return
		}
		delay(2)
	}
}

func res(device string) {
	adb.ScreenFive(device)
	var arr [12]entity.Point
	arr[0] = entity.ZY
	arr[1] = entity.ZY2
	arr[2] = entity.ZY3
	arr[3] = entity.ZY4
	arr[4] = entity.ZY5
	arr[5] = entity.ZY6
	arr[6] = entity.ZY7
	arr[7] = entity.ZY8
	arr[8] = entity.ZY9
	arr[9] = entity.ZY10
	arr[10] = entity.ZY12
	arr[11] = entity.ZY11
	for i := 0; i < len(arr); i++ {
		if i == 3 {
			adb.ClickPoint(arr[i], 2, device)
			adb.ClickPoint_(arr[i], 2, device)
			adb.ClickMore(entity.P36, 3, 1, device)
			adb.ClickPoint(entity.Back, 2, device)
			continue
		}
		adb.ClickPoint(arr[i], 2, device)
		adb.ClickPoint(entity.P36, 2, device)
		adb.ClickPoint(arr[i], 2, device)
		adb.ClickPoint_(arr[i], 2, device)
		adb.ClickMore(entity.P36, 3, 1, device)
		adb.ClickPoint(entity.Back, 2, device)
	}
	adb.ScreenFive(device)
	adb.ClickPoint(entity.ZYZ1, 2, device)
	adb.ClickPoint(entity.P38, 2, device)
	var arr1 [5]entity.Point
	arr1[0] = entity.ZY14
	arr1[1] = entity.ZY16
	arr1[2] = entity.ZY17
	arr1[3] = entity.ZY13
	arr1[4] = entity.ZY15
	for i := 0; i < len(arr1); i++ {
		if i == 4 {
			adb.ScreenFive(device)
		}
		adb.ClickPoint(arr1[i], 2, device)
		adb.ClickPoint(entity.P36, 2, device)
		adb.ClickPoint(arr1[i], 2, device)
		adb.ClickPoint_(arr1[i], 2, device)
		adb.ClickMore(entity.P36, 3, 1, device)
		adb.ClickPoint(entity.Back, 2, device)
	}
}

func createAccount(device string, index string) *entity.Account {
	fmt.Println(utils.Now(), device, "创建账号")
	acc := &entity.Account{}
	acc.Id = index
	acc.Name = device
	acc.ChengBao = 6
	acc.ChengQiang = 5
	acc.CangKu = 4
	acc.ZhanZhen = 4
	acc.YiYuan = 4
	acc.XueYuan = 4
	acc.RongLian = 4
	acc.ShiGuan = 4
	acc.ZhenCha = 4
	acc.GuaiWu = 0
	acc.BingYing1 = 4
	acc.BingYing2 = 4
	acc.NongChang = 5
	acc.MuChang = 5
	acc.ShiBing = 5
	acc.JiJiu = 5
	return acc
}
