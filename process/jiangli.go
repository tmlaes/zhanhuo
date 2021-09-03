package process

import (
	"fmt"
	"time"
	"zhanhuo/adb"
	"zhanhuo/entity"
	"zhanhuo/utils"
)

func jiangLi(account *entity.Account) {
	device := account.Name
	if account.LastUpdateTime == "" {
		account.LastUpdateTime = utils.DATE_FORMAT
	}
	t, _ := time.ParseInLocation(utils.DATE_FORMAT, account.LastUpdateTime, time.Local)
	now := time.Now()
	sub := now.Sub(t)
	checkBack(device)
	if sub.Hours() > 12 {
		zhengWuJiangli(account.Id, device)
		account.LastUpdateTime = now.Format(utils.DATE_FORMAT)
		if statuMap[device] {
			return
		}
	}
	maCheJiangLi(device)
	if sub.Hours() > 24 {
		taskJiangLi(device)
		account.LastUpdateTime = now.Format(utils.DATE_FORMAT)
	}
}

func taskJiangLi(device string) {
	fmt.Println(utils.Now(), device, "点击任务奖励")
	adb.ClickPoint(entity.P41, 3, device)
	if !Compare1(entity.JiangLiTaskImg(device), device) {
		fmt.Println(utils.Now(), device, "不是奖励页面")
		if Compare1(entity.BackImg(device), device) {
			adb.ClickPoint(entity.Back, 2, device)
		}
		return
	}
	fmt.Println(utils.Now(), "领取军备奖励")
	text := GetText(device, entity.Img2(device))
	if text == "建设任务" {
		adb.ClickPoint(entity.P51, 2, device)
	}
	for {
		adb.ClickPoint(entity.P42, 1, device)
		if Compare1(entity.TaskGoImg(device), device) {
			adb.ClickPoint(entity.Back, 2, device)
			break
		}
	}
	adb.ClickPoint(entity.P51, 2, device)
	fmt.Println(utils.Now(), "领取建设奖励")
	adb.ClickPoint(entity.P52, 2, device)
	for {
		adb.ClickPoint(entity.P44, 1, device)
		if Compare1(entity.TaskGoImg(device), device) {
			adb.ClickPoint(entity.Back, 2, device)
			break
		}
	}
	adb.ClickPoint(entity.P52, 2, device)
	fmt.Println(utils.Now(), "领取其他奖励")
	adb.ClickPoint(entity.P53, 2, device)
	for {
		adb.ClickPoint(entity.P46, 1, device)
		if Compare1(entity.TaskGoImg(device), device) {
			adb.ClickPoint(entity.Back, 2, device)
			break
		}
	}
	adb.ClickPoint(entity.Back, 3, device)
}

func zhengWuJiangli(id, device string) {
	fmt.Println(utils.Now(), device, "领取政务日常奖励")
	adb.ScreenOne(device)
	adb.ClickPoint(entity.Zhengwu, 2, device)
	adb.ClickPoint(entity.P55, 2, device)
	adb.ClickPoint(entity.Zhengwu, 2, device)
	adb.ClickPoint_(entity.Zhengwu, 3, device)

	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到政务日常，退出")
			return
		}
		if Compare1(entity.JiangLiImg(device), device) {
			break
		}
		adb.ScreenOne(device)
		adb.ClickPoint(entity.Zhengwu, 2, device)
		adb.ClickPoint_(entity.Zhengwu, 3, device)
	}
	defer adb.ClickPoint(entity.Back, 5, device)
	var arr [6]entity.Point
	arr[0] = entity.P66
	arr[1] = entity.P67
	arr[2] = entity.P68
	arr[3] = entity.P69
	arr[4] = entity.P70
	arr[5] = entity.P71
	for i := 0; i < len(arr); i++ {
		adb.ClickPoint(arr[i], 2, device)
		if i == 2 || i == 5 {
			adb.ClickPoint(entity.P73, 1, device)
		} else {
			adb.ClickPoint(entity.P72, 1, device)
		}
		adb.ClickPoint(entity.P55, 2, device)
	}
}

func maCheJiangLi(device string) {
	fmt.Println(utils.Now(), device, "领取马车奖励")
	checkLianMengLiWu(device)
	adb.ScreenTwo(device)
	adb.ClickPoint(entity.MaChe, 1, device)
	adb.ClickPoint(entity.P55, 2, device)
}
