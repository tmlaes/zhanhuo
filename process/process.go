package process

import (
	"fmt"
	"time"
	"zhanhuo/adb"
	P "zhanhuo/entity"
	"zhanhuo/utils"
)

const (
	GameX = 640
	GameY = 120
)

func Process() {
	//获取在线设备
	devices := adb.Devices()
	//读取配置
	for _, device := range devices {
		account := utils.ReadJons(device)
		start(device)
		ChengBao(account)
		ZhaoMu(account)
		YeGuai(account)
		CaiGi(account)
		utils.WriteJons(account)
	}
}

func start(device string) {
	adb.Click(GameX, GameY, 5, device)
	ticker := time.NewTicker(3 * time.Second)
	for range ticker.C {
		if CheckStart(device) {
			fmt.Println(time.Now(), device, "已进入游戏界面")
			break
		}
		fmt.Println(time.Now(), device, "正在加载游戏中.....")
	}
}

func task(device string) {
	//第一步打开任务列表
	adb.ScreenTwo(device)
	fmt.Println(time.Now(), device, "打开任务列表")
	adb.Click(P.Task.X, P.Task.Y, 2, device)
}

func sleep(times int) {
	time.Sleep(time.Duration(times) * time.Second)
}
