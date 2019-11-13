package spiderKit

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	DialTimeout = 10 * time.Second // 拨号超时连接
	RWTimeout   = 10 * time.Second // 读写超时连接
)

var (
	httpClient http.Client
)

func init() {
	httpClient = http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				// 设置连接请求超时时间
				Conn, err := net.DialTimeout(netw, addr, DialTimeout)
				if err != nil {
					return nil, err
				}
				// 设置连接的读写超时时间
				deadline := time.Now().Add(RWTimeout)
				Conn.SetDeadline(deadline)
				return Conn, nil
			},
		},
	}
}

// DownloadFileWithClient 下载图片
func DownloadFileWithClient(url, fileName string) {
	fmt.Println("DownloadFileWithClient...")
	resp, err := httpClient.Get(url)
	if err != nil {
		fmt.Println(fileName, "下载失败")
		return
	}
	defer resp.Body.Close()

	imgBytes, _ := ioutil.ReadAll(resp.Body)
	err = ioutil.WriteFile(fileName, imgBytes, 0644)
	if err != nil {
		fmt.Println(fileName, "下载失败")
	}
	fmt.Println(fileName, "下载成功")
}

// GetHtml 获取网页源码
func GetHtml(url string) string {
	resp, err := httpClient.Get(url)
	HandleError(err,`httpClient.Get(url)`)
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	html := string(bytes)
	return html
}
