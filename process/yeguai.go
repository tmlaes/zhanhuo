package process

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"zhanhuo/adb"
	P "zhanhuo/entity"
	"zhanhuo/utils"
)

func YeGuai(account *P.Account) {
	device := account.Name
	fmt.Println(time.Now(), device, "读取体力值")
	text := GetText1(device, P.TiLiImg(device))
	v, _ := strconv.Atoi(text)
	if v < 100 {
		return
	}
	fmt.Println(time.Now(), device, "体力值：", v)
	zv := zhanli(device)
	fmt.Println(time.Now(), device, "战力值：", zv)
	adb.Click(P.Reset.X, P.Reset.Y, 5, device)
	for i := 0; i < 10; i++ {
		fmt.Println(time.Now(), device, "点击搜索按钮")
		adb.Click(P.Search.X, P.Search.Y, 2, device)
		fmt.Println(time.Now(), device, "点击怪物图标")
		adb.Click(P.GW.X, P.GW.Y, 1, device)
		if i == 0 {
			//怪物最高等级
			text1 := GetText1(device, P.GuaiWuImg(device))
			cb := utils.GetBetweenStr(text1, "战", "级")
			level, _ := strconv.Atoi(cb)
			if P.ZhanLiMap[cb] > zv {
				level = level - 1
			}
			fmt.Println(time.Now(), device, "点击等级输入框")
			adb.Click(P.GWL.X, P.GWL.Y, 1, device)
			fmt.Println(time.Now(), device, "输入等级")
			adb.Text(device, strconv.Itoa(level))
			fmt.Println(time.Now(), device, "回车")
			adb.Key(device, "4")
			time.Sleep(time.Duration(1) * time.Second)
		}
		fmt.Println(time.Now(), device, "点击前往")
		adb.Click(P.GoTO.X, P.GoTO.Y, 1, device)
		fmt.Println(time.Now(), device, "关闭前往对话框")
		adb.Click(P.CloseW.X, P.CloseW.Y, 1, device)
		for {
			fmt.Println(time.Now(), device, "点击怪物")
			adb.Click(P.GgW.X, P.GgW.Y, 2, device)
			shotImg := adb.Screenshot(device)
			img := P.GjImg(device)
			CutImage(shotImg, img)
			flag := Compare(img.Name, GojImg)
			if flag {
				break
			}
		}
		adb.Click(P.GoJi.X, P.GoJi.Y, 3, device)
		text1 := GetText1(device, P.TimeImg(device))
		fmt.Println(time.Now(), device, "获取行军时间", text1)
		split := strings.Split(text1, ":")
		min, _ := strconv.Atoi(split[0])
		sec, _ := strconv.Atoi(split[1])
		t := 2*(min*60+sec) + 3
		adb.Click(P.CF.X, P.CF.Y, t, device)
	}
	fmt.Println(time.Now(), device, "怪物攻击完成")
	adb.Click(P.Reset.X, P.Reset.Y, 5, device)
}

func zhanli(device string) int {
	fmt.Println(time.Now(), device, "读取战力值")
	for {
		adb.ScreenOne(device)
		fmt.Println(time.Now(), device, "点击战争广场")
		adb.Click(P.ZhanZhen.X, P.ZhanZhen.Y, 2, device)
		fmt.Println(time.Now(), device, "点击战争广场招募")
		adb.Click(P.ZhanZhen.X_+110, P.ZhanZhen.Y_-20, 3, device)
		shotImg := adb.Screenshot(device)
		img := P.BuDuiImg(device)
		CutImage(shotImg, img)
		flag := Compare(img.Name, BuDuiImg)
		if flag {
			break
		}
	}
	fmt.Println(time.Now(), device, "点击编队一")
	adb.Click(P.BD1.X, P.BD1.Y, 2, device)
	fmt.Println(time.Now(), device, "获取战力")
	text := GetText1(device, P.ZhanliImg(device))
	v, _ := strconv.Atoi(text)
	fmt.Println(time.Now(), device, "点击返回")
	adb.Click(P.Back.X, P.Back.Y, 2, device)
	fmt.Println(time.Now(), device, "点击返回")
	adb.Click(P.Back.X, P.Back.Y, 2, device)
	return v
}
