package server

import (
	"log"
	"net/http"
	"strconv"
)

// WebServer 自定义web服务器对象
type WebServer struct {
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

// DoStart 启动服务器
func (s *WebServer) DoStart(port int) {
	// 设置请求处理器
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 处理请求
		log.Println("Received request: " + request.Method + " " + request.RequestURI)
		// 发送响应
		writer.Header().Set("Content-Type", "text/plain")
		writer.Write([]byte("Hello from rice HTTP server"))
	})
	// 启动服务器并监听端口
	var portStr = strconv.Itoa(port)
	if err := http.ListenAndServe(":"+portStr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	} else {
		log.Println("Server is now listening on port: " + portStr)
	}
}
