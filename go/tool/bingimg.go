package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"syscall"
	"time"
	"unsafe"
)

const (
	UserAgent      = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36 Edg/86.0.622.51"
	CurrentPathDir = "./bingimg/"
)

func init() {
	_ = os.Mkdir(CurrentPathDir, 0755)
}

// PathExists 判断文件是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// SetWindowsWallpaper 设置windows壁纸
func SetWindowsWallpaper(imagePath string) error {
	dll := syscall.NewLazyDLL("user32.dll")
	proc := dll.NewProc("SystemParametersInfoW")
	_t, _ := syscall.UTF16PtrFromString(imagePath)
	ret, _, _ := proc.Call(20, 1, uintptr(unsafe.Pointer(_t)), 0x1|0x2)
	if ret != 1 {
		return errors.New("System invoke failed.")
	}
	return nil
}

// GetBingBackgroundImageURL 获取bing主页的背景图片链接
func GetBingBackgroundImageURL() (string, error) {
	request, err := http.NewRequest("GET", "https://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1", nil)
	if err != nil {
		return "", err
	}
	request.Header.Set("user-agent", UserAgent)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	type BingImage struct {
		URL string `json:"url"`
	}
	type BingImages struct {
		Images []BingImage `json:"images"`
	}
	var bing = new(BingImages)
	err = json.Unmarshal(data, bing)
	if err != nil {
		return "", err
	}
	if len(bing.Images) > 0 && bing.Images[0].URL != "" {
		return "https://www.bing.com" + bing.Images[0].URL, nil
	}
	return "", errors.New(string(data))
}

// DownloadImage 下载图片,图片已存在则结束程序,保存并返回保存的文件名的绝对路径
func DownloadImage(imageURL string) (string, error) {
	day := time.Now().Format("2006-01-02")
	path := CurrentPathDir + day + ".jpg"
	if PathExists(path) {
		println(path + " is Exists.")
		os.Exit(1)
	}
	request, err := http.NewRequest("GET", imageURL, nil)
	if err != nil {
		return "", err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	f, err := os.Create(path)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(f, response.Body)
	if err != nil {
		return "", err
	}
	f.Close()
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return absPath, nil
}

// GetAndSet get bing img and set wallpaper
func GetAndSet() bool {
	println("get bing img url...")
	imageURL, err := GetBingBackgroundImageURL()
	if err != nil {
		println("get bing img url failed: " + err.Error())
		return false
	}
	println("get success: " + imageURL)

	println("download bing img...")
	imagePath, err := DownloadImage(imageURL)
	if err != nil {
		println("download bing img failed: " + err.Error())
		return false
	}
	println("set wallpaper...", imagePath)
	err = SetWindowsWallpaper(imagePath)
	if err != nil {
		println("set wallpaper failed: " + err.Error())
		return false
	}
	println("set wallpaper success.")
	return true
}

func main() {
	for i := 0; i < 10; i++ {
		if GetAndSet() {
			os.Exit(0)
		}
		time.Sleep(time.Second * time.Duration(i+1))
	}
}
