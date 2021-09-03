package process

import (
	"fmt"
	"zhanhuo/adb"
	P "zhanhuo/entity"
	"zhanhuo/utils"
)

func ZhaoMu(account *P.Account) {
	device := account.Name
	status := getBingYingStatus(device, account.BingYing3)
	for i := 0; i < len(status); i++ {
		if status[i] == 0 {
			fmt.Println(utils.Now(), device, "兵营", (i + 1), "不是空闲和收取状态")
			continue
		}
		adb.ClickPoint(P.Task, 2, device)
		if i == 0 {
			adb.ClickPoint(P.Task3, 2, device)
		}
		if i == 1 {
			adb.ClickPoint(P.Task4, 2, device)
		}
		if i == 2 {
			adb.ClickPoint(P.Task5, 2, device)
		}
		zhaoMuProcess1(P.P78, i, device, status[i])
	}
}

func zhaoMuProcess1(p P.Point, index int, device string, status int) {
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

		fmt.Println(utils.Now(), device, "点击兵营")
		adb.ClickPoint(p, 2, device)
		fmt.Println(utils.Now(), device, "点击招募按钮")
		adb.ClickPointOffset(p, 110, 100, 3, device)
		if Compare1(P.ZhaoMuImg(device), device) {
			break
		}
		adb.ClickPoint(P.Task, 2, device)
		if index == 0 {
			adb.ClickPoint(P.Task3, 2, device)
		}
		if index == 1 {
			adb.ClickPoint(P.Task4, 2, device)
		}
		if index == 2 {
			adb.ClickPoint(P.Task5, 2, device)
		}

	}
	if index == 0 {
		fmt.Println(utils.Now(), device, "点击法师")
		adb.ClickPoint(P.FaShi, 2, device)
	} else {
		fmt.Println(utils.Now(), device, "点击骑士")
		adb.ClickPoint(P.QiBin, 2, device)
	}
	adb.ClickPoint(P.Shengji, 2, device)
	checkZiYuanBao(device)
}

func zhaoMuProcess(p P.Point, index int, device string, status int) {
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

		fmt.Println(utils.Now(), device, "点击兵营")
		adb.ClickPoint(p, 2, device)
		fmt.Println(utils.Now(), device, "点击招募按钮")
		adb.ClickPointOffset(p, 110, 100, 3, device)
		if Compare1(P.ZhaoMuImg(device), device) {
			break
		}
	}
	if index == 2 {
		fmt.Println(utils.Now(), device, "点击法师")
		adb.ClickPoint(P.FaShi, 2, device)
	} else {
		fmt.Println(utils.Now(), device, "点击骑士")
		adb.ClickPoint(P.QiBin, 2, device)
	}
	adb.ClickPoint(P.Shengji, 2, device)
	checkZiYuanBao(device)
}

func getBingYingStatus(device string, level3 int) [4]int {
	task(device)
	var status [4]int
	if level3 > 0 {
		if Compare1(P.ZhaoMuBY1KongImg(device), device) {
			fmt.Println(utils.Now(), device, "第一个兵营状态为空闲中")
			status[0] = 1
		}
		if Compare1(P.ZhaoMuBY1ShouImg(device), device) {
			fmt.Println(utils.Now(), device, "第一个兵营状态为可收取")
			status[0] = 2
		}
		if Compare1(P.ZhaoMuBY2KongImg(device), device) {
			fmt.Println(utils.Now(), device, "第二个兵营状态为空闲中")
			status[1] = 1
		}
		if Compare1(P.ZhaoMuBY2ShouImg(device), device) {
			fmt.Println(utils.Now(), device, "第二个兵营状态为可收取")
			status[1] = 2
		}

		if Compare1(P.ZhaoMuBY3KongImg(device), device) {
			fmt.Println(utils.Now(), device, "第三个兵营状态为空闲中")
			status[2] = 1
		}
		if Compare1(P.ZhaoMuBY3ShouImg(device), device) {
			fmt.Println(utils.Now(), device, "第三个兵营状态为可收取")
			status[2] = 2
		}
	} else {
		if Compare1(P.ZhaoMuBY1KongImg(device), device) {
			fmt.Println(utils.Now(), device, "第一个兵营状态为空闲中")
			status[0] = 1
		}
		if Compare1(P.ZhaoMuBY1ShouImg(device), device) {
			fmt.Println(utils.Now(), device, "第二个兵营状态为可收取")
			status[0] = 2
		}
		if Compare1(P.ZhaoMuBY2KongImg(device), device) {
			fmt.Println(utils.Now(), device, "第二个兵营状态为空闲中")
			status[1] = 1
		}
		if Compare1(P.ZhaoMuBY2ShouImg(device), device) {
			fmt.Println(utils.Now(), device, "第二个兵营状态为可收取")
			status[1] = 2
		}
	}
	adb.ClickPoint(P.P55, 2, device)
	fmt.Println(utils.Now(), device, status)
	return status
}
