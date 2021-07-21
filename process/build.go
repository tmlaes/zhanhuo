package process

import (
	"zhanhuo/adb"
	P "zhanhuo/entity"
)

type Build struct {
	Name string
	p    P.Point
}

func (b *Build) Create(device string) {
	adb.Click(b.p.X, b.p.Y, 2, device)
	for {
		text := GetText1(device, P.Task1Img(device))
		if text == b.Name {
			break
		}
		//点击左向箭头
		//adb.Click(p.X, p.Y, 2, device)
	}
	//点击建造
	//点击免费按钮4次
	//点击返回

}

func (b *Build) Update(device string) {
	adb.Click(b.p.X, b.p.Y, 2, device)
	for {
		//校验是否进入升级界面
	}
	//点击升级

}
