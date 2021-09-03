package process

import (
	"fmt"
	"sync"
	"time"
	"zhanhuo/adb"
	"zhanhuo/entity"
	"zhanhuo/utils"
)

/**
 * 创建账号
 */
var wg sync.WaitGroup

func Create() {
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
		startText := GetText(device, entity.StartImg(device))
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
	YeGuai(index, device, 0)
	account := createAccount(device, index)
	Caiji(account)
	account.BuildTime = getBuildTime(device)
	updateLastTime(account)
	adb.Quit(account.Id)
	utils.WriteJons(account)
	wg.Done()
}

func createTwo(id, device string) {
	delay(2)
	fmt.Println(utils.Now(), "开始贝拉对话")
	adb.ClickMore(entity.P1, 4, 2, device)
	delay(4)
	adb.ClickMore(entity.P1, 2, 2, device)
	fmt.Println(utils.Now(), "开始建造兵营")
	adb.ClickPoint(entity.P3, 3, device)
	fmt.Println(utils.Now(), "点击建造")
	adb.ClickPoint(entity.P4, 3, device)
	fmt.Println(utils.Now(), "开始招募士兵")
	adb.ClickPoint(entity.P5, 3, device)
	fmt.Println(utils.Now(), "点击主界面招募")
	adb.ClickPoint(entity.P6, 3, device)
	fmt.Println(utils.Now(), "点击招募界面招募")
	adb.ClickPoint(entity.P7, 3, device)
	fmt.Println(utils.Now(), "点击主界面招募完成图标")
	adb.ClickPoint(entity.P8, 5, device)
	fmt.Println(utils.Now(), "开始贝拉对话")
	adb.ClickMore(entity.P1, 2, 2, device)
	fmt.Println(utils.Now(), "开始攻击野怪")
	delay(11)
	fmt.Println(utils.Now(), "野怪攻击结束")
	fmt.Println(utils.Now(), "开始贝拉对话")
	adb.ClickMore(entity.P1, 2, 2, device)
	delay(3)
	fmt.Println(utils.Now(), "开始建造农场")
	adb.ClickPoint(entity.P9, 2, device)
	fmt.Println(utils.Now(), "点击建造农场")
	adb.ClickPoint(entity.P10, 3, device)
	fmt.Println(utils.Now(), "点击农场收集")
	adb.ClickPoint(entity.P11, 5, device)
	fmt.Println(utils.Now(), "开始贝拉对话")
	adb.ClickMore(entity.P1, 3, 2, device)
	delay(1)
	fmt.Println(utils.Now(), "开始建造医院")
	adb.ClickPoint(entity.P12, 2, device)
	fmt.Println(utils.Now(), "点击建造")
	adb.ClickPoint(entity.P10, 3, device)
	fmt.Println(utils.Now(), "点击治疗伤兵")
	adb.ClickPoint(entity.P13, 2, device)
	fmt.Println(utils.Now(), "点击治疗图标")
	adb.ClickPoint(entity.P14, 3, device)
	fmt.Println(utils.Now(), "点击治疗界面开始治疗")
	adb.ClickPoint(entity.P15, 3, device)
	fmt.Println(utils.Now(), "点击主界面治疗完成")
	adb.ClickPoint(entity.P16, 5, device)
	fmt.Println(utils.Now(), "开始贝拉对话")
	adb.ClickMore(entity.P1, 4, 2, device)
	delay(2)
	fmt.Println(utils.Now(), "点击无尽之塔")
	adb.ClickPoint(entity.P17, 2, device)
	fmt.Println(utils.Now(), "点击经典战役")
	adb.ClickPoint(entity.P18, 1, device)
	fmt.Println(utils.Now(), "进入战斗界面")
	delay(6)
	fmt.Println(utils.Now(), "点击对话")
	adb.ClickMore(entity.P19, 3, 2, device)
	fmt.Println(utils.Now(), "建造攻击塔")
	adb.ClickPoint(entity.P20, 2, device)
	fmt.Println(utils.Now(), "点击加速")
	adb.ClickPoint(entity.P21, 3, device)
	fmt.Println(utils.Now(), "建造第二个塔")
	adb.ClickPoint(entity.P22, 1, device)
	adb.ClickPoint(entity.P23, 3, device)
	fmt.Println(utils.Now(), "建造第三个塔")
	adb.ClickPoint(entity.P24, 1, device)
	adb.ClickPoint(entity.P25, 18, device)
	fmt.Println(utils.Now(), "点击结束界面好的")
	adb.ClickPoint(entity.P26, 8, device)
	fmt.Println(utils.Now(), "开始贝拉对话")
	adb.ClickMore(entity.P1, 3, 2, device)
	delay(3)
	fmt.Println(utils.Now(), "开始升级城堡")
	adb.ClickPoint(entity.P27, 2, device)
	fmt.Println(utils.Now(), "点击升级按钮")
	adb.ClickPoint(entity.P28, 3, device)
	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到城堡升级按钮，退出")
			return
		}
		img := entity.ChengBaoImg(device)
		if Compare1(img, device) {
			break
		}
		adb.ClickPoint(entity.P27, 2, device)
		adb.ClickPoint(entity.P28, 3, device)
	}

	fmt.Println(utils.Now(), "点击升级界面升级")
	adb.ClickPoint(entity.P29, 5, device)
	fmt.Println(utils.Now(), "开始贝拉对话")
	adb.ClickMore(entity.P1, 3, 2, device)
	delay(6)
	adb.ClickMore(entity.P1, 3, 2, device)
	delay(3)
	fmt.Println(utils.Now(), "开始王者之路")
	adb.ClickPoint(entity.P30, 3, device)
	fmt.Println(utils.Now(), "点击王者之路界面升级城墙")
	adb.ClickPoint(entity.P31, 3, device)
	fmt.Println(utils.Now(), "点击升级城墙")
	adb.ClickPoint(entity.P32, 3, device)
	fmt.Println(utils.Now(), "点击升级城墙按钮")
	adb.ClickPoint(entity.P33, 3, device)
	fmt.Println(utils.Now(), "点击升级界面升级按钮")
	adb.ClickPoint(entity.P29, 5, device)
	fmt.Println(utils.Now(), "打开王者之路界面")
	adb.ClickPoint(entity.P30, 3, device)
	fmt.Println(utils.Now(), "点击升级城墙领取奖励")
	adb.ClickPoint(entity.P31, 3, device)
	fmt.Println(utils.Now(), "点击空白区域")
	adb.ClickPoint(entity.P34, 3, device)
	fmt.Println(utils.Now(), "点击王者之路界面升级城堡")
	adb.ClickPoint(entity.P31, 3, device)
	fmt.Println(utils.Now(), "开始贝拉对话")
	adb.ClickMore(entity.P1, 3, 2, device)
	delay(1)
	fmt.Println(utils.Now(), "开始升级城堡")
	adb.ClickPoint(entity.P27, 2, device)
	fmt.Println(utils.Now(), "点击升级城堡按钮")
	adb.ClickPoint(entity.P35, 3, device)
	fmt.Println(utils.Now(), "点击免费按钮")
	adb.ClickPoint(entity.P36, 3, device)
	fmt.Println(utils.Now(), "点击空白区域")
	adb.ClickPoint(entity.P34, 3, device)
	fmt.Println(utils.Now(), "点击跳转按钮")
	adb.ClickPoint(entity.P37, 2, device)
	fmt.Println(utils.Now(), "点击免费按钮")
	adb.ClickPoint(entity.P36, 2, device)
	fmt.Println(utils.Now(), "点击跳转按钮")
	adb.ClickPoint(entity.P37, 2, device)
	fmt.Println(utils.Now(), "点击免费按钮")
	adb.ClickPoint(entity.P36, 3, device)
	fmt.Println(utils.Now(), "点击空白区域")
	adb.ClickPoint(entity.P34, 3, device)
	fmt.Println(utils.Now(), "点击跳转按钮")
	adb.ClickPoint(entity.P37, 2, device)
	fmt.Println(utils.Now(), "点击免费按钮")
	adb.ClickPoint(entity.P36, 4, device)
	check1(id, device)
	if statuMap[device] {
		return
	}
	fmt.Println(utils.Now(), "点击跳转按钮")
	adb.ClickPoint(entity.P37, 2, device)
	fmt.Println(utils.Now(), "点击升级界面升级按钮")
	adb.ClickPoint(entity.P29, 3, device)
	fmt.Println(utils.Now(), "打开王者之路界面")
	adb.ClickPoint(entity.P30, 3, device)
	fmt.Println(utils.Now(), "点击领取奖励")
	adb.ClickPoint(entity.P31, 3, device)
	fmt.Println(utils.Now(), "点击空白区域")
	adb.ClickPoint(entity.P34, 3, device)
	fmt.Println(utils.Now(), "点击第一个任务")
	adb.ClickPoint(entity.P31, 3, device)
	res(device)
	fmt.Println(utils.Now(), "点击免费按钮")
	adb.ClickPoint(entity.P40, 1, device)
	fmt.Println(utils.Now(), "点击空白区域")
	adb.ClickPoint(entity.P34, 3, device)
	adb.ScreenTwo(device)
	fmt.Println(utils.Now(), "升级城墙")
	adb.ClickPoint(entity.CC, 2, device)
	adb.ClickPoint_(entity.CC, 3, device)
	fmt.Println(utils.Now(), "点击升级界面升级按钮")
	adb.ClickPoint(entity.P29, 3, device)
}

