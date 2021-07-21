package entity

type Account struct {
	Name            string     `json:"name"`
	ChengBaoLevel   int        `json:"chengBaoLevel"`
	ChengQiangLevel int        `json:"chengQiangLevel"`
	CangKuLevel     int        `json:"cangKuLevel"`
	ZhanZhenLevel   int        `json:"zhanZhenLevel"`
	YiYuanLevel     int        `json:"yiYuanLevel"`
	XueYuanLevel    int        `json:"xueYuanLevel"`
	RongLianLevel   int        `json:"rongLianLevel"`
	ShiGuanLevel    int        `json:"shiGuanLevel"`
	ZhenChaLevel    int        `json:"zhenChaLevel"`
	BingYing        []bingYing `json:"bingYing"`
	ZiYuan          ziYuan     `json:"ziYuan"`
	KaiHuang1       int        `json:"kaiHuang1"`
	KaiHuang2       int        `json:"kaiHuang2"`
	KaiHuang3       int        `json:"kaiHuang3"`
	KaiHuang4       int        `json:"kaiHuang4"`
	KaiHuang5       int        `json:"kaiHuang5"`
}

type bingYing struct {
	Level   int    `json:"level"`
	ZmTimes string `json:"zmTimes"`
}
type ziYuan struct {
	NongChang []tian `json:"nongChang"`
	MuChang   []tian `json:"muChang"`
	ShiTou    []tian `json:"shiTou"`
	Tie       []tian `json:"tie"`
	JiJiu     []tian `json:"jiJiu"`
	ShiBing   []tian `json:"shiBing"`
}

type tian struct {
	Id    int `json:"id"`
	Level int `json:"level"`
}

func (a *Account) GetLevel(name string) (level int) {
	switch name {
	case "chengQiang":
		level = a.ChengQiangLevel
	case "bingYing1":
		level = a.BingYing[0].Level
	case "bingYing2":
		level = a.BingYing[1].Level
	case "xueYuan":
		level = a.XueYuanLevel
	case "shiGuan":
		level = a.ShiGuanLevel
	case "rongLian":
		level = a.RongLianLevel
	case "zhanZhen":
		level = a.ZhanZhenLevel
	case "kaiHuang1":
		level = a.KaiHuang1
	case "kaiHuang2":
		level = a.KaiHuang2
	case "cangKu":
		level = a.CangKuLevel
	case "bingYing3":
		level = a.BingYing[2].Level
	case "shiBing":
		level = a.ZiYuan.ShiBing[0].Level
	case "jiJiu":
		level = a.ZiYuan.JiJiu[0].Level
	case "yiYuan":
		level = a.YiYuanLevel
	case "kaiHuang3":
		level = a.KaiHuang3
	case "kaiHuang4":
		level = a.KaiHuang4
	case "shiTou":
		level = a.ZiYuan.ShiTou[0].Level
	case "bingYing4":
		level = a.BingYing[3].Level
	case "kaiHuang5":
		level = a.KaiHuang5
	case "zhenCha":
		level = a.ZhenChaLevel
	case "faMu":
		level = a.ZiYuan.MuChang[0].Level
	case "tie":
		level = a.ZiYuan.Tie[0].Level
	}
	return level
}

func (a *Account) SetLevel(name string) {
	switch name {
	case "chengQiang":
		a.ChengQiangLevel = a.ChengQiangLevel + 1
	case "bingYing1":
		a.BingYing[0].Level = a.BingYing[0].Level + 1
	case "bingYing2":
		a.BingYing[1].Level = a.BingYing[1].Level + 1
	case "xueYuan":
		a.XueYuanLevel = a.XueYuanLevel + 1
	case "shiGuan":
		a.ShiGuanLevel = a.ShiGuanLevel + 1
	case "rongLian":
		a.RongLianLevel = a.RongLianLevel + 1
	case "zhanZhen":
		a.ZhanZhenLevel = a.ZhanZhenLevel + 1
	case "kaiHuang1":
		a.KaiHuang1 = a.KaiHuang1 + 1
	case "kaiHuang2":
		a.KaiHuang2 = a.KaiHuang2 + 1
	case "cangKu":
		a.CangKuLevel = a.CangKuLevel + 1
	case "bingYing3":
		a.BingYing[2].Level = a.BingYing[2].Level + 1
	case "shiBing":
		a.ZiYuan.ShiBing[0].Level = a.ZiYuan.ShiBing[0].Level + 1
	case "jiJiu":
		a.ZiYuan.JiJiu[0].Level = a.ZiYuan.JiJiu[0].Level + 1
	case "yiYuan":
		a.YiYuanLevel = a.YiYuanLevel + 1
	case "kaiHuang3":
		a.KaiHuang3 = a.KaiHuang3 + 1
	case "kaiHuang4":
		a.KaiHuang4 = a.KaiHuang4 + 1
	case "shiTou":
		a.ZiYuan.ShiTou[0].Level = a.ZiYuan.ShiTou[0].Level + 1
	case "bingYing4":
		a.BingYing[3].Level = a.BingYing[3].Level + 1
	case "kaiHuang5":
		a.KaiHuang5 = a.KaiHuang5 + 1
	case "zhenCha":
		a.ZhenChaLevel = a.ZhenChaLevel + 1
	case "faMu":
		a.ZiYuan.MuChang[0].Level = a.ZiYuan.MuChang[0].Level + 1
	case "tie":
		a.ZiYuan.Tie[0].Level = a.ZiYuan.Tie[0].Level + 1
	}
}
