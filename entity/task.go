package entity

type task struct {
	Name  string
	Level int
}

var ZhanLiArr [18]int64

var level6 = []task{{"ChengQiang", 6}, {"BingYing1", 6}}
var level7 = []task{{"ChengQiang", 7}, {"CangKu", 7}}
var level8 = []task{{"ChengQiang", 8}, {"BingYing1", 8}, {"ShiBing", 8}}
var level9 = []task{{"ChengQiang", 9}, {"JiJiu", 9}}
var level10 = []task{{"ChengQiang", 10}, {"XueYuan", 10}}
var level11 = []task{{"ChengQiang", 11}, {"YiYuan", 11}}
var level12 = []task{{"ChengQiang", 12}, {"XueYuan", 12}, {"KaiHuang3", 1}, {"KaiHuang4", 1}, {"ShiTou", 12}}
var level13 = []task{{"ChengQiang", 13}, {"BingYing1", 13}, {"ShiGuan", 13}}
var level14 = []task{{"ChengQiang", 14}, {"ZhanZhen", 14}}
var level15 = []task{{"ChengQiang", 15}, {"KaiHuang5", 15}, {"ZhenCha", 15}}
var level16 = []task{{"ChengQiang", 16}, {"MuChang", 16}}
var level17 = []task{{"ChengQiang", 17}, {"Tie", 17}}
var level18 = []task{{"ChengQiang", 18}, {"RongLian", 18}}
var level19 = []task{{"ChengQiang", 19}, {"BingYing1", 19}, {"XueYuan", 19}}

var TaskMap = make(map[string][]task)
var KejiArr = []Point{}
var RollSkillArr = []Point{}

func init() {
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

	ZhanLiArr[0] = 1000
	ZhanLiArr[1] = 1000
	ZhanLiArr[2] = 1000
	ZhanLiArr[3] = 2500
	ZhanLiArr[4] = 3900
	ZhanLiArr[5] = 4600
	ZhanLiArr[6] = 6700
	ZhanLiArr[7] = 8800
	ZhanLiArr[8] = 14200
	ZhanLiArr[9] = 20400
	ZhanLiArr[10] = 29900
	ZhanLiArr[11] = 37400
	ZhanLiArr[12] = 47400
	ZhanLiArr[13] = 57400
	ZhanLiArr[14] = 75000
	ZhanLiArr[15] = 91000
	ZhanLiArr[16] = 110500
	ZhanLiArr[17] = 134000

	for i := 0; i < 11; i++ {
		if i < 4 {
			KejiArr = append(KejiArr, KeJiJunShi1)
			continue
		}
		if i < 7 {
			KejiArr = append(KejiArr, KeJiFaZhan1)
			continue
		}
		if i < 10 {
			KejiArr = append(KejiArr, KeJiFaZhan3)
			continue
		}
	}

	RollSkillArr = append(RollSkillArr, SkillFaZhan1)
	RollSkillArr = append(RollSkillArr, SkillFaZhan2)
	RollSkillArr = append(RollSkillArr, SkillFaZhan3)
	RollSkillArr = append(RollSkillArr, SkillFaZhan4)
	RollSkillArr = append(RollSkillArr, SkillFaZhan5)
	RollSkillArr = append(RollSkillArr, SkillFaZhan6)
	RollSkillArr = append(RollSkillArr, SkillFaZhan7)
	RollSkillArr = append(RollSkillArr, SkillFaZhan8)
	RollSkillArr = append(RollSkillArr, SkillFaZhan9)
	RollSkillArr = append(RollSkillArr, SkillFaZhan11)
	RollSkillArr = append(RollSkillArr, SkillFaZhan12)
	RollSkillArr = append(RollSkillArr, SkillFaZhan14)
	RollSkillArr = append(RollSkillArr, SkillFaZhan15)
	RollSkillArr = append(RollSkillArr, SkillFaZhan17)
	RollSkillArr = append(RollSkillArr, SkillFaZhan20)
	RollSkillArr = append(RollSkillArr, SkillFaZhan21)
	RollSkillArr = append(RollSkillArr, SkillFaZhan23)
	RollSkillArr = append(RollSkillArr, SkillFaZhan24)
	RollSkillArr = append(RollSkillArr, SkillFaZhan25)
	RollSkillArr = append(RollSkillArr, SkillZhanZheng1)
	RollSkillArr = append(RollSkillArr, SkillZhanZheng2)
	RollSkillArr = append(RollSkillArr, SkillZhanZheng3)
	RollSkillArr = append(RollSkillArr, SkillZhanZheng4)
}
