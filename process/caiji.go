package process

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"zhanhuo/adb"
	P "zhanhuo/entity"
	"zhanhuo/utils"
)

func Caiji(account *P.Account) {
	closePop(account.Name)
	adb.ClickPoint(P.Reset, 5, account.Name)
	r := lianMengKuang(account.Name)
	if !r {
		puTongKuang(account)
	}
	fmt.Println(utils.Now(), account.Name, "采集完成")
	//account.CaiJiTime = getCaijiTime(account.Name)
	adb.ClickPoint(P.Reset, 5, account.Name)
	closePop(account.Name)
}

func puTongKuang(account *P.Account) {
	device := account.Name
	fmt.Println(utils.Now(), device, "采集普通矿")
	var arr []int
	if account.ChengBao >= 15 {
		arr = rad(4)
	} else if account.ChengBao > 10 {
		arr = rad(3)
	} else {
		arr = rad(2)
	}
	index := 0
	for i := 0; i < 4; i++ {
		adb.ClickPoint(P.Search, 2, device)
		if i >= len(arr) {
			break
		}
		r := arr[i]
		if r == 0 {
			fmt.Println(utils.Now(), device, "采集粮食")
			adb.ClickPoint(P.NC, 1, device)
		} else if r == 1 {
			fmt.Println(utils.Now(), device, "采集木材")
			adb.ClickPoint(P.MC, 1, device)
		} else if r == 2 {
			fmt.Println(utils.Now(), device, "采集石头")
			adb.ClickPoint(P.SC, 1, device)
		} else if r == 3 {
			fmt.Println(utils.Now(), device, "采集铁矿")
			adb.ClickPoint(P.TC, 1, device)
		}
		if index == 0 {
			adb.ClickMore(P.P56, 2, 1, device)
			adb.ClickPoint(P.P62, 1, device)
		}
		if index > 2 {
			fmt.Println(utils.Now(), device, "5次未寻找到资源")
			adb.ClickPoint(P.P62, 1, device)
		}
		if index > 4 {
			fmt.Println(utils.Now(), device, "10次未寻找到资源")
			break
		}
		adb.ClickPoint(P.GoTO, 3, device)
		adb.ClickPoint(P.P55, 2, device)
		adb.ClickPoint(P.P63, 2, device)
		if Compare1(P.Img6(device), device) {
			fmt.Println(utils.Now(), device, "队伍不足")
			adb.ClickPoint(P.Back, 2, device)
			break
		}
		text1 := GetText(device, P.TimeImg(device))
		fmt.Println(utils.Now(), device, "获取行军时间", text1)
		if text1 == "" {
			i = i - 1
			index = index + 1
			continue
		}
		adb.ClickPoint(P.CF, 2, device)
		index = 0
	}
	adb.ClickPoint(P.P55, 2, device)
}

func lianMengKuang(device string) bool {
	fmt.Println(utils.Now(), device, "采集联盟矿")
	list := ChaoKuangList(device)
	for i := 0; i < len(list); i++ {
		adb.ClickPoint(P.P59, 2, device)
		adb.ClickPoint(P.P60, 2, device)
		text := GetText(device, list[i])
		if text == "" {
			list[i].Y = list[i].Y + 30
			text = GetText(device, list[i])
			if text == "" {
				adb.ClickPoint(P.Back, 2, device)
				fmt.Println(utils.Now(), device, "未检测到超级矿")
				return false
			}
		}
		fmt.Println(utils.Now(), device, "开始采集超级矿")

		pp := P.Point{X: list[i].X + 80, Y: list[i].Y + 100}
		adb.ClickPoint(pp, 3, device)
		adb.ClickPoint(P.P61, 2, device)

		if Compare1(P.Img6(device), device) {
			fmt.Println(utils.Now(), device, "队伍不足")
			adb.ClickPoint(P.Back, 2, device)
			return true
		}
		text1 := GetText(device, P.TimeImg(device))
		fmt.Println(utils.Now(), device, "获取行军时间", text1)
		if text1 == "" {
			fmt.Println(utils.Now(), device, "未获取到采集时间")
			return false
		}
		//adb.ClickPoint(P.Back, 2, device)
		fmt.Println(utils.Now(), device, "前往采集")
		adb.ClickPoint(P.CF, 2, device)
	}
	return true
}

func getCaijiTime(device string) string {
	text := GetText(device, P.Img7(device))
	fmt.Println(utils.Now(), device, "获取采集时间", text)
	st := addTime(text)
	fmt.Println(utils.Now(), device, "采集结束时间", st)
	return st
}

func ChaoKuangList(device string) [6]*P.Img {
	var list [6]*P.Img
	list[0] = &P.Img{Name: device + "_chaokuang_1.png", X: 180, Y: 340, W: 200, H: 40}
	list[1] = &P.Img{Name: device + "_chaokuang_2.png", X: 180, Y: 490, W: 200, H: 40}
	list[2] = &P.Img{Name: device + "_chaokuang_3.png", X: 180, Y: 640, W: 200, H: 40}
	list[3] = &P.Img{Name: device + "_chaokuang_4.png", X: 180, Y: 790, W: 200, H: 40}
	list[4] = &P.Img{Name: device + "_chaokuang_5.png", X: 180, Y: 950, W: 200, H: 40}
	list[5] = &P.Img{Name: device + "_chaokuang_6.png", X: 180, Y: 1100, W: 200, H: 40}
	return list
}

func rad(max int) []int {
	var arr []int
	m := make(map[int]int)
	index := 0
	for {

		n, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
		i := int(n.Int64())
		if m[i] == 0 {
			m[i] = 1
			arr = append(arr, i)
			index = index + 1
		}
		if index >= max {
			return arr
		}

	}
}
