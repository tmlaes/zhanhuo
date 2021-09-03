package adb

import (
	"errors"
	"fmt"
	"github.com/corona10/goimagehash"
	"image"
	"image/png"
	"os"
	"zhanhuo/entity"
	"zhanhuo/utils"
)

const (
	BasePath = "C:\\Users\\Administrator\\Desktop\\temp\\"
)

func close1(device string) {
	shotImg := Screenshot(device)
	img1 := entity.CloseImg(device)
	cutImage2(shotImg, img1)
	if compare2(img1) {
		fmt.Println(utils.Now(), device, "关闭弹窗")
		ClickPoint(entity.Close, 2, device)
		ClickPoint(entity.P55, 2, device)
	}
	img2 := entity.BackImg(device)
	cutImage2(shotImg, img2)
	if compare2(img2) {
		fmt.Println(utils.Now(), device, "返回弹窗")
		ClickPoint(entity.Back, 2, device)
		ClickPoint(entity.P55, 2, device)
	}
}

func compare3(cutImg *entity.Img, device string) bool {
	shotImg := Screenshot(device)
	cutImage2(shotImg, cutImg)
	return compare2(cutImg)
}

func compare2(cutImg *entity.Img) bool {
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

func cutImage2(shotImg image.Image, cutImg *entity.Img) error {
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
