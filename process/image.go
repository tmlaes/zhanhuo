package process

import (
	"errors"
	"fmt"
	"github.com/Comdex/imgo"
	"github.com/corona10/goimagehash"
	"image"
	"image/png"
	"os"
	"os/exec"
	"strings"
	"zhanhuo/adb"
	"zhanhuo/entity"
)

const (
	BasePath = "C:\\Users\\Administrator\\Desktop\\temp\\"
)

func CheckStart(deviceName string) bool {
	//关闭窗口
	closePop(deviceName)
	//校验是否进入主界面
	return checkComplete(deviceName)
}

func checkComplete(deviceName string) bool {
	//裁剪领主头像
	return Compare1(entity.MasterImg(deviceName), deviceName)
}

//校验是否有需要关闭窗口
func closePop(device string) {
	if Compare1(entity.Img8(device), device) {
		adb.ClickPoint(entity.P64, 2, device)
		adb.ClickPoint(entity.P55, 1, device)
	}
	if Compare1(entity.CloseImg(device), device) {
		adb.ClickPoint(entity.Close, 2, device)
	}
	checkBack(device)
}

func checkBack(device string) {
	for {
		if Compare1(entity.BackImg(device), device) {
			adb.ClickPoint(entity.Back, 2, device)
		} else {
			return
		}
	}

}

func GetText(device string, cutImg *entity.Img) string {
	shotImg := adb.Screenshot(device)
	CutImage(shotImg, cutImg)
	return ocr(cutImg.Name)
}

func GetText1(device string, cutImg *entity.OcrImg) string {
	shotImg := adb.Screenshot(device)
	CutImage1(shotImg, cutImg)
	transparentImg(cutImg)
	return ocr(cutImg.Name)
}

func Compare1(cutImg *entity.Img, device string) bool {
	shotImg := adb.Screenshot(device)
	CutImage(shotImg, cutImg)
	return Compare(cutImg)
}

func CutImage(shotImg image.Image, cutImg *entity.Img) error {
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

func CutImage1(shotImg image.Image, cutImg *entity.OcrImg) error {
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
	os.Remove(BasePath + imgPath)
	s := string(body)
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\r\n", "")
	return s
}

func Compare(cutImg *entity.Img) bool {
	queryFile, _ := os.Open(cutImg.Path)
	file1, _ := os.Open(BasePath + cutImg.Name)
	defer os.Remove(BasePath + cutImg.Name)
	defer queryFile.Close()
	defer file1.Close()
	imgQuery, _ := png.Decode(queryFile)
	img1, _ := png.Decode(file1)
	queryHash, _ := goimagehash.DifferenceHash(imgQuery)
	hash1, _ := goimagehash.DifferenceHash(img1)
	distance, _ := queryHash.Distance(hash1)
	if distance <= 5 {
		return true
	}
	return false
}

//把图片背景设置成透明
func transparentImg(img *entity.OcrImg) {
	imgArr := imgo.MustRead(BasePath + img.Name)
	for i := 0; i < img.H; i++ {
		for j := 0; j < img.W; j++ {
			if img.RGB[0][0] <= imgArr[i][j][0] && imgArr[i][j][0] <= img.RGB[0][1] &&
				img.RGB[1][0] <= imgArr[i][j][1] && imgArr[i][j][1] <= img.RGB[1][1] &&
				img.RGB[2][0] <= imgArr[i][j][2] && imgArr[i][j][2] <= img.RGB[2][1] {
				imgArr[i][j][0] = 0
				imgArr[i][j][1] = 0
				imgArr[i][j][2] = 0
				imgArr[i][j][3] = 255
			} /*else {
				imgArr[i][j][0]=255
				imgArr[i][j][1]=255
				imgArr[i][j][2]=255
				imgArr[i][j][3] = 255
			}*/
		}
	}
	err := imgo.SaveAsPNG(BasePath+img.Name, imgArr)
	if err != nil {
		panic(err)
	}
}
