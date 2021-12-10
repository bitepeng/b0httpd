package httpd

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// fileHandle 文件上传界面
func (s *Server) fileHandle(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"title": "Easyhttpd File Upload"}
	s.template.ExecuteTemplate(w, "upload.html", data)

}

// uploadHandle 上传文件操作
func (s *Server) uploadHandle(w http.ResponseWriter, r *http.Request) {
	file, head, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		io.WriteString(w, err.Error())
		return
	}
	defer file.Close()
	filePath := s.root + head.Filename
	fW, err := os.Create(filePath)
	if err != nil {
		io.WriteString(w, "文件创建失败")
		return
	}
	defer fW.Close()
	_, err = io.Copy(fW, file)
	if err != nil {
		io.WriteString(w, "文件保存失败")
		return
	}
	io.WriteString(w, "save to "+filePath)
}
