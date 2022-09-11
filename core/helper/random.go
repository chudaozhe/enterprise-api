package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func Random(length int, char string) string {
	if len(char) == 0 {
		char = "1234567890asdfghklzxcvbnmqwertyuiop"
	}
	rand.Seed(time.Now().UnixNano())
	str := ""
	for i := 0; i < length; i++ {
		str += fmt.Sprintf("%c", char[rand.Intn(len(char))])
	}
	return str
}
