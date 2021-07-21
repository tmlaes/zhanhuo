package process

import (
	"fmt"
	"time"
	"zhanhuo/adb"
	"zhanhuo/entity"
)

type ChengBao struct {
	account *entity.Account
}

func (c *ChengBao) ShengJi() {
	//定位
	/*	device := c.account.Name
		locate(device)
		level := c.account.ChengBaoLevel
		key := "level" + strconv.Itoa(level)
		for _, item := range entity.TaskMap[key] {
			build := buildMap[item.Name]
			c.account.SetLevel(item.Name)
			if c.account.GetLevel(item.Name) == 0 {
				build.Create(device)
				return
			}
			if c.account.GetLevel(item.Name) < item.Level {
				build.Update(device)
				return
			}
		}
		for  {
			adb.ScreenOne(device)
			fmt.Println(time.Now(), device, "点击城堡")
			adb.Click(entity.CB.X, entity.CB.Y,2,device)
			fmt.Println(time.Now(), device, "点击升级按钮")
			adb.Click(entity.CB.X_, entity.CB.Y_,3,device)
			shotImg := adb.Screenshot(device)
			img := entity.ChengBaoImg(device)
			CutImage(shotImg, img)
			flag := Compare(img.Name, ChengBaoImg)
			if flag {
				break
			}
		}
		fmt.Println(time.Now(), device, "点击升级界面升级")
		adb.Click(entity.Shengji.X, entity.Shengji.Y,3,device)*/
	c.account.SetLevel("bingYing4")
	c.account.ChengBaoLevel = c.account.ChengBaoLevel + 1
}

func locate(device string) {
	for {
		adb.ScreenOne(device)
		//读取当前城堡等级
		fmt.Println(time.Now(), device, "点击城堡")
		adb.Click(entity.CB.X, entity.CB.Y, 2, device)
		fmt.Println(time.Now(), device, "点击升级按钮")
		adb.Click(entity.CB.X_, entity.CB.Y_, 3, device)

		shotImg := adb.Screenshot(device)
		img := entity.ChengBaoImg(device)
		CutImage(shotImg, img)
		flag := Compare(img.Name, ChengBaoImg)
		if flag {
			return
		}
	}
}

var buildMap = make(map[string]Build)

func init() {
	buildMap["bingYing2"] = Build{"", entity.Bingyin2}
	buildMap["chengQiang"] = Build{"", entity.CC}
	buildMap["xueYuan"] = Build{"", entity.XueYuan}
	buildMap["shiGuan"] = Build{"", entity.ShiGuan}
	buildMap["rongLian"] = Build{"", entity.RongLian}
	buildMap["zhanZhen"] = Build{"", entity.ZhanZhen}
	buildMap["kaiHuang1"] = Build{"", entity.ZYZ1}
	buildMap["kaiHuang2"] = Build{"", entity.ZYZ2}
	buildMap["cangKu"] = Build{"", entity.Cangku}
	buildMap["bingYing3"] = Build{"", entity.Bingyin3}
	buildMap["shiBing"] = Build{"", entity.ZY19}
	buildMap["jiJiu"] = Build{"", entity.ZY20}
	buildMap["yiYuan"] = Build{"", entity.YiYuan}
	buildMap["kaiHuang3"] = Build{"", entity.ZYZ1}
	buildMap["kaiHuang4"] = Build{"", entity.ZYZ1}
	buildMap["shiTou"] = Build{"", entity.ZY20}
	buildMap["bingYing4"] = Build{"", entity.Bingyin4}
	buildMap["zhenCha"] = Build{"", entity.ZhenCha}
	buildMap["faMu"] = Build{"", entity.ZY}
	buildMap["tie"] = Build{"", entity.ZY20}
}
