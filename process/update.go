package process

import (
	"fmt"
	"zhanhuo/adb"
	"zhanhuo/entity"
	"zhanhuo/utils"
)

func update(id, operation, device string) bool {
	switch operation {
	case "ChengBao":
		updateChengbao(id, device)
	case "ChengQiang":
		updateChengQiang(id, device)
	case "CangKu":
		updateCangku(id, device)
	case "ZhanZhen":
		updateZhanZheng(id, device)
	case "YiYuan":
		updateYiYuan(id, device)
	case "XueYuan":
		updateXueyuan(id, device)
	case "RongLian":
		updateRongLian(id, device)
	case "ShiGuan":
		updateShiGuan(id, device)
	case "ZhenCha":
		updateZhenCha(id, device)
	case "BingYing1":
		updateBingYing1(id, device)
	case "BingYing2":
		updateBingYing2(id, device)
	case "NongChang":
		updateNongChang(id, device)
	case "MuChang":
		updateMuChang(id, device)
	case "ShiBing":
		updateShiBing(id, device)
	case "JiJiu":
		updateJiJiu(id, device)
	default:
		fmt.Println(utils.Now(), "序号", id, operation, "未找到升级建筑")
		statuMap[device] = true
		return false
	}
	if statuMap[device] {
		return false
	}
	return checkZiYuanBao(device)
}

func updateChengbao(id, device string) {
	fmt.Println(utils.Now(), id, device, "升级城堡")
	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到城堡升级按钮，退出")
			return
		}
		adb.ScreenOne(device)
		adb.ClickPoint(entity.CB, 2, device)
		adb.ClickPoint_(entity.CB, 3, device)
		img := entity.ChengBaoImg(device)
		if utils.Compare1(img, device) {
			break
		}
	}
	adb.ClickPoint(entity.Shengji, 3, device)
}

func updateChengQiang(id, device string) {
	fmt.Println(utils.Now(), id, device, "升级城墙")
	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到城墙升级按钮，退出")
			return
		}
		checkLianMengLiWu(device)
		adb.ScreenTwo(device)
		adb.ClickPoint(entity.CC, 2, device)
		adb.ClickPoint_(entity.CC, 3, device)
		img := entity.ChengQiangImg(device)
		if utils.Compare1(img, device) {
			break
		}
	}
	adb.ClickPoint(entity.Shengji, 3, device)
}

func updateCangku(id, device string) {
	fmt.Println(utils.Now(), id, device, "升级仓库")
	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到仓库升级按钮，退出")
			return
		}
		adb.ScreenOne(device)
		adb.ClickPoint(entity.Cangku, 2, device)
		adb.ClickPoint_(entity.Cangku, 3, device)
		img := entity.CangKuImg(device)
		if utils.Compare1(img, device) {
			break
		}
	}
	adb.ClickPoint(entity.Shengji, 2, device)
}

func updateZhanZheng(id, device string) {
	fmt.Println(utils.Now(), id, device, "升级战争广场")
	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到战争广场升级按钮，退出")
			return
		}
		adb.ScreenOne(device)
		adb.ClickPoint(entity.ZhanZhen, 2, device)
		adb.ClickPoint_(entity.ZhanZhen, 3, device)
		img := entity.ZhanZhengImg(device)
		if utils.Compare1(img, device) {
			break
		}
	}
	adb.ClickPoint(entity.Shengji, 3, device)
}

func updateYiYuan(id, device string) {
	fmt.Println(utils.Now(), id, device, "升级医院")
	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到医院升级按钮，退出")
			return
		}
		checkLianMengLiWu(device)
		adb.ScreenTwo(device)
		adb.ClickPoint(entity.YiYuan, 2, device)
		adb.ClickPoint_(entity.YiYuan, 3, device)
		img := entity.YiYuanImg(device)
		if utils.Compare1(img, device) {
			break
		}
	}
	adb.ClickPoint(entity.Shengji, 3, device)
}

func updateXueyuan(id, device string) {
	fmt.Println(utils.Now(), id, device, "升级学院")
	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到学院升级按钮，退出")
			return
		}
		checkLianMengLiWu(device)
		adb.ScreenTwo(device)
		adb.ClickPoint(entity.XueYuan, 2, device)
		adb.ClickPoint_(entity.XueYuan, 3, device)
		img := entity.XueYuanImg(device)
		if utils.Compare1(img, device) {
			break
		}
	}
	adb.ClickPoint(entity.Shengji, 3, device)
}

func updateRongLian(id, device string) {
	fmt.Println(utils.Now(), id, device, "升级熔炼")
	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到熔炼升级按钮，退出")
			return
		}
		checkLianMengLiWu(device)
		adb.ScreenTwo(device)
		adb.ClickPoint(entity.RongLian, 2, device)
		adb.ClickPoint_(entity.RongLian, 3, device)
		img := entity.RongLianImg(device)
		if utils.Compare1(img, device) {
			break
		}
	}
	adb.ClickPoint(entity.Shengji, 3, device)
}

func updateShiGuan(id, device string) {
	fmt.Println(utils.Now(), id, device, "升级使馆")
	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到使馆升级按钮，退出")
			return
		}
		checkLianMengLiWu(device)
		adb.ScreenTwo(device)
		adb.ClickPoint(entity.ShiGuan, 2, device)
		adb.ClickPoint_(entity.ShiGuan, 3, device)
		img := entity.ShiGuanImg(device)
		if utils.Compare1(img, device) {
			break
		}
	}
	adb.ClickPoint(entity.Shengji, 3, device)
}

