package helper

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strings"
)

func VerifyEmail(email string) bool {
	emailRegex := regexp.MustCompile("^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$")
	return emailRegex.MatchString(email)
}

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[strings.ToLower(t.Field(i).Name)] = v.Field(i).Interface()
	}
	return data
}

func Md5(str string) string {
	data := []byte(str)
	return fmt.Sprintf("%x", md5.Sum(data))
}

// 支持大文件
func Md5File(file string) string {
	f, _ := os.Open(file)
	defer f.Close()
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		panic(err.Error())
	}
	return fmt.Sprintf("%x", md5hash.Sum(nil))
}
