package process

import (
	"fmt"
	"strconv"
	"strings"
	"zhanhuo/adb"
	P "zhanhuo/entity"
	"zhanhuo/utils"
)

func YeGuai(id, device string, level int) int {
	closePop(device)
	fmt.Println(utils.Now(), "打野")
	text := GetText(device, P.TiLiImg(device))
	v, _ := strconv.Atoi(text)
	fmt.Println(utils.Now(), device, "体力值：", v)
	if v < 10 {
		return level
	}
	adb.ClickPoint(P.Reset, 5, device)
	maxZhanli := getZhanLi(id, device)
	if statuMap[device] || maxZhanli == 0 {
		adb.ClickPoint(P.Reset, 5, device)
		return level
	}
	maxLevel := getMaxLevel(maxZhanli)
	fmt.Println(utils.Now(), device, "最高怪物等级：", maxLevel)
	if maxLevel == 0 || maxLevel < level {
		adb.ClickPoint(P.Reset, 5, device)
		return level
	}
	i := 0
	for ; i < 15; i++ {
		if v < 10 {
			break
		}
		closePop(device)
		adb.ClickPoint(P.Search, 2, device)
		adb.ClickPoint(P.GW, 1, device)
		if maxLevel > level {
			if level > 0 {
				adb.ClickPoint(P.P56, 1, device)
			}
			level = level + 1
		}
		adb.ClickPoint(P.GoTO, 2, device)
		adb.ClickPoint(P.GgW, 1, device)
		j := 0
		for ; j < 5; j++ {
			fmt.Println(utils.Now(), device, "点击怪物")
			adb.ClickPoint(P.GgW, 1, device)
			if Compare1(P.Img3(device), device) {
				break
			}
			if Compare1(P.Img9(device), device) {
				adb.ClickPoint(P.GoJi, 2, device)
				adb.ClickPoint(P.P77, 1, device)
				adb.ClickPoint(P.GoTO, 2, device)
			}

		}
		if j >= 4 {
			adb.ClickPoint(P.P55, 1, device)
			continue
		}
		adb.ClickPoint(P.GoJi, 2, device)
		text1 := GetText(device, P.TimeImg(device))
		fmt.Println(utils.Now(), device, "获取行军时间", text1)
		split := strings.Split(text1, ":")
		min, _ := strconv.Atoi(split[0])
		sec, _ := strconv.Atoi(split[1])
		t := 2*(min*60+sec) + 3
		adb.ClickPoint(P.CF, t, device)
		v = v - 10
	}
	if i > 13 {
		adb.Quit(id)
		statuMap[device] = true
		return level
	}
	fmt.Println(utils.Now(), device, "打野完成")
	zhiliao(device)
	return level
}

func getZhanLi(id, device string) int64 {
	adb.ClickPoint(P.Search, 2, device)
	adb.ClickPoint(P.GW, 1, device)
	adb.ClickPoint(P.GoTO, 2, device)
	adb.ClickPoint(P.P55, 2, device)
	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到怪物，退出")
			return 0
		}
		fmt.Println(utils.Now(), device, "点击怪物")
		adb.ClickPoint(P.GgW, 1, device)
		if Compare1(P.Img3(device), device) {
			break
		}
	}
	adb.ClickPoint(P.GoJi, 2, device)
	if Compare1(P.Img6(device), device) {
		fmt.Println(utils.Now(), device, "队伍不足")
		adb.ClickPoint(P.Back, 2, device)
		return 0
	}
	adb.ClickPoint(P.P57, 2, device)
	adb.ClickPoint(P.P58, 2, device)
	text := GetText(device, P.Img4(device))
	zl, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(utils.Now(), device, "攻击力：", zl)
	adb.ClickPoint(P.Back, 2, device)
	return zl
}

func getMaxLevel(zl int64) int {
	for index, value := range P.ZhanLiArr {
		if value > zl {
			return index
		}
	}
	return 1
}

func zhiliao(device string) {
	closePop(device)
	adb.ClickPoint(P.Reset, 5, device)
	closePop(device)
	fmt.Println(utils.Now(), device, "治疗伤兵")
	adb.ClickPointOffset(P.YiYuan, -10, -80, 3, device)
	if Compare1(P.BackImg(device), device) {
		adb.ClickPoint(P.CF, 1, device)
		checkZiYuanBao(device)
	}
}

func Zhiliao(device string) {
	fmt.Println(utils.Now(), device, "点击治疗伤完成")
	adb.ClickPointOffset(P.YiYuan, -10, -80, 3, device)
}