func createThree(id, device string) {
	fmt.Println(utils.Now(), "升级兵营")
	adb.ClickPoint(entity.Bingyin1, 3, device)
	adb.ClickPoint_(entity.Bingyin1, 3, device)
	adb.ClickMore(entity.P36, 3, 1, device)
	adb.ClickPoint(entity.Back, 3, device)
	adb.ScreenOne(device)
	adb.ClickPoint(entity.Bingyin2, 3, device)
	for i := 0; i < 6; i++ {
		if Compare1(entity.BingYing0Img(device), device) {
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
	fmt.Println(utils.Now(), "点击王者之路")
	adb.ClickPoint(entity.P30, 3, device)
	adb.ClickMore(entity.P47, 6, 1, device)
	fmt.Println(utils.Now(), "点击领取奖励")
	adb.ClickPoint(entity.P31, 2, device)
	fmt.Println(utils.Now(), "点击空白区域")
	adb.ClickPoint(entity.P34, 2, device)
	adb.ClickPoint(entity.P31, 2, device)
	adb.ClickPoint(entity.P34, 2, device)
	adb.ClickPoint(entity.Back, 3, device)
	fmt.Println(utils.Now(), "招募")
	zhaoMuProcess(entity.Bingyin1, 2, device, 0)
	zhaoMuProcess(entity.Bingyin2, 1, device, 0)
	fmt.Println(utils.Now(), "升级仓库")
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
		if Compare1(img, device) {
			break
		}
		adb.ScreenOne(device)
		adb.ClickPoint(entity.Cangku, 2, device)
		adb.ClickPoint_(entity.Cangku, 3, device)
	}
	adb.ClickMore(entity.P36, 3, 1, device)
	adb.ClickPoint(entity.Back, 3, device)
	fmt.Println(utils.Now(), "升级战争广场")
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
		if Compare1(img, device) {
			break
		}
		adb.ScreenOne(device)
		adb.ClickPoint(entity.ZhanZhen, 2, device)
		adb.ClickPoint_(entity.ZhanZhen, 3, device)
	}
	adb.ClickMore(entity.P36, 3, 2, device)
	adb.ClickPoint(entity.Back, 3, device)
	fmt.Println(utils.Now(), "领取宠物")
	adb.ClickPoint(entity.P48, 3, device)
	adb.ClickMore(entity.P49, 2, 2, device)
	adb.ClickPoint(entity.P50, 10, device)
	adb.ClickPoint(entity.Back, 3, device)
}

