package spiderKit

import (
	"fmt"
	"strings"
)

// GetPageImgInfosToChan 获得页面上的全部文件信息
func GetPageImgInfosToChan(url,imgDir string, chImgMaps chan<- map[string]string) {
	imgInfos := GetPageImgInfos(url,imgDir)
	for _, imgInfoMap := range imgInfos {
		chImgMaps <- imgInfoMap
	}
}

// GetPageImgUrls 获得页面上的所有图片链接
func GetPageImgInfos(url, imgDir string) []map[string]string {
	html := GetHtml(url)
	// 出现乱码需要转码
	//bytes := ConvertToByte(html, "gbk", "utf8")
	//html = string(bytes)

	//re := regexp.MustCompile(reImgTagStr)
	rets := reImgTag.FindAllStringSubmatch(html, -1)
	fmt.Println("捕获图片张数：", len(rets))

	// 返回切片
	imgInfos := make([]map[string]string, 0)
	for _, ret := range rets {
		imgInfo := make(map[string]string)
		imgInfo["url"] = ret[1]
		imgInfo["filename"] = GetImgNameFromTag(ret[0], ret[1], imgDir)
		imgInfos = append(imgInfos, imgInfo)
	}
	return imgInfos
}

// GetImgNameFromTag 从img标签中获取文件名
// 有alt使用alt做文件名，没有则使用时间戳_随机数做文件名
// dir 文件路径
// suffix 后缀
func GetImgNameFromTag(imgTag, imgUrl, imgDir string) (fileName string) {
	suffix := ".jpg"

	// 先获取图片格式
	imgName := GetImgNameFromImgUrl(imgUrl)
	if imgName != "" {
		suffix = imgName[strings.LastIndex(imgName, "."):]
	}

	// 尝试从imgTag中提取alt
	//re := regexp.MustCompile(reTagAltStr)
	rets := reTagAlt.FindAllStringSubmatch(imgTag, 1)

	if len(rets) > 0 {
		// 文件名字首选alt
		alt := rets[0][1]
		alt = strings.Replace(alt, ":", "_", -1)
		fileName = alt + suffix
	} else if imgName != "" {
		// 文件名字次选链接中的文件名
		fileName = imgName
	} else {
		// 文件名字末选时间戳+_随机数
		fileName = GetRandomFileName() + suffix
	}
	fileName = imgDir + fileName
	return fileName
}

// GetImgNameFromImgUrl 获取图片的本身的自己的名字
func GetImgNameFromImgUrl(imgUrl string) string {
	//re := regexp.MustCompile(reImgNameStr)
	rets := reImgName.FindAllStringSubmatch(imgUrl, -1)
	if len(rets) > 0 {
		return rets[0][1]
	} else {
		return ""
	}
}
