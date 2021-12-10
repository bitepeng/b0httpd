package main

import (
	"flag"

	"easyhttpd/core/httpd"
)

var (
	RootPath   string //根目录
	ServerPort string //服务端口
)

func main() {
	//命令行参数
	flag.StringVar(&RootPath, "r", "./", "-r for HttpServer root path")
	flag.StringVar(&ServerPort, "p", ":8008", "-p for HttpServer server port")
	flag.Parse()

	//http服务器
	httpd := httpd.Server{}
	httpd.Start(RootPath, ServerPort)
}
