// @Author Gopher
// @Date 2025/2/1 19:55:00
// @Desc
package main

import (
	"filestore-server/handler"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMeaHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server, err : %s", err.Error())
	}
}
