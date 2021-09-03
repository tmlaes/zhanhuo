package entity

type Img struct {
	Name       string
	Path       string
	X, Y, W, H int
}

type OcrImg struct {
	Name       string
	Path       string
	X, Y, W, H int
	RGB        [3][2]uint8
}

func StartImg(deviceName string) *Img {
	return &Img{
		deviceName + "_beila.png",
		"images/beila.png",
		330, 1050, 55, 30,
	}
}

func LevelImg(deviceName string) *Img {
	return &Img{
		deviceName + "_level.png",
		"",
		250, 550, 220, 30,
	}
}

func Task1Img(deviceName string) *OcrImg {
	return &OcrImg{
		deviceName + "_task.png",
		"",
		220, 360, 90, 40,
		[3][2]uint8{{28, 40}, {40, 50}, {40, 55}},
	}
}
func TiLiImg(deviceName string) *Img {
	return &Img{
		deviceName + "_tili.png",
		"",
		150, 10, 40, 30,
	}
}

func TimeImg(deviceName string) *Img {
	return &Img{
		deviceName + "_time.png",
		"",
		540, 1177, 100, 25,
	}
}
func MasterImg(deviceName string) *Img {
	return &Img{
		deviceName + "_pic_0.png",
		"images/pic_0.png",
		20, 20, 80, 80,
	}
}

func CloseImg(deviceName string) *Img {
	return &Img{
		deviceName + "_close.png",
		"images/close.png",
		630, 250, 40, 30,
	}
}

func Close1Img(deviceName string) *Img {
	return &Img{
		deviceName + "_close1.png",
		"images/close1.png",
		640, 210, 20, 30,
	}
}
func Close2Img(deviceName string) *Img {
	return &Img{
		deviceName + "_close2.png",
		"images/close2.png",
		615, 250, 30, 30,
	}
}

func ZiYuanBaoImg(deviceName string) *Img {
	return &Img{
		deviceName + "_ziyuanbao.png",
		"images/ziyuanbao.png",
		265, 265, 190, 30,
	}
}

func ZiYuanBaoActImg(deviceName string) *Img {
	return &Img{
		deviceName + "_ziyuanbaoact.png",
		"images/ziyuanbaoact.png",
		240, 910, 240, 50,
	}
}

func ChengBaoImg(deviceName string) *Img {
	return &Img{
		deviceName + "_chengbao.png",
		"images/build/chengbao.png",
		330, 20, 60, 30,
	}
}

func ChengQiangImg(deviceName string) *Img {
	return &Img{
		deviceName + "_chengqiang.png",
		"images/build/chengqiang.png",
		330, 20, 60, 30,
	}
}

func XueYuanImg(deviceName string) *Img {
	return &Img{
		deviceName + "_xueyuan.png",
		"images/build/xueyuan.png",
		330, 20, 60, 30,
	}
}

func RongLianImg(deviceName string) *Img {
	return &Img{
		deviceName + "_ronglian.png",
		"images/build/ronglian.png",
		300, 20, 120, 30,
	}
}

func ShiGuanImg(deviceName string) *Img {
	return &Img{
		deviceName + "_shiguan.png",
		"images/build/shiguan.png",
		330, 20, 60, 30,
	}
}

func BingYing0Img(deviceName string) *Img {
	return &Img{
		deviceName + "_bingying0.png",
		"images/build/bingying0.png",
		330, 20, 60, 30,
	}
}

func BingYingImg(deviceName string) *Img {
	return &Img{
		deviceName + "_bingying.png",
		"images/build/bingying.png",
		330, 20, 60, 30,
	}
}

func XieYiImg(deviceName string) *Img {
	return &Img{
		deviceName + "_xieyi.png",
		"images/xieyi.png",
		290, 370, 130, 30,
	}
}

func ZhanZhengImg(deviceName string) *Img {
	return &Img{
		deviceName + "_zhangzheng.png",
		"images/build/zhangzheng.png",
		300, 20, 120, 30,
	}
}

func CangKuImg(deviceName string) *Img {
	return &Img{
		deviceName + "_cangku.png",
		"images/build/cangku.png",
		330, 20, 60, 30,
	}
}

func YiYuanImg(deviceName string) *Img {
	return &Img{
		deviceName + "_yiyuan.png",
		"images/build/yiyuan.png",
		330, 20, 60, 30,
	}
}

func ZhenChaImg(deviceName string) *Img {
	return &Img{
		deviceName + "_zhencha.png",
		"images/build/zhencha.png",
		310, 20, 100, 30,
	}
}

func NongChangImg(deviceName string) *Img {
	return &Img{
		deviceName + "_nongchang.png",
		"images/build/nongchang.png",
		330, 20, 60, 30,
	}
}

func MuChangImg(deviceName string) *Img {
	return &Img{
		deviceName + "_muchang.png",
		"images/build/muchang.png",
		310, 20, 100, 30,
	}
}

