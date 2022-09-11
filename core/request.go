package core

import "strconv"

func ToInt(str string) (num int) {
	num, _ = strconv.Atoi(str)
	return
}

func ToStr(int1 int) (str string) {
	str = strconv.Itoa(int1)
	return
}
