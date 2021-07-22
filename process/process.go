package process

import (
	"fmt"
	"time"
	"zhanhuo/adb"
	P "zhanhuo/entity"
	"zhanhuo/utils"
)

func Process() {
	//获取在线设备
	devices := adb.Devices()
	//读取配置
	for _, device := range devices {
		account := utils.ReadJons(device)
		ChengBao(account)
		ZhaoMu(account)
		YeGuai(account)
		CaiGi(account)
		utils.WriteJons(account)
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
