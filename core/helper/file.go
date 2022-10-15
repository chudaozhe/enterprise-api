package helper

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/**
// 写入文件
_ = FilePutContents("1.txt", "abc", 0644)

// 读取文件
value, _ := FileGetContents("1.txt")
fmt.Print(value)
*/
// 读取整个文件
func FileGetContents(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	return string(data), err
}

// 把字符串写入文件
func FilePutContents(filename string, data string, mode os.FileMode) error {
	return os.WriteFile(filename, []byte(data), mode)
}

// 获取远程内容
func GetContent(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}
