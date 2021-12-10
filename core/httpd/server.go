package httpd

import (
	"log"
	"net/http"
	"text/template"

	"easyhttpd/core/tools"
	"easyhttpd/resource"
)

// Server 配置Httpd服务
type Server struct {
	root     string
	port     string
	template *template.Template
}

// Start 启动Httpd服务器
func (s *Server) Start(root, port string) {

	//配置服务
	s.root = root
	s.port = port
	tpl, _ := template.ParseFS(resource.Templates, "templates/*.html")
	s.template = tpl

	//文件服务器
	mux := http.DefaultServeMux
	fs := http.FileServer(http.Dir(s.root))
	mux.Handle("/", http.StripPrefix("/", fs))

	//服务路由
	s.router(mux)

	//启动服务器
	log.Printf("EasyHttpd Server Started! (root=> %s,port=>%s)\n", s.root, s.port)
	tools.Open("http://127.0.0.1" + s.port)
	log.Fatal(http.ListenAndServe(s.port, nil))

}

// router 服务路由注册
func (s *Server) router(mux *http.ServeMux) {

	//上传文件
	mux.HandleFunc("/file", s.fileHandle)
	mux.HandleFunc("/upload", s.uploadHandle)

}
