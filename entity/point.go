package entity

type Point struct {
	X  int
	Y  int
	X_ int
	Y_ int
}

var (
	// Reset 重置到左上角坐标
	Reset = Point{X: 70, Y: 1220}
	// Close 弹窗关闭按钮坐标
	Close = Point{X: 650, Y: 270}
	// Back 返回按钮坐标
	Back = Point{X: 40, Y: 30}
	// TouXiang 头像坐标
	TouXiang = Point{X: 50, Y: 50}
	// Task 任务列表坐标
	Task = Point{X: 20, Y: 650}
	// One 第一屏坐标
	One = Point{0, 0, 180, 30}
	// Three 第三屏坐标
	Three = Point{200, 0, 0, 0}
	//Four 第四屏坐标
	Four = Point{400, 0, 0, 20}
	//Five 第五屏坐标
	Five = Point{150, 150, 0, 0}
	//CB 城堡坐标
	CB = Point{280, 330, 140, 480} //升级坐标X-130,y+150
	//CC 城墙坐标
	CC = Point{560, 750, 500, 870} //升级坐标X-60,y+120
	//Cangku 仓库坐标
	Cangku = Point{CB.X - 30, CB.Y + 350, CB.X - 40, CB.Y + 470}
	//Zhengwu 政务大厅坐标
	Zhengwu = Point{CB.X + 80, CB.Y + 240, CB.X - 100, CB.Y + 310}
	//ZhanZhen 战争广场坐标
	ZhanZhen = Point{CB.X - 150, CB.Y + 460, 280, 780}
	//Bingyin3 3号兵营坐标
	Bingyin3 = Point{CB.X + 80, CB.Y + 480, 0, 0}
	////Bingyin4 4号兵营坐标
	Bingyin4 = Point{CB.X + 210, CB.Y + 530, 0, 0}
	//Bingyin2 2号兵营坐标
	Bingyin2 = Point{CB.X + 210, CB.Y + 350, CB.X + 200, CB.Y + 470}
	//Bingyin1 1号兵营坐标
	Bingyin1 = Point{CC.X - 400, CC.Y - 120, CC.X - 400, CB.Y}
	//YiYuan 医院
	YiYuan = Point{CC.X - 230, CC.Y - 30, CC.X - 290, CC.Y + 80}
	//XueYuan 学院
	XueYuan = Point{CC.X - 200, CC.Y - 270, CC.X - 260, CC.Y - 160}
	//RongLian 熔炼工坊
	RongLian = Point{CC.X - 150, CC.Y - 310, CC.X - 200, CC.Y - 220}
	//ShiGuan 使馆
	ShiGuan = Point{CC.X - 340, CC.Y - 200, CC.X - 380, CC.Y - 100}
	//MaChe 马车
	MaChe = Point{X: CC.X - 40, Y: CC.Y - 430}
	//MaChe 马车
	ShenQi = Point{X: CC.X - 340, Y: CC.Y - 550}
	//空地1 最上面的
	Kong1 = Point{X: CC.X - 320, Y: CC.Y - 380}
	//空地2 靠近熔炼工坊
	Kong2 = Point{X: CC.X - 230, Y: CC.Y - 330}
	//Kong3 靠近城墙
	Kong3 = Point{X: CC.X - 60, Y: CC.Y - 150}
	//WJ 无尽之塔
	WJ = Point{330, 780, 390, 880}
	//ShangChuan 商船
	ShangChuan = Point{X: WJ.X + 160, Y: WJ.Y - 290}
	//HeiShang 黑商
	HeiShang = Point{X: WJ.X + 70, Y: CL.Y - 180}
	//CL 材料工坊
	CL = Point{X: 100, Y: 700}
	//XingKuang 星矿
	XingKuang = Point{X: CL.X + 200, Y: CL.Y + 110}
	//ZhenCha 侦察塔
	ZhenCha = Point{CL.X + 170, CL.Y - 300, CL.X + 170, CL.Y - 210}
	//ShenXiang 神像
	ShenXiang = Point{X: CL.X - 20, Y: CL.Y - 470}

	//ZY 资源田1号
	ZY = Point{460, 490, 510, 580}
	//ZY2 资源田2号
	ZY2 = Point{ZY.X - 80, ZY.Y + 40, ZY.X - 30, ZY.Y + 120}
	//ZY3 资源田3号
	ZY3 = Point{ZY.X, ZY.Y + 70, ZY.X + 50, ZY.Y + 150}
	//ZY4 资源田4号
	ZY4 = Point{ZY.X + 80, ZY.Y + 100, ZY.X + 130, ZY.Y + 180}
	//ZY5 资源田5号
	ZY5 = Point{ZY.X - 110, ZY.Y + 130, ZY.X - 60, ZY.Y + 210}
	//ZY 资源田6号
	ZY6 = Point{ZY.X - 20, ZY.Y + 170, ZY.X + 30, ZY.Y + 250}
	//ZY7 资源田7号
	ZY7 = Point{ZY.X - 180, ZY.Y + 190, ZY.X - 130, ZY.Y + 240}
	//ZY8 资源田8号
	ZY8 = Point{ZY.X - 90, ZY.Y + 220, ZY.X - 40, ZY.Y + 300}
	//ZY9 资源田9号
	ZY9 = Point{ZY.X - 250, ZY.Y + 230, ZY.X - 200, ZY.Y + 310}
	//ZY10 资源田10号
	ZY10 = Point{ZY.X - 170, ZY.Y + 270, ZY.X - 120, ZY.Y + 350}
	//ZY11 资源田11号
	ZY11 = Point{ZY.X - 325, ZY.Y + 285, ZY.X - 275, ZY.Y + 365}
	//ZY12 资源田12号
	ZY12 = Point{ZY.X - 240, ZY.Y + 320, ZY.X - 190, ZY.Y + 400}
	//ZY13 资源田13号
	ZY13 = Point{ZY.X - 140, ZY.Y + 390, ZY.X - 90, ZY.Y + 470}
	//ZY 资源田1号
	ZY14 = Point{ZY.X - 50, ZY.Y + 350, ZY.X, ZY.Y + 430}
	//ZY 资源田1号
	ZY15 = Point{ZY.X + 30, ZY.Y + 390, ZY.X + 80, ZY.Y + 470}
	//ZY 资源田1号
	ZY16 = Point{ZY.X - 60, ZY.Y + 430, ZY.X - 10, ZY.Y + 510}
	//ZY 资源田17号
	ZY17 = Point{ZY.X + 20, ZY.Y + 470, ZY.X + 70, ZY.Y + 550}
	//ZYZ1 资源桩
	ZYZ1 = Point{X: ZY.X + 120, Y: ZY.Y + 150}
	//ZYZ2 资源桩
	ZYZ2 = Point{X: ZY.X - 20, Y: ZY.Y + 300}
	//ZY18 资源田18号
	ZY18 = Point{ZY.X + 90, ZY.Y + 190, ZY.X + 140, ZY.Y + 270}
	//ZY19 资源田19号
	ZY19 = Point{ZY.X + 20, ZY.Y + 240, ZY.X + 70, ZY.Y + 320}
	//ZY20 资源田20号
	ZY20 = Point{ZY.X + 120, ZY.Y + 290, ZY.X + 170, ZY.Y + 370}
	//ZY21 资源田21号
	ZY21 = Point{ZY.X + 180, ZY.Y + 170, 510, 750}
	//ZY22 资源田22号
	ZY22 = Point{ZY.X + 160, ZY.Y + 230, 500, 810}
	//ZY23 资源田23号
	ZY23 = Point{ZY.X + 220, ZY.Y + 290, 520, 860}

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
	TiaojianTZ = Point{620, 820, 620, 870}
	//Shengji 升级按钮
	Shengji = Point{X: 530, Y: 1220}
	//ZiYuanBao 资源包确认按钮
	ZiYuanBao = Point{X: 350, Y: 940}
	//Search 搜索按钮
	Search = Point{X: 660, Y: 1000}
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
	//GWL 怪物等级
	GWL = Point{X: GW.X - 270, Y: GW.Y + 110}
	//GoTO 前往
	GoTO = Point{X: GW.X, Y: GW.Y + 170}
	// GgW 怪物本体
	GgW = Point{X: 360, Y: 500}
	// CloseW 关闭前往对话框
	CloseW = Point{X: 150, Y: 70}
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
)
