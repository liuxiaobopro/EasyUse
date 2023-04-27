package logic

import (
	"mime/multipart"
	"path"

	"github.com/gin-gonic/gin"
	respx "github.com/liuxiaobopro/gobox/resp"
	"github.com/zeromicro/go-zero/core/logx"
)

// UploadBook 上传书籍
func (*logic) UploadBook(c *gin.Context, f *multipart.FileHeader) (interface{}, respx.Pt) {
	// // 打开文件句柄
	// file, err := f.Open()
	// if err != nil {
	// 	logx.Error(err)
	// 	return nil, respx.InternalErrT
	// }
	// defer file.Close()

	// // 创建目标文件
	// dst, err := os.Create("/" + f.Filename)
	// if err != nil {
	// 	logx.Error(err)
	// 	return nil, respx.InternalErrT
	// }
	// defer dst.Close()

	// // 复制文件内容到目标文件
	// if _, err = io.Copy(dst, file); err != nil {
	// 	logx.Error(err)
	// 	return nil, respx.InternalErrT
	// }

	dst := path.Join("frontend/public/book", f.Filename)
	if err := c.SaveUploadedFile(f, dst); err != nil {
		logx.Error(err)
		return nil, respx.InternalErrT
	}

	return nil, nil
}