func createFour(id, device string) {
	adb.ScreenTwo(device)
	fmt.Println(utils.Now(), "升级医院")
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
		if Compare1(img, device) {
			break
		}
		adb.ScreenTwo(device)
		adb.ClickPoint(entity.YiYuan, 2, device)
		adb.ClickPoint_(entity.YiYuan, 2, device)
	}
	adb.ClickMore(entity.P36, 3, 1, device)
	adb.ClickPoint(entity.Back, 3, device)
	fmt.Println(utils.Now(), "创建学院")
	adb.ClickPoint(entity.XueYuan, 3, device)
	adb.ClickPoint(entity.P36, 3, device)
	adb.ClickPoint(entity.XueYuan, 2, device)
	adb.ClickPoint_(entity.XueYuan, 3, device)
	adb.ClickMore(entity.P36, 3, 1, device)
	adb.ClickPoint(entity.Back, 3, device)
	fmt.Println(utils.Now(), "创建使馆")
	adb.ClickPoint(entity.ShiGuan, 3, device)
	adb.ClickPoint(entity.P36, 3, device)
	adb.ClickPoint(entity.ShiGuan, 2, device)
	adb.ClickPoint_(entity.ShiGuan, 3, device)
	adb.ClickMore(entity.P36, 3, 1, device)
	adb.ClickPoint(entity.Back, 3, device)
	fmt.Println(utils.Now(), "创建熔炼工坊")
	adb.ClickPoint(entity.RongLian, 3, device)
	adb.ClickPoint(entity.P36, 3, device)
	adb.ClickPoint(entity.RongLian, 2, device)
	adb.ClickPoint_(entity.RongLian, 3, device)
	adb.ClickMore(entity.P36, 3, 1, device)
	adb.ClickPoint(entity.Back, 3, device)
	fmt.Println(utils.Now(), "点击免费按钮")
	adb.ClickPoint(entity.P40, 1, device)
	fmt.Println(utils.Now(), "点击空白区域")
	adb.ClickPoint(entity.P34, 3, device)
	updateChengbao(id, device)
	taskJiangLi(device)
	fmt.Println(utils.Now(), device, "创建18号资源")
	adb.ScreenSeven(device)
	adb.ScreenSeven(device)
	adb.ClickPoint(entity.ZY18, 3, device)
	for i := 0; i < 5; i++ {
		img := entity.ShiBingImg(device)
		if Compare1(img, device) {
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
	fmt.Println(utils.Now(), device, "升级侦察塔")
	adb.ClickPoint(entity.ZhenCha, 2, device)
	adb.ClickPoint_(entity.ZhenCha, 3, device)
	adb.ClickMore(entity.P36, 3, 1, device)
	adb.ClickPoint(entity.Back, 3, device)
	adb.ScreenThree(device)
	fmt.Println(utils.Now(), device, "加工材料")
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
	closePop(device)

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
		if Compare1(img, device) {
			fmt.Println(utils.Now(), "点击用户协议确认按钮")
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
	fmt.Println(utils.Now(), "创建账号")
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
	acc.GuaiWu = 3
	acc.BingYing1 = 4
	acc.BingYing2 = 4
	acc.NongChang = 5
	acc.MuChang = 5
	acc.ShiBing = 5
	acc.JiJiu = 5
	return acc
}
