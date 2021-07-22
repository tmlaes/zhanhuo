package utils

import "strings"

func GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}

func GetStr(str, start string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	}
	n = n + len(start)
	str = string([]byte(str)[n:])
	return str
}
