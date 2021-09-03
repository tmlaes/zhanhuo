package entity

import "reflect"

type Account struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	ChengBao       int    `json:"chengBao"`
	ChengQiang     int    `json:"chengQiang"`
	CangKu         int    `json:"cangKu"`
	ZhanZhen       int    `json:"zhanZhen"`
	YiYuan         int    `json:"yiYuan"`
	XueYuan        int    `json:"xueYuan"`
	RongLian       int    `json:"rongLian"`
	ShiGuan        int    `json:"shiGuan"`
	ZhenCha        int    `json:"zhenCha"`
	BingYing1      int    `json:"bingYing1"`
	BingYing2      int    `json:"bingYing2"`
	BingYing3      int    `json:"bingYing3"`
	BingYing4      int    `json:"bingYing4"`
	NongChang      int    `json:"nongChang"`
	MuChang        int    `json:"muChang"`
	ShiTou         int    `json:"shiTou"`
	Tie            int    `json:"tie"`
	ShiBing        int    `json:"shiBing"`
	JiJiu          int    `json:"jiJiu"`
	GuaiWu         int    `json:"guaiWu"`
	BuildTime      string `json:"buildTime"`
	LastUpdateTime string `json:"lastUpdateTime"`
	FengShouTime   string `json:"fengShouTime"`
}

type BingYingVO struct {
	Id      int    `json:"id"`
	Level   int    `json:"level"`
	ZmTimes string `json:"zmTimes"`
}

type ZiYuanVO struct {
	Id    int `json:"id"`
	Level int `json:"level"`
}

func (acc *Account) SetField(field string, value interface{}) {
	v := reflect.ValueOf(acc)
	v = v.Elem()
	byName := v.FieldByName(field)
	if byName.IsValid() {
		switch value.(type) {
		case string:
			vv, _ := value.(string)
			byName.SetString(vv)
		case int:
			vv, _ := value.(int)
			byName.SetInt(int64(vv))
		}
	}
}

func (acc *Account) GetFieldInt(field string) int {
	v := reflect.ValueOf(acc)
	v = v.Elem()
	byName := v.FieldByName(field)
	if byName.IsValid() {
		return byName.Interface().(int)
	}
	return 0
}

func (acc *Account) GetFieldString(field string) string {
	v := reflect.ValueOf(acc)
	v = v.Elem()
	byName := v.FieldByName(field)
	if byName.IsValid() {
		return byName.Interface().(string)
	}
	return ""
}
