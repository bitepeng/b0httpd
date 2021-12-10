package tools

import (
	"embed"
	"errors"
	"io"
	"io/fs"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strings"
)

// EmbedUI 嵌入静态资源
type EmbedUI struct {
	Fs   embed.FS // 静态资源
	Path string   // 设置embed文件到静态资源的相对路径，也就是embed注释里的路径
}

// EmbedUI.Open 静态资源访问
func (w *EmbedUI) Open(name string) (http.File, error) {
	log.Println("EmbedUI open: ", name)
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) {
		return nil, errors.New("http: invalid character in file path")
	}
	fullName := filepath.Join(w.Path, filepath.FromSlash(path.Clean("/"+name)))
	file, err := w.Fs.Open(fullName)
	wf := &EmbedUIFile{
		File: file,
	}
	log.Println("EmbedUI open err: ", err.Error())
	return wf, err
}

// EmbedUIFile 静态资源文件
type EmbedUIFile struct {
	io.Seeker
	fs.File
}

// EmbedUIFile.Readdir 静态资源目录
func (*EmbedUIFile) Readdir(count int) ([]fs.FileInfo, error) {
	return nil, nil
}