func updateZhenCha(id, device string) {
	fmt.Println(utils.Now(), id, device, "升级侦察")
	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到侦察塔升级按钮，退出")
			return
		}
		adb.ScreenFour(device)
		adb.ClickPoint(entity.ZhenCha, 2, device)
		adb.ClickPoint_(entity.ZhenCha, 3, device)
		img := entity.ZhenChaImg(device)
		if utils.Compare1(img, device) {
			break
		}
	}
	adb.ClickPoint(entity.Shengji, 3, device)
}

func updateBingYing1(id, device string) {
	fmt.Println(utils.Now(), id, device, "升级兵营1")
	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到兵营1升级按钮，退出")
			return
		}
		checkLianMengLiWu(device)
		adb.ScreenTwo(device)
		adb.ClickPoint(entity.Bingyin1, 2, device)
		adb.ClickPoint_(entity.Bingyin1, 3, device)
		img := entity.BingYingImg(device)
		if utils.Compare1(img, device) {
			break
		}
	}
	adb.ClickPoint(entity.Shengji, 3, device)
}

func updateBingYing2(id, device string) {
	fmt.Println(utils.Now(), id, device, "升级兵营2")
	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到兵营2升级按钮，退出")
			return
		}
		adb.ScreenOne(device)
		adb.ClickPoint(entity.Bingyin2, 2, device)
		adb.ClickPoint_(entity.Bingyin2, 3, device)
		img := entity.BingYingImg(device)
		if utils.Compare1(img, device) {
			break
		}
	}
	adb.ClickPoint(entity.Shengji, 3, device)
}

func updateNongChang(id, device string) {
	fmt.Println(utils.Now(), id, device, "升级农场")
	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到农场升级按钮，退出")
			return
		}
		adb.ScreenFive(device)
		adb.ClickPoint(entity.ZY3, 2, device)
		adb.ClickPoint_(entity.ZY3, 3, device)
		img := entity.NongChangImg(device)
		if utils.Compare1(img, device) {
			break
		}
	}
	adb.ClickPoint(entity.Shengji, 3, device)
}

func updateMuChang(id, device string) {
	fmt.Println(utils.Now(), id, device, "升级木场")
	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到木场升级按钮，退出")
			return
		}
		adb.ScreenFive(device)
		adb.ClickPoint(entity.ZY, 2, device)
		adb.ClickPoint_(entity.ZY, 3, device)
		img := entity.MuChangImg(device)
		if utils.Compare1(img, device) {
			break
		}
	}
	adb.ClickPoint(entity.Shengji, 3, device)
}

func updateShiBing(id, device string) {
	fmt.Println(utils.Now(), id, device, "升级士兵训练营")
	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到士兵训练营升级按钮，退出")
			return
		}
		adb.ScreenFive(device)
		adb.ClickPoint(entity.ZY14, 2, device)
		adb.ClickPoint_(entity.ZY14, 3, device)
		img := entity.ShiBingImg(device)
		if utils.Compare1(img, device) {
			break
		}
	}
	adb.ClickPoint(entity.Shengji, 3, device)
}

func updateJiJiu(id, device string) {
	fmt.Println(utils.Now(), id, device, "升级急救帐篷")
	for i := 0; i < 10; i++ {
		if i > 8 {
			adb.Quit(id)
			statuMap[device] = true
			fmt.Println(utils.Now(), "序号", id, "未获取到急救帐篷升级按钮，退出")
			return
		}
		adb.ScreenFive(device)
		adb.ClickPoint(entity.ZY15, 2, device)
		adb.ClickPoint_(entity.ZY15, 3, device)
		img := entity.JiJiuImg(device)
		if utils.Compare1(img, device) {
			break
		}
	}
	adb.ClickPoint(entity.Shengji, 3, device)
}

func updateShiTou(device string) {
	//fmt.Println(utils.Now(), "升级石头")
	//for  {
	//	adb.ScreenSeven(device)
	//	adb.ClickPoint(entity.ZY, 2, device)
	//	adb.ClickPoint_(entity.ZY, 3, device)
	//	img := entity.JiJiuImg(device)
	//	if Compare1(img, device) {
	//		break
	//	}
	//}
	//adb.ClickPoint(entity.Shengji, 3, device)
}

func updateTie(device string) {
	//fmt.Println(utils.Now(), "升级铁")
	//for  {
	//	adb.ScreenSeven(device)
	//	adb.ClickPoint(entity.ZY, 2, device)
	//	adb.ClickPoint_(entity.ZY, 3, device)
	//	img := entity.JiJiuImg(device)
	//	if Compare1(img, device) {
	//		break
	//	}
	//}
	//adb.ClickPoint(entity.Shengji, 3, device)
}

func checkZiYuanBao(device string) bool {
	flag := utils.Compare1(entity.ZiYuanBaoImg(device), device)
	if flag {
		fmt.Println(utils.Now(), device, "点击快速使用资源包")
		flag = utils.Compare1(entity.ZiYuanBaoActImg(device), device)
		if flag {
			fmt.Println(utils.Now(), device, "点击确认按钮")
			adb.ClickPoint(entity.ZiYuanBao, 2, device)
			return true
		} else {
			fmt.Println(utils.Now(), device, "资源包不够，关闭窗口")
			adb.ClickPoint(entity.Close, 2, device)
			adb.ClickPoint(entity.Back, 2, device)
			return false
		}
	}
	return true
}

func checkLianMengLiWu(device string) {
	adb.ClickPoint(entity.P85, 0, device)
	adb.ClickPoint(entity.P85, 2, device)
	if utils.Compare1(entity.Img11(device), device) {
		for {
			if utils.Compare1(entity.Img12(device), device) {
				adb.ClickMore(entity.P84, 2, 2, device)
			} else {
				adb.ClickPoint(entity.Back, 2, device)
				return
			}
		}
	}
}
