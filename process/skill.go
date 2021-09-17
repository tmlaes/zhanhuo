package process

import (
	"fmt"
	"time"
	"zhanhuo/adb"
	"zhanhuo/entity"
	"zhanhuo/utils"
)

func Skill(account *entity.Account) {
	xueYuanSkill(account)
	roleSkill(account)
}

func xueYuanSkill(account *entity.Account) {
	device := account.Name
	if account.Keji >= len(entity.KejiArr) {
		fmt.Println(utils.Now(), device, "科技队列超出范围")
		return
	}
	adb.ClickMore(entity.P55, 3, 1, device)
	adb.ScreenZero(device)
	for i := 0; i < 4; i++ {
		task(device)
		if adb.Compare1(entity.Img13(device), device) {
			break
		}
		adb.ClickPoint(entity.P55, 2, device)
		if i >= 3 {
			fmt.Println(utils.Now(), device, "科技队列不处于空闲状态")
			return
		}
	}

	adb.ClickPoint(entity.P86, 2, device)
	p := entity.KejiArr[account.Keji]

	adb.ClickPoint(p, 1, device)
	adb.ClickPoint_(p, 2, device)
	adb.ClickPoint(entity.P87, 2, device)
	checkZiYuanBao(device)
	account.Keji = account.Keji + 1
	utils.WriteJons(account)
	adb.ClickPoint(entity.Back, 2, device)
	adb.ClickPoint(entity.Back, 2, device)
}

func roleSkill(account *entity.Account) {
	t, _ := time.ParseInLocation(utils.DATE_FORMAT, account.DateTime, time.Local)
	sub := time.Now().Sub(t)
	if sub.Hours() < 24 {
		return
	}
	account.DateTime = utils.Now()
	defer utils.WriteJons(account)
	device := account.Name
	adb.ClickMore(entity.P88, 2, 2, device)
	defer adb.ClickMore(entity.Back, 2, 2, device)
	if adb.Compare1(entity.Img1(device), device) {
		return
	}
	skill := account.Skill
	screenRoll(skill, device)
	points := entity.RollSkillArr[skill:]
	for i := 0; i < len(points); i++ {
		adb.ClickPoint(points[i], 2, device)
		adb.ClickPoint(entity.P90, 2, device)
		if adb.Compare1(entity.Img1(device), device) {
			account.Skill = skill + i
			return
		}
	}

}

func screenRoll(skill int, device string) {
	if skill > 20 {
		return
	}
	adb.ClickPoint(entity.P89, 1, device)
	if skill > 5 {
		adb.ScreenRollSkill(device)
	}
	if skill > 11 {
		adb.ScreenRollSkill(device)
	}
	if skill > 17 {
		adb.ScreenRollSkill(device)
	}
}
