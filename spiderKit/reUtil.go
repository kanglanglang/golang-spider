package spiderKit

import "regexp"

var (
	// 网页图片
	reImgTagStr = `<img.+?src="(http.+?)".*?>`
	// Img标签alt属性
	reTagAltStr = `alt="([\s\S]+?)"`
	// 原本图片的名称
	reImgNameStr = `/(\w+\.((jpg)|(jpeg)|(png)|(gif)|(bmp)|(webp)|(swf)|(ico)))`

	// 预编译正则对象
	reImgTag, reTagAlt, reImgName *regexp.Regexp
)

func init() {
	reImgTag = regexp.MustCompile(reImgTagStr)
	reTagAlt = regexp.MustCompile(reTagAltStr)
	reImgName = regexp.MustCompile(reImgNameStr)
}
