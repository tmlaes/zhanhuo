package process

import (
	"fmt"
	"zhanhuo/adb"
	"zhanhuo/entity"
	"zhanhuo/utils"
)

func ZhaoMu(account *entity.Account) {
	device := account.Name
	status := getBingYingStatus(device, account.BingYing3)
	for i := 0; i < len(status); i++ {
		if status[i] == 0 {
			fmt.Println(utils.Now(), account.Id, device, "兵营", (i + 1), "不是空闲和收取状态")
			continue
		}
		adb.ClickPoint(entity.Task, 2, device)
		if i == 0 {
			adb.ClickPoint(entity.Task3, 2, device)
		}
		if i == 1 {
			adb.ClickPoint(entity.Task4, 2, device)
		}
		if i == 2 {
			adb.ClickPoint(entity.Task5, 2, device)
		}
		zhaoMuProcess1(entity.P78, i, device, status[i])
	}
}

func zhaoMuProcess1(p entity.Point, index int, device string, status int) {
	for i := 0; i < 6; i++ {
		if i == 5 {
			return
		}
		if status == 2 {
			adb.ClickPoint(p, 1, device)
			if i > 1 {
				adb.ClickPoint(p, 1, device)
			}
		}

		adb.ClickPoint(p, 2, device)
		fmt.Println(utils.Now(), device, "点击招募按钮")
		adb.ClickPointOffset(p, 110, 100, 3, device)
		if adb.Compare1(entity.ZhaoMuImg(device), device) {
			break
		}
		adb.ClickPoint(entity.Task, 2, device)
		if index == 0 {
			adb.ClickPoint(entity.Task3, 2, device)
		}
		if index == 1 {
			adb.ClickPoint(entity.Task4, 2, device)
		}
		if index == 2 {
			adb.ClickPoint(entity.Task5, 2, device)
		}

	}
	if index == 0 {
		adb.ClickPoint(entity.FaShi, 2, device)
	} else {
		adb.ClickPoint(entity.QiBin, 2, device)
	}
	adb.ClickPoint(entity.Shengji, 2, device)
	checkZiYuanBao(device)
}

func zhaoMuProcess(p entity.Point, index int, device string, status int) {
	for i := 0; i < 6; i++ {
		if i == 5 {
			return
		}
		if index == 1 {
			adb.ScreenOne(device)
		} else {
			checkLianMengLiWu(device)
			adb.ScreenTwo(device)
		}
		if status == 2 {
			adb.ClickPoint(p, 1, device)
			if i > 1 {
				adb.ClickPoint(p, 1, device)
			}
		}

		adb.ClickPoint(p, 2, device)
		adb.ClickPointOffset(p, 110, 100, 3, device)
		if adb.Compare1(entity.ZhaoMuImg(device), device) {
			break
		}
	}
	if index == 2 {
		adb.ClickPoint(entity.FaShi, 2, device)
	} else {
		adb.ClickPoint(entity.QiBin, 2, device)
	}
	adb.ClickPoint(entity.Shengji, 2, device)
	checkZiYuanBao(device)
}

func getBingYingStatus(device string, level3 int) [4]int {
	task(device)
	var status [4]int
	if level3 > 0 {
		if adb.Compare1(entity.ZhaoMuBY1KongImg(device), device) {
			fmt.Println(utils.Now(), device, "第一个兵营状态为空闲中")
			status[0] = 1
		}
		if adb.Compare1(entity.ZhaoMuBY1ShouImg(device), device) {
			fmt.Println(utils.Now(), device, "第一个兵营状态为可收取")
			status[0] = 2
		}
		if adb.Compare1(entity.ZhaoMuBY2KongImg(device), device) {
			fmt.Println(utils.Now(), device, "第二个兵营状态为空闲中")
			status[1] = 1
		}
		if adb.Compare1(entity.ZhaoMuBY2ShouImg(device), device) {
			fmt.Println(utils.Now(), device, "第二个兵营状态为可收取")
			status[1] = 2
		}

		if adb.Compare1(entity.ZhaoMuBY3KongImg(device), device) {
			fmt.Println(utils.Now(), device, "第三个兵营状态为空闲中")
			status[2] = 1
		}
		if adb.Compare1(entity.ZhaoMuBY3ShouImg(device), device) {
			fmt.Println(utils.Now(), device, "第三个兵营状态为可收取")
			status[2] = 2
		}
	} else {
		if adb.Compare1(entity.ZhaoMuBY1KongImg(device), device) {
			fmt.Println(utils.Now(), device, "第一个兵营状态为空闲中")
			status[0] = 1
		}
		if adb.Compare1(entity.ZhaoMuBY1ShouImg(device), device) {
			fmt.Println(utils.Now(), device, "第二个兵营状态为可收取")
			status[0] = 2
		}
		if adb.Compare1(entity.ZhaoMuBY2KongImg(device), device) {
			fmt.Println(utils.Now(), device, "第二个兵营状态为空闲中")
			status[1] = 1
		}
		if adb.Compare1(entity.ZhaoMuBY2ShouImg(device), device) {
			fmt.Println(utils.Now(), device, "第二个兵营状态为可收取")
			status[1] = 2
		}
	}
	adb.ClickPoint(entity.P55, 2, device)
	fmt.Println(utils.Now(), device, status)
	return status
}
