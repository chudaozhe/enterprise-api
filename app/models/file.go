package models

import (
	"encoding/base64"
	"enterprise-api/app/config"
	"enterprise-api/core"
	"enterprise-api/core/helper"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
)

// 保存管理员头像
func SaveAdminAvatar(id int, content string) (dst string) {
	root, err0 := os.Getwd()
	if err0 != nil {
		return ""
	}
	conf := config.GetConfig().FileConfig
	dst = conf.Prefix + conf.Avatar
	dst = strings.ReplaceAll(dst, "[admin_id]", core.ToStr(id))
	dst = strings.ReplaceAll(dst, "[ext]", "jpg")

	_, err := os.Stat(root + "/" + path.Dir(dst))
	if os.IsNotExist(err) {
		os.MkdirAll(root+"/"+path.Dir(dst), 0777)
	}
	logrus.Info(root + "/" + dst)
	fileObj, err := os.OpenFile(root+"/"+dst, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "%s", Base64Decode(content))
	return
}

// 保存图片空间图片
func SavePhoto(content string) (dst string) {
	root, err0 := os.Getwd()
	if err0 != nil {
		return ""
	}
	fileUnixName := strconv.FormatInt(helper.GetUnix(), 10)
	conf := config.GetConfig().FileConfig
	dst = conf.Prefix + conf.Photo
	dst = strings.ReplaceAll(dst, "[date]", helper.GetDay()+"/"+fileUnixName)
	dst = strings.ReplaceAll(dst, "[ext]", "jpg")

	_, err := os.Stat(root + "/" + path.Dir(dst))
	if os.IsNotExist(err) {
		os.MkdirAll(root+"/"+path.Dir(dst), 0777)
	}
	logrus.Info(root + "/" + dst)
	fileObj, err := os.OpenFile(root+"/"+dst, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "%s", Base64Decode(content))
	return
}

// 保存编辑器图片
func SaveEditorImage(c *gin.Context, file *multipart.FileHeader) (dst string, err error) {
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		err = errors.New("文件类型不合法")
		return
	}

	root, err0 := os.Getwd()
	if err0 != nil {
		err = errors.New("根目录获取失败")
		return
	}
	fileUnixName := strconv.FormatInt(helper.GetUnix(), 10)
	conf := config.GetConfig().FileConfig
	dst = conf.Prefix + conf.Editor
	dst = strings.ReplaceAll(dst, "[date]", helper.GetDay()+"/"+fileUnixName)
	dst = strings.ReplaceAll(dst, "[ext]", extName)

	_, err1 := os.Stat(root + "/" + path.Dir(dst))
	if os.IsNotExist(err1) {
		os.MkdirAll(root+"/"+path.Dir(dst), 0777)
	}
	logrus.Info(root + "/" + dst)

	err2 := c.SaveUploadedFile(file, root+"/"+dst)
	if err2 != nil {
		err = errors.New("上传失败")
		return
	}
	return
}

// 删除文件
func DeleteFile(filePath string) bool { //需要测试
	return false
	err := os.Remove(filePath)
	if err != nil {
		fmt.Printf("%s", err)
		return false
	} else {
		return true
	}
}

func Base64Decode(str string) string {
	reader := strings.NewReader(str)
	decoder := base64.NewDecoder(base64.RawStdEncoding, reader)
	buf := make([]byte, 1024)
	dst := ""
	for {
		n, err := decoder.Read(buf)
		dst += string(buf[:n])
		if n == 0 || err != nil {
			break
		}
	}
	return dst
}
