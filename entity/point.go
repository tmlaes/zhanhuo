package entity

type Point struct {
	X     int
	Y     int
	X_    int
	Y_    int
	Check bool
}

var (
	//Game 游戏坐标
	Game = Point{X: 640, Y: 120}
	// Reset 重置到左上角坐标
	Reset = Point{X: 70, Y: 1220, Check: true}
	// Close 弹窗关闭按钮坐标
	Close = Point{X: 650, Y: 270}
	// Back 返回按钮坐标
	Back = Point{X: 40, Y: 30}
	// TouXiang 头像坐标
	TouXiang = Point{X: 50, Y: 50}
	// Task 任务列表坐标
	Task = Point{X: 20, Y: 650, Check: true}
	//CB 城堡坐标
	CB = Point{340, 350, 200, 500, true} //升级坐标X-130,y+150
	//CC 城墙坐标
	CC = Point{560, 750, 500, 870, true} //升级坐标X-60,y+120
	//Cangku 仓库坐标
	Cangku = Point{CB.X - 30, CB.Y + 350, CB.X - 40, CB.Y + 470, true}
	//Zhengwu 政务大厅坐标
	Zhengwu = Point{CB.X + 80, CB.Y + 240, CB.X - 100, CB.Y + 310, true}
	//ZhanZhen 战争广场坐标
	ZhanZhen = Point{CB.X - 150, CB.Y + 460, CB.X - 170, CB.Y + 600, true}
	//Bingyin3 3号兵营坐标
	Bingyin3 = Point{CB.X + 80, CB.Y + 480, 0, 0, true}
	////Bingyin4 4号兵营坐标
	Bingyin4 = Point{CB.X + 210, CB.Y + 530, 0, 0, true}
	//Bingyin2 2号兵营坐标
	Bingyin2 = Point{CB.X + 210, CB.Y + 350, CB.X + 200, CB.Y + 470, true}
	//Bingyin1 1号兵营坐标
	Bingyin1 = Point{CC.X - 400, CC.Y - 120, CC.X - 400, CC.Y, true}
	//YiYuan 医院
	YiYuan = Point{CC.X - 230, CC.Y - 30, CC.X - 290, CC.Y + 80, true}
	//XueYuan 学院
	XueYuan = Point{CC.X - 250, CC.Y - 250, CC.X - 260, CC.Y - 160, true}
	//RongLian 熔炼工坊
	RongLian = Point{CC.X - 150, CC.Y - 310, CC.X - 200, CC.Y - 220, true}
	//ShiGuan 使馆
	ShiGuan = Point{CC.X - 340, CC.Y - 200, CC.X - 380, CC.Y - 100, true}
	//MaChe 马车
	MaChe = Point{X: CC.X - 40, Y: CC.Y - 420, Check: true}
	//MaChe 神器
	ShenQi = Point{X: CC.X - 340, Y: CC.Y - 550, Check: true}
	//空地1 最上面的
	Kong1 = Point{X: CC.X - 320, Y: CC.Y - 380, Check: true}
	//空地2 靠近熔炼工坊
	Kong2 = Point{X: CC.X - 230, Y: CC.Y - 330, Check: true}
	//Kong3 靠近城墙
	Kong3 = Point{X: CC.X - 60, Y: CC.Y - 150, Check: true}
	//CL 材料工坊
	CL = Point{X: 560, Y: 700, Check: true}
	//HeiShang 黑商
	HeiShang = Point{X: CL.X - 260, Y: CL.Y + 10, Check: true}
	//ShangChuan 商船
	ShangChuan = Point{X: CL.X - 220, Y: CL.Y - 130, Check: true}
	//WJ 无尽之塔
	WJ = Point{CL.X - 440, CL.Y, CL.X - 380, CL.Y + 100, true}
	//ShenXiang 神像
	ShenXiang = Point{X: CL.X - 410, Y: CL.Y - 220, Check: true}

	//ZhenCha 侦察塔
	ZhenCha = Point{300, 390, 300, 490, true}
	//XingKuang 星矿
	XingKuang = Point{X: 340, Y: 800, Check: true}

	//ZY 资源田1号
	ZY = Point{430, 490, 430, 580, true}
	//ZY2 资源田2号
	ZY2 = Point{ZY.X - 80, ZY.Y + 40, ZY.X - 80, ZY.Y + 130, true}
	//ZY3 资源田3号
	ZY3 = Point{ZY.X, ZY.Y + 70, ZY.X - 10, ZY.Y + 160, true}
	//ZY4 资源田4号
	ZY4 = Point{ZY.X + 70, ZY.Y + 100, ZY.X + 70, ZY.Y + 200, true}
	//ZY5 资源田5号
	ZY5 = Point{ZY.X - 110, ZY.Y + 130, ZY.X - 120, ZY.Y + 230, true}
	//ZY 资源田6号
	ZY6 = Point{ZY.X - 30, ZY.Y + 170, ZY.X - 30, ZY.Y + 260, true}
	//ZY7 资源田7号
	ZY7 = Point{ZY.X - 190, ZY.Y + 180, ZY.X - 190, ZY.Y + 280, true}
	//ZY8 资源田8号
	ZY8 = Point{ZY.X - 100, ZY.Y + 220, ZY.X - 100, ZY.Y + 320, true}
	//ZY9 资源田9号
	ZY9 = Point{ZY.X - 250, ZY.Y + 230, ZY.X - 260, ZY.Y + 330, true}
	//ZY10 资源田10号
	ZY10 = Point{ZY.X - 180, ZY.Y + 270, ZY.X - 180, ZY.Y + 370, true}
	//ZY11 资源田11号
	ZY11 = Point{ZY.X - 330, ZY.Y + 280, 270, 870, true}
	//ZY12 资源田12号
	ZY12 = Point{ZY.X - 250, ZY.Y + 320, ZY.X - 250, ZY.Y + 420, true}
	//ZY13 资源田13号
	ZY13 = Point{ZY.X + 180, ZY.Y + 170, 510, 750, true}
	//ZY14 资源田14号
	ZY14 = Point{ZY.X + 90, ZY.Y + 190, ZY.X + 140, ZY.Y + 280, true}
	//ZY15 资源田15号
	ZY15 = Point{ZY.X + 160, ZY.Y + 230, 500, 810, true}
	//ZY16 资源田16号
	ZY16 = Point{ZY.X + 10, ZY.Y + 240, ZY.X + 60, ZY.Y + 330, true}
	//ZY17 资源田17号
	ZY17 = Point{ZY.X + 110, ZY.Y + 290, ZY.X + 160, ZY.Y + 380, true}
	//ZY18 资源田18号
	ZY18 = Point{190, 780, 240, 870, true}
	//ZYZ1 资源桩
	ZYZ1 = Point{X: ZY.X + 120, Y: ZY.Y + 150, Check: true}
	//ZYZ2 资源桩
	ZYZ2 = Point{X: ZY.X - 30, Y: ZY.Y + 310, Check: true}

	//ZY19 资源田19号
	ZY19 = Point{ZY.X - 60, ZY.Y + 350, ZY.X, ZY.Y + 440, true}
	//ZY20 资源田20号
	ZY20 = Point{ZY.X + 20, ZY.Y + 390, ZY.X + 80, ZY.Y + 380, true}
	//ZY21 资源田21号
	ZY21 = Point{ZY.X - 150, ZY.Y + 390, ZY.X - 90, ZY.Y + 380, true}
	//ZY22 资源田22号
	ZY22 = Point{ZY.X - 70, ZY.Y + 430, ZY.X - 20, ZY.Y + 530, true}
	//ZY23 资源田23号
	ZY23 = Point{ZY.X - 170, ZY.Y + 470, ZY.X - 110, ZY.Y + 570, true}
	//ZY24 资源田24号
	ZY24 = Point{ZY.X + 10, ZY.Y + 470, ZY.X + 70, ZY.Y + 570, true}

	//Task1 任务1
	Task1 = Point{X: 440, Y: 380}
	//Task2 任务2
	Task2 = Point{X: Task1.X, Y: Task1.Y + 50}
	//Task3 任务3
	Task3 = Point{X: Task1.X, Y: Task1.Y + 180}
	//Task4 任务4
	Task4 = Point{X: Task1.X, Y: Task1.Y + 230}
	//Task5 任务5
	Task5 = Point{X: Task1.X, Y: Task1.Y + 280}
	//Task6 任务6
	Task6 = Point{X: Task1.X, Y: Task1.Y + 330}
	//Task7 任务7
	Task7 = Point{X: Task1.X, Y: Task1.Y + 530}

	//TiaojianTZ 城堡里升级条件按钮
	TiaojianTZ = Point{620, 820, 620, 870, false}
	//Shengji 升级按钮
	Shengji = Point{X: 530, Y: 1220}
	//Free 升级按钮
	Free = Point{X: 150, Y: 1220}
	//ZiYuanBao 资源包确认按钮
	ZiYuanBao = Point{X: 350, Y: 940}
	//Search 搜索按钮
	Search = Point{X: 660, Y: 1000, Check: true}
	// 怪物
	GW = Point{X: 600, Y: 1050}
	// 粮食
	NC = Point{X: GW.X - 540, Y: GW.Y}
	//MC 木材
	MC = Point{X: GW.X - 430, Y: GW.Y}
	//SC 石头
	SC = Point{X: GW.X - 320, Y: GW.Y}
	//TC 铁
	TC = Point{X: GW.X - 220, Y: GW.Y}
	//ZC 钻石
	ZC = Point{X: GW.X - 110, Y: GW.Y}
	//ZYL 資源等级
	ZYL = Point{X: 290, Y: 1165}
	//GoTO 前往
	GoTO = Point{X: GW.X, Y: GW.Y + 170}
	// GgW 怪物本体
	GgW = Point{X: 360, Y: 500}
	// GoJi 对话框攻击按钮
	GoJi = Point{X: 350, Y: 850}
	//CF 出发
	CF = Point{X: 600, Y: 1230}

	//BD1 编队1
	BD1 = Point{X: 60, Y: 770}
	//QiBin 骑兵
	QiBin = Point{X: 280, Y: 130}
	//FaShi 法师
	FaShi = Point{X: 630, Y: 130}
	//Left 建造页面左向箭头
	Left = Point{X: 630, Y: 275}

	//创建账号相关坐标
	//贝拉对话框按钮
	P1 = Point{X: 600, Y: 1100}
	//启动游戏授权按钮
	P2 = Point{X: 590, Y: 690}
	//
	P3  = Point{X: 370, Y: 710}
	P4  = Point{X: 570, Y: 1210}
	P5  = Point{X: 320, Y: 680}
	P6  = Point{X: 420, Y: 810}
	P7  = Point{X: 560, Y: 1210}
	P8  = Point{X: 320, Y: 610}
	P9  = Point{X: 350, Y: 730}
	P10 = Point{X: 550, Y: 1210}
	P11 = Point{X: 350, Y: 680}
	P12 = Point{X: 310, Y: 740}
	P13 = Point{X: 320, Y: 710}
	P14 = Point{X: 360, Y: 840}
	P15 = Point{X: 600, Y: 1230}
	P16 = Point{X: 310, Y: 650}
	P17 = Point{X: 380, Y: 630}
	P18 = Point{X: 320, Y: 770}
	P19 = Point{X: 350, Y: 640}
	P20 = Point{X: 240, Y: 530}
	P21 = Point{X: 680, Y: 130}
	P22 = Point{X: 300, Y: 780}
	P23 = Point{X: 200, Y: 660}
	P24 = Point{X: 440, Y: 920}
	P25 = Point{X: 330, Y: 820}
	P26 = Point{X: 340, Y: 830}
	P27 = Point{X: 350, Y: 630}
	P28 = Point{X: 240, Y: 820}
	P29 = Point{X: 530, Y: 1220}
	P30 = Point{X: 50, Y: 950}
	P31 = Point{X: 580, Y: 520}
	P32 = Point{X: 360, Y: 580}
	P33 = Point{X: 310, Y: 700}
	P34 = Point{X: 300, Y: 300}
	P35 = Point{X: 240, Y: 750}
	P36 = Point{X: 200, Y: 1220}
	P37 = Point{X: 620, Y: 820}
	P38 = Point{X: 350, Y: 800}
	P39 = Point{X: 500, Y: 870}
	P40 = Point{X: 50, Y: 200}
	P41 = Point{X: 210, Y: 1240}
	P42 = Point{X: 620, Y: 250}
	P43 = Point{X: 660, Y: 970}
	P44 = Point{X: 640, Y: 360}
	P45 = Point{X: 660, Y: 1090}
	P46 = Point{X: 630, Y: 470}
	P47 = Point{X: 350, Y: 1150}
	P48 = Point{X: 430, Y: 1070}
	P49 = Point{X: 680, Y: 640}
	P50 = Point{X: 360, Y: 1220}
	P51 = Point{X: 670, Y: 170}
	P52 = Point{X: 670, Y: 280}
	P53 = Point{X: 670, Y: 390}
	P54 = Point{X: 640, Y: 230}
	P55 = Point{X: 600, Y: 20}
	//P55 = Point{X:200,Y:75}
	//野外等级+号
	P56 = Point{X: 480, Y: 1215}
	P57 = Point{X: 120, Y: 1230}
	P58 = Point{X: 120, Y: 1150}
	P59 = Point{X: 50, Y: 900}
	P60 = Point{X: 480, Y: 130}
	P61 = Point{X: 510, Y: 450}
	P62 = Point{X: 50, Y: 1220}
	P63 = Point{X: 510, Y: 570}
	P64 = Point{X: 350, Y: 950}
	P65 = Point{X: 670, Y: 800}
	//日常奖励
	P66 = Point{X: 100, Y: 190}
	P67 = Point{X: 190, Y: 190}
	P68 = Point{X: 310, Y: 190}
	P69 = Point{X: 410, Y: 190}
	P70 = Point{X: 510, Y: 190}
	P71 = Point{X: 670, Y: 190}
	P72 = Point{X: 350, Y: 770}
	P73 = Point{X: 350, Y: 890}
	//技能
	P74 = Point{X: 670, Y: 890}
	P75 = Point{X: 130, Y: 720}
	P76 = Point{X: 360, Y: 1230}
	//怪物等级减
	P77 = Point{X: 160, Y: 1220}
	//从任务点击招募
	P78 = Point{350, 580, 460, 680, true}
	//联盟捐献
	P79 = Point{X: 550, Y: 1240}
	P80 = Point{X: 630, Y: 680}
	P81 = Point{X: 100, Y: 430}
	P82 = Point{X: 220, Y: 820}
	//一键捐献按钮
	P83 = Point{X: 230, Y: 700}
	//领取联盟礼物
	P84 = Point{X: 620, Y: 710}
	//滑屏2起点
	P85 = Point{X: 690, Y: 1110}
)
