package adb

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os/exec"
	"strconv"
	"time"
	"zhanhuo/entity"
	"zhanhuo/utils"
)

func Devices() []string {
	for {
		devices := adbShellDevices()
		if len(devices) > 0 {
			return devices
		}
		delay(2)
	}
}

func Devices_(arr []string) map[string]string {
	deviceMap := make(map[string]string)
	for _, s := range arr {
		fmt.Println("启动序号：", s)
		Launch(s)
		for {
			ly := ListLy()
			devices := adbShellDevices()
			l := len(deviceMap)
			if len(ly) > l && len(devices) > l {
				key := devices[l]
				value := ly[l]
				deviceMap[key] = value
				break
			}
			<-time.After(time.Duration(1) * time.Second)
		}
	}
	return deviceMap
}

func ClickPoint(point entity.Point, times int, device string) {
	if point.Check {
		fmt.Println(utils.Now(), device, "校验弹窗")
		ClosePop(device)
	}
	fmt.Println(utils.Now(), device, point.Desc)
	adbShellInputClick(strconv.Itoa(point.X), strconv.Itoa(point.Y), device)
	if times > 0 {
		delay(times)
	}
}

func ClickPoint_(point entity.Point, times int, device string) {
	fmt.Println(utils.Now(), device, point.Desc)
	adbShellInputClick(strconv.Itoa(point.X_), strconv.Itoa(point.Y_), device)
	delay(times)
}

func ClickPointOffset(point entity.Point, x, y int, times int, device string) {
	if point.Check {
		fmt.Println(utils.Now(), device, "校验弹窗")
		ClosePop(device)
	}
	fmt.Println(utils.Now(), device, point.Desc)
	adbShellInputClick(strconv.Itoa(point.X+x), strconv.Itoa(point.Y+y), device)
	delay(times)
}

func ClickPointOffset_(point entity.Point, x, y int, times int, device string) {
	adbShellInputClick(strconv.Itoa(point.X_+x), strconv.Itoa(point.Y_+y), device)
	delay(times)
}

func ClickMore(point entity.Point, n, times int, device string) {
	if point.Check {
		fmt.Println(utils.Now(), device, "校验弹窗")
		ClosePop(device)
	}
	fmt.Println(utils.Now(), device, point.Desc)
	for i := 0; i < n; i++ {
		adbShellInputClick(strconv.Itoa(point.X), strconv.Itoa(point.Y), device)
		fmt.Println("第", i+1, "次点击")
		delay(times)
	}
}

func Swipe(x, y, x1, y1 int, times int, device string) {
	adbShellInputSwipe(strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(x1), strconv.Itoa(y1), strconv.Itoa(times), device)
}

func ScreenShot(imgName, device string) {
	adbShellScreenShot(imgName, device)
}

func Screenshot(device string) (img image.Image) {
	cmd := exec.Command("adb", "-s", device, "shell", "screencap", "-p")
	out, _ := cmd.StdoutPipe()
	cmd.Start()
	outBytes, err := ioutil.ReadAll(out)
	if err != nil {
		fmt.Println("截屏失败", err.Error())
		return nil
	}
	x := bytes.Replace(outBytes, []byte("\r\n"), []byte("\n"), -1)
	img, _ = png.Decode(bytes.NewReader(x))
	cmd.Wait()
	out.Close()
	return
}

func Text(s, device string) {
	adbShellInputText(s, device)
	delay(1)
}

func Key(key, device string) {
	adbShellInputKey(key, device)
	delay(1)
}

func ScreenZero(device string) {
	adbShellInputSwipe("0", "0", "1000", "1000", "500", device)
	delay(1)
}

func ScreenOne(device string) {
	ScreenZero(device)
	adbShellInputSwipe("200", "500", "150", "150", "3000", device)
	delay(2)
}
func ScreenTwo(device string) {
	ScreenZero(device)
	adbShellInputSwipe("690", "1110", "120", "450", "5000", device)
	delay(2)
}
func ScreenThree(device string) {
	ScreenZero(device)
	adbShellInputSwipe("950", "100", "150", "100", "6000", device)
	adbShellInputSwipe("200", "600", "200", "200", "5000", device)
	delay(2)
}
func ScreenFour(device string) {
	ScreenZero(device)
	adbShellInputSwipe("1100", "100", "0", "100", "6000", device)
	adbShellInputSwipe("200", "600", "200", "200", "3000", device)
	delay(2)
}

func ScreenFive(device string) {
	ScreenZero(device)
	adbShellInputSwipe("0", "1000", "0", "100", "5000", device)
	adbShellInputSwipe("1000", "1000", "0", "1000", "5000", device)
	delay(2)
}

func ScreenSix(device string) {
	ScreenZero(device)
	adbShellInputSwipe("0", "1000", "0", "100", "6000", device)
	adbShellInputSwipe("500", "1000", "0", "1000", "3000", device)
	delay(2)
}

func ScreenSeven(device string) {
	adbShellInputSwipe("1000", "1000", "0", "0", "500", device)
	delay(1)
	adbShellInputSwipe("400", "1000", "500", "1000", "3000", device)
	delay(2)
}

func ScreenRollSkill(device string) {
	adbShellInputSwipe("500", "1240", "500", "220", "5000", device)
	delay(1)
}

func delay(times int) {
	<-time.After(time.Duration(times) * time.Second)
}
