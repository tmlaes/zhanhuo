package process

import (
	"fmt"
	"time"
	"zhanhuo/adb"
	P "zhanhuo/entity"
)

func CaiGi(account *P.Account) {
	device := account.Name
	adb.Click(P.Reset.X, P.Reset.Y, 5, device)
	fmt.Println(time.Now(), device, "点击搜索按钮")
	adb.Click(P.Search.X, P.Search.Y, 2, device)
	fmt.Println(time.Now(), device, "点击怪物图标")
	adb.Click(P.NC.X, P.NC.Y, 1, device)

	fmt.Println(time.Now(), device, "点击等级输入框")
	adb.Click(P.GWL.X, P.GWL.Y, 1, device)
	fmt.Println(time.Now(), device, "输入等级")
	adb.Text(device, "5")
	fmt.Println(time.Now(), device, "回车")
	adb.Key(device, "4")
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println(time.Now(), device, "点击前往")
	adb.Click(P.GoTO.X, P.GoTO.Y, 1, device)
	fmt.Println(time.Now(), device, "关闭前往对话框")
	adb.Click(P.CloseW.X, P.CloseW.Y, 1, device)
	adb.Click(P.GoJi.X, P.GoJi.Y, 3, device)
	adb.Click(P.CF.X, P.CF.Y, 3, device)
}
