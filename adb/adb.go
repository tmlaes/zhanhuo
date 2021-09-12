package adb

import (
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"os/exec"
	"strings"
)

// AdbShellInputClick 屏幕点击
func adbShellInputClick(x, y string, device string) {
	cmd := exec.Command("adb", "-s", device, "shell", "input", "tap", x, y)
	cmd.Start()
	cmd.Wait()
}

// AdbShellInputSwipe 屏幕滑动
//在屏幕上做划屏操作，前四个数为坐标点，后面是滑动的时间（单位毫秒）
func adbShellInputSwipe(x, y, x1, y1, time string, device string) {
	cmd := exec.Command("adb", "-s", device, "shell", "input", "swipe", x, y, x1, y1, time)
	cmd.Start()
	cmd.Wait()
}

// AdbShellScreenShot 截屏
func adbShellScreenShot(imgName, device string) {
	cmd := exec.Command("adb", "-s", device, "shell", "/system/bin/screencap", "-p", "/sdcard/Pictures/"+imgName)
	cmd.Start()
	cmd.Wait()
}

//AdbShellInputText 输入“字符”
func adbShellInputText(s, device string) {
	cmd := exec.Command("adb", "-s", device, "shell", "input", "text", s)
	cmd.Start()
	cmd.Wait()
}

//AdbShellInputKey 输入按键
func adbShellInputKey(key, device string) {
	cmd := exec.Command("adb", "-s", device, "shell", "input", "keyevent", key)
	cmd.Start()
	cmd.Wait()
}

// AdbShellDevices 获取设备列表
func adbShellDevices() []string {
	MyCmd := exec.Command("adb", "devices")
	MyOut, _ := MyCmd.StdoutPipe()
	MyCmd.Start()
	MyBytes, _ := ioutil.ReadAll(MyOut)
	MyCmd.Wait()
	MyOut.Close()
	return parseDevices(string(MyBytes))
}

func ListLy() []string {
	var enc mahonia.Decoder
	enc = mahonia.NewDecoder("gbk")
	cmd := exec.Command("D:\\BaiZhi\\LSPlayer64\\lsconsole.exe", "list2")
	out, _ := cmd.StdoutPipe()
	cmd.Start()
	reader := enc.NewReader(out)
	bytes, _ := ioutil.ReadAll(reader)
	cmd.Wait()
	out.Close()
	s := string(bytes)
	//索引，标题，顶层窗口句柄，绑定窗口句柄，是否进入android，进程PID，VBox进程PID
	//fmt.Println(s)
	return parseDevices1(s)
}

//action <--name mnq_name | --index mnq_idx> --key --value
func Shake(index string) {
	//dnconsole.exe action --name *** --key call.shake --value null
	cmd := exec.Command("D:\\BaiZhi\\LSPlayer64\\lsconsole.exe", "action", "--index", index, "--key", "call.shake", "--value", "null")
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	cmd.Wait()
}

func Launch(index string) {
	//lsconsole.exe launch --name 雷神模拟器
	//lsconsole.exe launch --index 0
	cmd := exec.Command("D:\\BaiZhi\\LSPlayer64\\lsconsole.exe", "launch", "--index", index)
	cmd.Start()
	cmd.Wait()
}

func Quit(index string) {
	//lsconsole.exe launch --name 雷神模拟器
	//lsconsole.exe launch --index 0
	cmd := exec.Command("D:\\BaiZhi\\LSPlayer64\\lsconsole.exe", "quit", "--index", index)
	cmd.Start()
	cmd.Wait()
}

func Quitall() {
	cmd := exec.Command("D:\\BaiZhi\\LSPlayer64\\lsconsole.exe", "quitall")
	cmd.Start()
	cmd.Wait()
}

func parseDevices(s string) []string {
	devices := make([]string, 0)
	if s == "" {
		return devices
	}
	split := strings.Split(s, "\r\n")
	for i, s2 := range split {
		if i == 0 || s2 == "" || strings.Index(s2, "emulator") < 0 {
			continue
		}
		s3 := strings.Split(s2, "\t")
		devices = append(devices, s3[0])
	}
	return devices
}

func parseDevices1(s string) []string {
	devices := make([]string, 0)
	if s == "" {
		return devices
	}
	split := strings.Split(s, "\r\n")
	for _, s2 := range split {
		if s2 == "" {
			continue
		}
		s3 := strings.Split(s2, ",")
		if s3[4] == "1" {
			devices = append(devices, s3[0])
		}
	}
	return devices
}
