package util

import (
	"encoding/base64"
	"io"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"crypto/md5"
	"os"
	"path"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

// 创建GUID
func CreateGUID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	return base64.URLEncoding.EncodeToString(b)
}

// 创建随机串
func RandString(n int) string {
	guid := CreateGUID()
	return SubString(guid, 0, n)
}

// 截取子串
func SubString(str string, begin, length int) (substr string) {
	rs := []rune(str)
	lth := len(rs)

	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	return string(rs[begin:end])
}

//返回字符串的MD5值
func Md5(value string) string {
	h := md5.New()
	h.Write([]byte(value))
	return fmt.Sprintf("%s", hex.EncodeToString(h.Sum(nil)))
}

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

//创建文件
func CreateFile(filePath string) {
	os.MkdirAll(path.Dir(filePath), os.ModePerm)
	os.Create(filePath)
}

// 读取文本文件
func ReadFile(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)

	return string(fd)
}

//byte转map
func Byte2Map(data []byte) (dataMap map[string]interface{}, err error) {
	err = json.Unmarshal(data, &dataMap)
	return
}


//string转float
func Str2Float(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}