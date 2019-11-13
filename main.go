package main

import (
	"strconv"
	"studygo/spider2/spiderKit"
	"sync"
)

var (
	// 图片存储地址
	imgDir = `G:\Go\src\studygo\spider\imgs\`
	// 管道 放map存图片地址和名字
	chImgMaps = make(chan map[string]string, 100)
	// 写一个管道
	chSem = make(chan int, 10)
	// 图片信息
	wg4ImgInfo sync.WaitGroup
	// 下载
	wg4Download sync.WaitGroup
)

func main() {
	baseUrl := "https://www.duotoo.com/zt/rbmn/index"
	for i := 1; i < 25; i++ {
		var url string
		if i == 1 {
			url = baseUrl + ".html"
		} else {
			url = baseUrl + "_" + strconv.Itoa(i) + ".html"
		}

		// 这里也写一个携程 获取imageInfo
		wg4ImgInfo.Add(1)
		go func(theUrl string) {
			spiderKit.GetPageImgInfosToChan(theUrl, imgDir, chImgMaps)
			wg4ImgInfo.Done()
		}(url)
	}

	go func() {
		wg4ImgInfo.Wait()
		close(chImgMaps)
	}()

	// 从管道中读取下载
	for imgMap := range chImgMaps {
		wg4Download.Add(1)
		go func(im map[string]string) {
			chSem <- 123
			spiderKit.DownloadFileWithClient(im["url"], im["filename"])
			<-chSem
			wg4Download.Done()
		}(imgMap)
	}

	// 全爬取完再停
	wg4Download.Wait()
}
