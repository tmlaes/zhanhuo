package process

import (
	"fmt"
	"strconv"
	"time"
	"zhanhuo/adb"
	"zhanhuo/entity"
)

type ChengBao struct {
	account *entity.Account
}

func (c *ChengBao) ShengJi() {
	//定位
	device := c.account.Name
	level := c.account.ChengBaoLevel
	key := "level" + strconv.Itoa(level)
	for _, item := range entity.TaskMap[key] {
		build := buildMap[item.Name]
		if c.account.GetLevel(item.Name) == 0 {
			fmt.Println(time.Now(), device, "建造", item.Name, build)
			build.Create(device)
			c.account.SetLevel(item.Name)
			c.account.SetLevel(item.Name)
			c.account.SetLevel(item.Name)
			c.account.SetLevel(item.Name)
			c.account.SetLevel(item.Name)
			return
		}
		if c.account.GetLevel(item.Name) < item.Level {
			fmt.Println(time.Now(), device, "升级", item.Name, build)
			build.Update(device)
			c.account.SetLevel(item.Name)
			return
		}
	}
	for {
		adb.ScreenOne(device)
		fmt.Println(time.Now(), device, "点击城堡")
		adb.Click(entity.CB.X, entity.CB.Y, 2, device)
		fmt.Println(time.Now(), device, "点击升级按钮")
		adb.Click(entity.CB.X_, entity.CB.Y_, 3, device)
		shotImg := adb.Screenshot(device)
		img := entity.ChengBaoImg(device)
		CutImage(shotImg, img)
		flag := Compare(img.Name, ChengBaoImg)
		if flag {
			break
		}
	}
	fmt.Println(time.Now(), device, "点击升级界面升级")
	adb.Click(entity.Shengji.X, entity.Shengji.Y, 3, device)
	CheckZiYuanBao(device)
	c.account.SetLevel("bingYing4")
	c.account.ChengBaoLevel = c.account.ChengBaoLevel + 1
}

var buildMap = make(map[string]Build)

func init() {
	buildMap["bingYing1"] = Build{"bingYing", 2, entity.Bingyin1}
	buildMap["bingYing2"] = Build{"bingYing", 1, entity.Bingyin2}
	buildMap["chengQiang"] = Build{"chengQiang", 2, entity.CC}
	buildMap["xueYuan"] = Build{"xueYuan", 2, entity.XueYuan}
	buildMap["shiGuan"] = Build{"shiGuan", 2, entity.ShiGuan}
	buildMap["rongLian"] = Build{"rongLian", 2, entity.RongLian}
	buildMap["zhanZhen"] = Build{"zhanZhen", 1, entity.ZhanZhen}
	buildMap["kaiHuang1"] = Build{"kaiHuang1", 5, entity.ZYZ1}
	buildMap["kaiHuang2"] = Build{"kaiHuang2", 5, entity.ZYZ2}
	buildMap["cangKu"] = Build{"cangKu", 1, entity.Cangku}
	buildMap["bingYing3"] = Build{"bingYing3", 1, entity.Bingyin3}
	buildMap["shiBing"] = Build{"shiBing", 5, entity.ZY19}
	buildMap["jiJiu"] = Build{"jiJiu", 5, entity.ZY20}
	buildMap["yiYuan"] = Build{"yiYuan", 2, entity.YiYuan}
	buildMap["kaiHuang3"] = Build{"kaiHuang3", 6, entity.ZYZ1}
	buildMap["kaiHuang4"] = Build{"kaiHuang4", 7, entity.ZYZ1}
	buildMap["shiTou"] = Build{"shiTou", 7, entity.ZY20}
	buildMap["bingYing4"] = Build{"bingYing4", 1, entity.Bingyin4}
	buildMap["zhenCha"] = Build{"zhenCha", 4, entity.ZhenCha}
	buildMap["faMu"] = Build{"faMu", 5, entity.ZY}
	buildMap["tie"] = Build{"tie", 7, entity.ZY20}
}
