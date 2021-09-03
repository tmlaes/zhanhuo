package process

import (
	"fmt"
	"strings"
	"sync"
	"time"
	"zhanhuo/adb"
	"zhanhuo/entity"
	"zhanhuo/utils"
)

var wg1 sync.WaitGroup
var statuMap = make(map[string]bool)

func Process() {
	deviceList := utils.ReadDevices()
	for _, list := range deviceList {
		//启动设备
		devices := adb.Devices_(list)
		//读取配置
		for device, _ := range devices {
			wg1.Add(1)
			go startDevice(device)
		}
		wg1.Wait()
	}
	/*var t = time.Tick(30 * time.Minute)
	for {
		<-t
		delayTask()
	}*/
}

func startDevice(device string) {
	account := utils.ReadJons(device)
	doProcess(account)
	wg1.Done()
}

func doProcess(account *entity.Account) {
	defer utils.WriteJons(account)
	device := account.Name
	start(device)
	jiangLi(account)
	if statuMap[device] {
		return
	}
	daily(account)
	build(account)
	if statuMap[device] {
		return
	}
	account.BuildTime = getBuildTime(device)
	utils.WriteJons(account)
	ZhaoMu(account)
	ygLevel := YeGuai(account.Id, device, account.GuaiWu)
	account.SetField("GuaiWu", ygLevel)
	utils.WriteJons(account)
	if statuMap[device] {
		return
	}
	Caiji(account)
	Zhiliao(device)
	//updateLastTime(account)
	adb.Quit(account.Id)
}

func start(device string) {
	adb.ClickPoint(entity.Game, 30, device)
	ticker := time.NewTicker(3 * time.Second)
	for range ticker.C {
		if CheckStart(device) {
			fmt.Println(utils.Now(), device, "已进入游戏界面")
			return
		}
		fmt.Println(utils.Now(), device, "正在加载游戏中.....")
	}
}

func task(device string) {
	//第一步打开任务列表
	adb.ScreenZero(device)
	adb.ClickPoint(entity.Task, 2, device)
}

func addTime(text string) string {
	split := strings.Split(text, ":")
	h := "0"
	m := "0"
	s := "0"
	if len(split) == 3 {
		h = split[0]
		m = split[1]
		s = split[2]
	}
	if len(split) == 2 {
		m = split[0]
		s = split[1]
	}
	now := time.Now()
	t := h + "h" + m + "m" + s + "s"
	mm, err := time.ParseDuration(t)
	if err != nil {
		return ""
	}
	mm1 := now.Add(mm)
	st := mm1.Format(utils.DATE_FORMAT)
	return st
}
func delay(times int) {
	<-time.After(time.Duration(times) * time.Second)
}

func delayTask() {
	deviceList := utils.ReadDevices()
	for _, list := range deviceList {
		//启动设备
		devices := adb.Devices_(list)
		//读取配置
		for device, _ := range devices {
			wg1.Add(1)
			go startDevice(device)
		}
		wg1.Wait()
	}
}
