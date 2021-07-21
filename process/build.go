package process

import (
	"fmt"
	"time"
	"zhanhuo/adb"
	P "zhanhuo/entity"
)

type Build struct {
	Name  string
	index int
	p     P.Point
}

func (b *Build) Create(device string) {
	adb.Scroll(b.index, device)
	adb.Click(b.p.X, b.p.Y, 3, device)
	for {
		shotImg := adb.Screenshot(device)
		img := P.BuildImg(device)
		CutImage(shotImg, img)
		oldImg := "./images/build/" + b.Name + ".png"
		flag := Compare(img.Name, oldImg)
		if flag {
			break
		}
		fmt.Println(time.Now(), device, "点击左向箭头")
		adb.Click(P.Left.X, P.Left.Y, 2, device)
	}
	fmt.Println(time.Now(), device, "点击免费建造按钮")
	adb.Click(P.Free.X, P.Free.Y, 3, device)
	fmt.Println(time.Now(), device, "点击建筑")
	adb.Click(b.p.X, b.p.Y, 2, device)
	fmt.Println(time.Now(), device, "点击升级")
	//TODO 这里需要确认是否进入升级界面
	adb.Click(b.p.X_, b.p.Y_, 3, device)
	fmt.Println(time.Now(), device, "点击3次免费建造按钮")
	adb.ClickMore(P.Free.X, P.Free.Y, 3, 1, device)
	fmt.Println(time.Now(), device, "点击升级按钮")
	adb.Click(P.Shengji.X, P.Shengji.Y, 3, device)
	//点击返回
	//adb.Click(P.Back.X, P.Back.Y,3, device)
}

func (b *Build) Update(device string) {
	for {
		adb.Scroll(b.index, device)
		fmt.Println(time.Now(), device, "点击建筑")
		adb.Click(b.p.X, b.p.Y, 2, device)
		fmt.Println(time.Now(), device, "点击升级")
		adb.Click(b.p.X_, b.p.Y_, 3, device)
		shotImg := adb.Screenshot(device)
		img := P.UpdateImg(device)
		CutImage(shotImg, img)
		flag := Compare(img.Name, UpdateImg)
		if flag {
			break
		}
	}
	fmt.Println(time.Now(), device, "点击升级按钮")
	adb.Click(P.Shengji.X, P.Shengji.Y, 3, device)
	CheckZiYuanBao(device)
}

func CheckZiYuanBao(device string) {
	shotImg := adb.Screenshot(device)
	img := P.ZiYuanBaoImg(device)
	CutImage(shotImg, img)
	flag := Compare(img.Name, ZiYuanBaoImg)
	if flag {
		fmt.Println(time.Now(), device, "点击快速使用资源包")
		adb.Click(P.ZiYuanBao.X, P.ZiYuanBao.Y, 2, device)
	}
}
