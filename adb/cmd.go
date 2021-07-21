package adb

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os/exec"
	"strconv"
	"strings"
	"time"
	P "zhanhuo/entity"
)

func Devices() []string {
	deviceStr := adbShellDevices()
	return parseDevices(deviceStr)
}

func Click(x, y int, times int, device string) {
	adbShellInputClick(strconv.Itoa(x), strconv.Itoa(y), device)
	time.Sleep(time.Duration(times) * time.Second)
}

func ClickMore(x, y int, n, times int, device string) {
	for i := 0; i < n; i++ {
		adbShellInputClick(strconv.Itoa(x), strconv.Itoa(y), device)
		time.Sleep(time.Duration(times) * time.Second)
	}
}

func Swipe(x, y, x1, y1 int, times int, device string) {
	adbShellInputSwipe(strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(x1), strconv.Itoa(y1), device)
	time.Sleep(time.Duration(times) * time.Second)
}

func ScreenShot(imgName, device string) {
	adbShellScreenShot(imgName, device)
}

func Screenshot(device string) (img image.Image) {
	cmd := exec.Command("adb", "-s", device, "shell", "screencap", "-p")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	x := bytes.Replace(out.Bytes(), []byte("\r\n"), []byte("\n"), -1)
	img, _ = png.Decode(bytes.NewReader(x))
	return
}

func Text(s, device string) {
	adbShellInputText(s, device)
	time.Sleep(time.Duration(1) * time.Second)
}

func Key(key, device string) {
	adbShellInputKey(key, device)
	time.Sleep(time.Duration(1) * time.Second)
}

func ScreenOne(device string) {
	resetPoint(device)
	Swipe(P.One.X, P.One.Y, P.One.X_, P.One.Y_, 2, device)
}
func ScreenTwo(device string) {
	resetPoint(device)
}
func ScreenThree(device string) {
	resetPoint(device)
	Swipe(P.Three.X, P.Three.Y, P.Three.X_, P.Three.Y_, 2, device)
}
func ScreenFour(device string) {
	resetPoint(device)
	Swipe(P.Four.X, P.Four.Y, P.Four.X_, P.Four.Y_, 2, device)
}

func ScreenFive(device string) {
	resetPoint(device)
	Swipe(P.Five.X, P.Five.Y, P.Five.X_, P.Five.Y_, 2, device)
}

func resetPoint(device string) {
	fmt.Println(time.Now(), "重置坐标", device)
	Click(P.Reset.X, P.Reset.Y, 5, device)
	Click(P.Reset.X, P.Reset.Y, 5, device)
}

func parseDevices(s string) []string {
	devices := make([]string, 0)
	split := strings.Split(s, "\r\n")
	for i, s2 := range split {
		if i == 0 || s2 == "" {
			continue
		}
		s3 := strings.Split(s2, "\t")
		devices = append(devices, s3[0])
	}
	return devices
}
