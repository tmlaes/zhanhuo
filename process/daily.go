package process

import (
	"fmt"
	"time"
	"zhanhuo/adb"
	"zhanhuo/entity"
	"zhanhuo/utils"
)

func daily(account *entity.Account) {
	device := account.Name
	now := time.Now()
	//摇一摇
	adb.Shake(account.Id)
	delay(2)
	//丰收技能
	if account.FengShouTime == "" {
		account.FengShouTime = now.Format(utils.DATE_FORMAT)
	} else {
		t, _ := time.ParseInLocation(utils.DATE_FORMAT, account.FengShouTime, time.Local)
		sub := now.Sub(t)
		if sub.Hours() >= 12 {
			adb.ClickPoint(entity.P74, 3, device)
			adb.ClickPoint(entity.P75, 2, device)
			adb.ClickPoint(entity.P76, 2, device)
			adb.ClickMore(entity.P55, 3, 1, device)
			add := now.Add(time.Duration(15) * time.Second)
			account.FengShouTime = add.Format(utils.DATE_FORMAT)
			utils.WriteJons(account)
		}
	}
}

func juanxian(device string) {
	fmt.Println(utils.Now(), device, "联盟捐献")
	adb.ClickPoint(entity.P79, 2, device)
	adb.ClickPoint(entity.P80, 2, device)
	adb.ClickPoint(entity.P81, 2, device)
	adb.ClickMore(entity.P82, 5, 1, device)
	if utils.Compare1(entity.Img10(device), device) {
		adb.ClickPoint(entity.P83, 2, device)
	}
	adb.ClickPoint(entity.Close, 2, device)
	adb.ClickMore(entity.Back, 2, 2, device)
}

func updateLastTime(account *entity.Account) {
	now := time.Now()
	if account.LastUpdateTime == "" {
		account.LastUpdateTime = now.Format(utils.DATE_FORMAT)
		return
	}
	t, _ := time.ParseInLocation(utils.DATE_FORMAT, account.LastUpdateTime, time.Local)
	sub := now.Sub(t)
	if sub.Hours() >= 24 {
		account.LastUpdateTime = now.Format(utils.DATE_FORMAT)
	}
}
