// @Author Gopher
// @Date 2025/2/1 19:55:00
// @Desc
package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// UploadHandler 处理文件上传
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// 	check the method of GET
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
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data, err : %s\n", err.Error())
			return
		}
		defer file.Close()

		// 在指定目录进行创建
		// notice : 这个地方一定要加上 "/" 符号，不然是到达不了这个目录里面去的
		newFile, err := os.Create("/Users/spring/workspace/code/src/project/dev/filestore-server/docs/" + head.Filename)
		if err != nil {
			fmt.Printf("Failed to create file, err : %s\n", err.Error())
			return
		}
		defer newFile.Close()

		// 保存指定文件
		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to save data into file , err : %s\n", err.Error())
			return
		}

		// 如果没有发生错误，重定向到保存成功的页面
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

// UploadSucHandler 上传已完成
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload finished !")
}
