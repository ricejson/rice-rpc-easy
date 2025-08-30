package server

import (
	"github.com/ricejson/rice-rpc-easy/serializer"
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
		nativeSerializer := serializer.NewNativeSerializer()
		handler := NewRequestHandler(writer, request, nativeSerializer)
		// 记录日志
		handler.doLog()
		// 解析rpc请求
		rpcRequest := handler.parseRpcRequest()
		// 反射调用实现类方法
		rpcResponse := handler.invoke(rpcRequest)
		// 写响应
		handler.doResponse(rpcResponse)
	})

	// 启动服务器并监听端口
	var portStr = strconv.Itoa(port)
	if err := http.ListenAndServe(":"+portStr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	} else {
		log.Println("Server is now listening on port: " + portStr)
	}
}
