package file

import (
	"io"
	"net/http"
	"os"
	"path"
)

// 下载网络io数据到文件
func Download(f, url string) error{
	fd, err := os.Create(f)
	if err != nil {
		return err
	}

	defer fd.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, err = io.Copy(fd, resp.Body)
	return err
}

// 获取文件目录
func Dir(fp string) string{
	return path.Dir(fp)
}

// 创建目录
func CreateDir(fp string) error{
	if IsExist(fp){
		return nil
	}
	return os.MkdirAll(fp, os.ModePerm)
}

// 创建文件,umask 0666如果存文件存在则中断
func Create(filename string) (*os.File, error){
	return os.Create(filename)
}

func OpenFile(filename string) (*os.File, error){
	return os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
}

// 删除文件
func RemoveFile(filename string) error{
	return os.Remove(filename)
}

// 判断文件或者目录是否存在
func IsExist(fp string) bool{
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}

func WriteBytes(filePath string, b []byte) (int, error){
	os.MkdirAll(path.Dir(filePath), os.ModePerm)
	fw, err := Create(filePath)
	if err != nil {
		return 0, err
	}
	defer fw.Close()
	return fw.Write(b)
}

func WriteString(filePath, s string) (int, error){
	return WriteBytes(filePath, []byte(s))
}
