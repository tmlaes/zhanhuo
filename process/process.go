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
		chengbao(account)
		utils.WriteJons(account)
	}

}

func chengbao(account *P.Account) {
	//校验队列1是否处于空闲状态
	/*	device :=account.Name
		task(device)
		for  {
			shotImg := adb.Screenshot(device)
			//队列1
			text := GetText(shotImg, P.Task1Img(device))
			fmt.Println(time.Now(), device, "建筑队列1状态", text)
			if text=="" {
				continue
			}
			if text == "空闲中" {
				break
			}
			return
		}

		//关闭任务列表
		sleep(1)
		adb.Click(P.Reset.X, P.Reset.Y,1, device)*/

	cb := ChengBao{
		account,
	}
	cb.ShengJi()
}

func task(device string) {
	//第一步打开任务列表
	fmt.Println(time.Now(), device, "打开任务列表")
	adb.Click(P.Task.X, P.Task.Y, 2, device)
}

func sleep(times int) {
	time.Sleep(time.Duration(times) * time.Second)
}
