package process

import (
	"fmt"
	"strconv"
	"time"
	"zhanhuo/adb"
	"zhanhuo/entity"
	"zhanhuo/utils"
)

func build(account *entity.Account) {
	//校验队列1是否处于空闲状态
	device := account.Name
	if account.BuildTime != "" {
		t, _ := time.ParseInLocation(utils.DATE_FORMAT, account.BuildTime, time.Local)
		if time.Now().Before(t) {
			fmt.Println(utils.Now(), device, "建筑时间还未到")
			return
		}
	} else {
		task(device)
		if !Compare1(entity.TaskBuild1Img(device), device) {
			fmt.Println(utils.Now(), device, "建筑队列不是空闲中")
			adb.ClickPoint(entity.P55, 2, device)
			return
		}
		adb.ClickPoint(entity.P55, 2, device)
	}
	level := account.ChengBao
	key := "level" + strconv.Itoa(level)
	for _, item := range entity.TaskMap[key] {
		oldLevel := account.GetFieldInt(item.Name)
		if oldLevel >= item.Level {
			continue
		}
		b := update(account.Id, item.Name, device)
		if statuMap[device] {
			return
		}
		if b {
			account.SetField(item.Name, oldLevel+1)
			account.BuildTime = getBuildTime(device)
			utils.WriteJons(account)
			return
		}
	}
	update(account.Id, "ChengBao", device)
	if statuMap[device] {
		return
	}
	account.SetField("ChengBao", level+1)
	account.BuildTime = getBuildTime(device)
	utils.WriteJons(account)
}

func getBuildTime(device string) string {
	fmt.Println(utils.Now(), device, "获取建筑时间")
	defer adb.ClickPoint(entity.P55, 2, device)
	task(device)
	for i := 0; i < 10; i++ {
		text := GetText1(device, entity.Task1Img(device))
		fmt.Println(utils.Now(), device, "建筑结束时间", text)
		if text != "" {
			st := addTime(text)
			return st
		}
	}
	fmt.Println(utils.Now(), device, "未获取到建筑时间")
	return ""
}
