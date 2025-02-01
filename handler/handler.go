// @Author Gopher
// @Date 2025/2/1 19:55:00
// @Desc
package handler

import (
	"io"
	"io/ioutil"
	"net/http"
)

// UploadHandler 处理文件上传
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 返回上传的 html 的页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internal server error")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		// 接收文件流及存储到本地的目录

	}
}
