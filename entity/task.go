package entity

type task struct {
	Name  string
	Level int
}

var level4 = []task{{"bingYing1", 4}, {"bingYing2", 4}}
var level5 = []task{{"chengQiang", 5}, {"xueYuan", 5}, {"shiGuan", 5}, {"rongLian", 5}, {"zhanZhen", 5}, {"kaiHuang1", 1}, {"kaiHuang2", 1}}
var level6 = []task{{"chengQiang", 6}, {"bingYing1", 6}}
var level7 = []task{{"chengQiang", 7}, {"cangKu", 7}}
var level8 = []task{{"chengQiang", 8}, {"bingYing3", 8}, {"shiBingXunLian", 8}}
var level9 = []task{{"chengQiang", 9}, {"jiJiu", 9}}
var level10 = []task{{"chengQiang", 10}, {"xueYuan", 10}}
var level11 = []task{{"chengQiang", 11}, {"yiYuan", 11}}
var level12 = []task{{"chengQiang", 12}, {"xueYuan", 12}, {"kaiHuang3", 1}, {"kaiHuang4", 1}, {"shiTou", 12}}
var level13 = []task{{"chengQiang", 13}, {"bingYing4", 13}, {"shiGuan", 13}}
var level14 = []task{{"chengQiang", 14}, {"zhanZhen", 14}}
var level15 = []task{{"chengQiang", 15}, {"kaiHuang5", 15}, {"zhenCha", 15}}
var level16 = []task{{"chengQiang", 16}, {"faMu", 16}}
var level17 = []task{{"chengQiang", 17}, {"tie", 17}}
var level18 = []task{{"chengQiang", 18}, {"rongLian", 18}}
var level19 = []task{{"chengQiang", 19}, {"bingYing1", 19}, {"xueYuan", 19}}
var TaskMap = make(map[string][]task)

func init() {
	TaskMap["level4"] = level4
	TaskMap["level5"] = level5
	TaskMap["level6"] = level6
	TaskMap["level7"] = level7
	TaskMap["level8"] = level8
	TaskMap["level9"] = level9
	TaskMap["level10"] = level10
	TaskMap["level11"] = level11
	TaskMap["level12"] = level12
	TaskMap["level13"] = level13
	TaskMap["level14"] = level14
	TaskMap["level15"] = level15
	TaskMap["level16"] = level16
	TaskMap["level17"] = level17
	TaskMap["level18"] = level18
	TaskMap["level19"] = level19
}
