package process

import (
	"errors"
	"fmt"
	"github.com/corona10/goimagehash"
	"image"
	"image/png"
	"os"
	"os/exec"
	"strings"
	"time"
	"zhanhuo/adb"
	P "zhanhuo/entity"
)

const (
	BasePath     = "C:\\Users\\Administrator\\Desktop\\temp\\"
	CloseImg     = "images/close.png"
	MasterImg    = "images/pic_0.png"
	RealImg      = "images/pic_1.png"
	JumpImg      = "images/jump.png"
	ActImg       = "images/act.png"
	GojImg       = "images/gongji.png"
	BuDuiImg     = "images/budui.png"
	ZiYuanBaoImg = "images/ziyuanbao.png"
	ChengBaoImg  = "images/chengbao.png"
	UpdateImg    = "images/update.png"
)

func CheckStart(deviceName string) bool {
	//截图
	shotImg := adb.Screenshot(deviceName)
	//关闭窗口
	checkClose(shotImg, deviceName)
	//校验是否进入主界面
	return checkComplete(shotImg, deviceName)
}

func checkComplete(shotImg image.Image, deviceName string) bool {
	//裁剪领主头像
	img := P.MasterImg(deviceName)
	CutImage(shotImg, img)
	return Compare(img.Name, MasterImg)
}

//校验是否有需要关闭窗口
func checkClose(shotImg image.Image, deviceName string) {
	img := P.CloseImg(deviceName)
	CutImage(shotImg, img)
	flag := Compare(img.Name, CloseImg)
	if flag {
		fmt.Println(time.Now(), "关闭弹窗")
		adb.Click(P.Close.X, P.Close.Y, 1, deviceName)
	}
}

func GetText(shotImg image.Image, cutImg *P.Img) string {
	CutImage(shotImg, cutImg)
	return ocr(cutImg.Name)
}

func GetText1(device string, cutImg *P.Img) string {
	shotImg := adb.Screenshot(device)
	CutImage(shotImg, cutImg)
	return ocr(cutImg.Name)
}

func CutImage(shotImg image.Image, cutImg *P.Img) error {
	var subImg image.Image
	var x = cutImg.X
	var y = cutImg.Y
	var w = cutImg.W
	var h = cutImg.H
	if rgbImg, ok := shotImg.(*image.YCbCr); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.YCbCr) //图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := shotImg.(*image.RGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.RGBA) //图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := shotImg.(*image.NRGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.NRGBA) //图片裁剪x0 y0 x1 y1
	} else {
		return errors.New("图片解码失败")
	}

	f, err := os.OpenFile(BasePath+cutImg.Name, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer f.Close()
	return png.Encode(f, subImg)
}

func ocr(imgPath string) string {
	body, err := exec.Command("tesseract", BasePath+imgPath, "stdout", "-l", "normal+eng+chi_sim").Output()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	//os.Remove(BasePath+imgPath)
	s := string(body)
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\r\n", "")
	return s
}

func Compare(imageName string, oldImg string) bool {
	queryFile, _ := os.Open(oldImg)
	file1, _ := os.Open(BasePath + imageName)
	//defer os.Remove(BasePath+imageName)
	defer queryFile.Close()
	defer file1.Close()
	imgQuery, _ := png.Decode(queryFile)
	img1, _ := png.Decode(file1)
	queryHash, _ := goimagehash.DifferenceHash(imgQuery)
	hash1, _ := goimagehash.DifferenceHash(img1)
	distance, _ := queryHash.Distance(hash1)
	if distance == 0 {
		return true
	}
	return false
}
