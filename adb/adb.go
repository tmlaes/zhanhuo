package adb

import (
	"io/ioutil"
	"os/exec"
)

// AdbShellInputClick 屏幕点击
func adbShellInputClick(x, y string, device string) {
	exec.Command("adb", "-s", device, "shell", "input", "tap", x, y).Run()
}

// AdbShellInputSwipe 屏幕滑动
//在屏幕上做划屏操作，前四个数为坐标点，后面是滑动的时间（单位毫秒）
func adbShellInputSwipe(x, y, x1, y1 string, device string) {
	exec.Command("adb", "-s", device, "shell", "input", "swipe", x, y, x1, y1).Run()
}

// AdbShellScreenShot 截屏
func adbShellScreenShot(imgName, device string) {
	exec.Command("adb", "-s", device, "shell", "/system/bin/screencap", "-p", "/sdcard/Pictures/"+imgName).Run()
}

//AdbShellInputText 输入“字符”
func adbShellInputText(s, device string) {
	exec.Command("adb", "-s", device, "shell", "input", "text", s).Run()
}

//AdbShellInputKey 输入按键
func adbShellInputKey(key, device string) {
	exec.Command("adb", "-s", device, "shell", "input", "keyevent", key).Run()
}

// AdbShellDevices 获取设备列表
func adbShellDevices() string {
	MyCmd := exec.Command("adb", "devices")
	MyOut, _ := MyCmd.StdoutPipe()
	MyCmd.Start()
	MyBytes, _ := ioutil.ReadAll(MyOut)
	MyCmd.Wait()
	MyOut.Close()
	return string(MyBytes)
}
