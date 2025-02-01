// @Author Gopher
// @Date 2025/2/1 20:56:00
// @Desc
package meta

// FileMeta 文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

// init 初始化创建
func init() {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMeta 新增或者更新文件的元信息
func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

// GetFileMeta 通过 sha1 获取文件的元信息的对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

// RemoveFileMeta 删除文件的元信息
func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas, fileSha1)
}