func ShiBingImg(deviceName string) *Img {
	return &Img{
		deviceName + "_shibing.png",
		"images/build/shibing.png",
		280, 20, 160, 30,
	}
}

func JiJiuImg(deviceName string) *Img {
	return &Img{
		deviceName + "_jijiu.png",
		"images/build/jijiu.png",
		300, 20, 120, 30,
	}
}

func TaskGoImg(deviceName string) *Img {
	return &Img{
		deviceName + "_taskgo.png",
		"images/taskgo.png",
		320, 1160, 70, 40,
	}
}

func ZhaoMuImg(deviceName string) *Img {
	return &Img{
		deviceName + "_zhaomu.png",
		"images/zhaomu.png",
		330, 20, 60, 30,
	}
}

func Img2(deviceName string) *Img {
	return &Img{
		deviceName + "_pic_2.png",
		"images/pic_2.png",
		220, 260, 120, 40,
	}
}

func Img3(deviceName string) *Img {
	return &Img{
		deviceName + "_pic_3.png",
		"images/pic_3.png",
		320, 840, 80, 30,
	}
}

func Img4(deviceName string) *Img {
	return &Img{
		deviceName + "_pic_4.png",
		"",
		450, 100, 150, 30,
	}
}
func Img5(deviceName string) *Img {
	return &Img{
		deviceName + "_pic_5.png",
		"",
		170, 60, 100, 27,
	}
}

func Img6(deviceName string) *Img {
	return &Img{
		deviceName + "_pic_6.png",
		"images/pic_6.png",
		300, 480, 120, 30,
	}
}

func Img7(deviceName string) *Img {
	return &Img{
		deviceName + "_pic_7.png",
		"",
		60, 240, 110, 20,
	}
}

func Img8(deviceName string) *Img {
	return &Img{
		deviceName + "_pic_8.png",
		"images/pic_8.png",
		300, 930, 120, 30,
	}
}

func Img9(deviceName string) *Img {
	return &Img{
		deviceName + "_pic_9.png",
		"images/pic_9.png",
		300, 840, 120, 30,
	}
}

func JiangLiImg(deviceName string) *Img {
	return &Img{
		deviceName + "_jiangli.png",
		"images/jiangli.png",
		300, 20, 120, 30,
	}
}

func BackImg(deviceName string) *Img {
	return &Img{
		deviceName + "_back.png",
		"images/back.png",
		20, 10, 60, 50,
	}
}

func ZhaoMuBY1KongImg(deviceName string) *Img {
	return &Img{
		deviceName + "_zhao_by1_kong.png",
		"images/task/zhao_by1_kong.png",
		220, 550, 75, 25,
	}
}
func ZhaoMuBY1ShouImg(deviceName string) *Img {
	return &Img{
		deviceName + "_zhao_by1_shou.png",
		"images/task/zhao_by1_shou.png",
		220, 550, 75, 25,
	}
}

func ZhaoMuBY2KongImg(deviceName string) *Img {
	return &Img{
		deviceName + "_zhao_by2_kong.png",
		"images/task/zhao_by2_kong.png",
		220, 600, 75, 25,
	}
}

func ZhaoMuBY2ShouImg(deviceName string) *Img {
	return &Img{
		deviceName + "_zhao_by2_shou.png",
		"images/task/zhao_by2_shou.png",
		220, 600, 75, 25,
	}
}

func ZhaoMuBY3KongImg(deviceName string) *Img {
	return &Img{
		deviceName + "_zhao_by3_kong.png",
		"images/task/zhao_by3_kong.png",
		220, 650, 75, 25,
	}
}

func ZhaoMuBY3ShouImg(deviceName string) *Img {
	return &Img{
		deviceName + "_zhao_by3_shou.png",
		"images/task/zhao_by3_shou.png",
		220, 650, 75, 25,
	}
}
func TaskBuild1Img(deviceName string) *Img {
	return &Img{
		deviceName + "_task_build1.png",
		"images/task/task_build1.png",
		220, 370, 75, 25,
	}
}

func Img10(deviceName string) *Img {
	return &Img{
		deviceName + "_pic_10.png",
		"images/pic_10.png",
		460, 680, 60, 30,
	}
}

func Img11(deviceName string) *Img {
	return &Img{
		deviceName + "_pic_11.png",
		"images/pic_11.png",
		290, 20, 140, 30,
	}
}

func Img12(deviceName string) *Img {
	return &Img{
		deviceName + "_pic_12.png",
		"images/pic_12.png",
		595, 705, 45, 20,
	}
}

func JiangLiTaskImg(deviceName string) *Img {
	return &Img{
		deviceName + "_jianglitask.png",
		"images/jianglitask.png",
		330, 20, 60, 30,
	}
}
