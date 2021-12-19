#  EasyHttpd
- EasyHttpd 超级简便的自建Http服务器
- 基于Go语言，极少依赖，一个文件，快速开启


##  下载快速开始
- 从bin目录下载对应操作系统的可执行文件
- windows平台，easyhttpd.exe 双击执行


## 命令行参数 
```
easyhttpd.exe -r ./ -p :8888
```
- 服务根文件目录 
  - `-r 根路径`
- 服务端口(可指定IP) 
  - `-p 0.0.0.0:8888` 
  - `-p :8888`

## 使用Go语言功能点
- [x]  net/http server
- [x]  go embed fs
- [x]  http template
- [x]  http file upload
- [x]  run cmd open
- [x]  file read and write
- [x]  host ip address
- [x]  multi platform cross compilation

## Docker镜像

### 构建容器镜像

`docker image build -t <namespace>/easyhttpd:<tag> .`

### 使用容器运行

`docker container run -t -v <local_html_path>:/srv/www -p 8008:8008 <namespace>/easyhttpd:<tag>`