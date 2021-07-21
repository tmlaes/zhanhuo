package entity

type Img struct {
	Name       string
	X, Y, W, H int
}

func LevelImg(deviceName string) *Img {
	return &Img{
		deviceName + "_level.png",
		250, 550, 220, 30,
	}
}

func Task1Img(deviceName string) *Img {
	return &Img{
		deviceName + "_task.png",
		220, 360, 80, 40,
	}
}
func TiLiImg(deviceName string) *Img {
	return &Img{
		deviceName + "_tili.png",
		150, 10, 40, 30,
	}
}
func TiLi1Img(deviceName string) *Img {
	return &Img{
		deviceName + "_tili1.png",
		450, 150, 100, 30,
	}
}

func TimeImg(deviceName string) *Img {
	return &Img{
		deviceName + "_time.png",
		540, 1177, 100, 25,
	}
}

func GuaiWuImg(deviceName string) *Img {
	return &Img{
		deviceName + "_guaiwu.png",
		440, 1150, 260, 30,
	}
}

func GjImg(deviceName string) *Img {
	return &Img{
		deviceName + "_gongji.png",
		180, 840, 360, 40,
	}
}

func ZhanliImg(deviceName string) *Img {
	return &Img{
		deviceName + "_zhanli.png",
		450, 100, 150, 30,
	}
}
func BuDuiImg(deviceName string) *Img {
	return &Img{
		deviceName + "_budui.png",
		290, 20, 140, 30,
	}
}

func MasterImg(deviceName string) *Img {
	return &Img{
		deviceName + "_pic_0.png",
		20, 20, 80, 80,
	}
}

func CloseImg(deviceName string) *Img {
	return &Img{
		deviceName + "_close.png",
		620, 240, 60, 60,
	}
}

func ZiYuanBaoImg(deviceName string) *Img {
	return &Img{
		deviceName + "_ziyuanbao.png",
		240, 910, 240, 50,
	}
}
func ChengBaoImg(deviceName string) *Img {
	return &Img{
		deviceName + "_chengbao.png",
		270, 20, 160, 40,
	}
}
