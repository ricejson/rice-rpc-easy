package server

import (
	"errors"
	"github.com/ricejson/rice-rpc-easy/model"
	"github.com/ricejson/rice-rpc-easy/registry"
	"github.com/ricejson/rice-rpc-easy/serializer"
	"io"
	"log"
	"net/http"
	"reflect"
)

type RequestHandler struct {
	writer     http.ResponseWriter
	request    *http.Request
	serializer serializer.Serializer
}

func NewRequestHandler(writer http.ResponseWriter, request *http.Request, serializer serializer.Serializer) *RequestHandler {
	return &RequestHandler{
		writer:     writer,
		request:    request,
		serializer: serializer,
	}
}

// doLog 记录日志
func (h *RequestHandler) doLog() {
	log.Println("Received request: " + h.request.Method + " " + h.request.RequestURI)
}

// parseRpcRequest 解析rpc请求
func (h *RequestHandler) parseRpcRequest() *model.RpcRequest {
	body, _ := io.ReadAll(h.request.Body)
	log.Println("rpc request body: " + string(body))
	defer h.request.Body.Close()
	// 反序列化
	rpcRequest := model.RpcRequest{}
	_ = h.serializer.Deserialize(body, &rpcRequest)
	return &rpcRequest
}

// invoke 反射调用对应实现类方法
func (h *RequestHandler) invoke(request *model.RpcRequest) *model.RpcResponse {
	// 根据服务名称获取对应实例
	impl, exists := registry.GetInstance().Get(request.ServiceName)
	resp := &model.RpcResponse{}
	if !exists {
		resp.Err = errors.New("service not found")
		log.Fatalf("service not found:")
		return resp
	}
	implValue := reflect.ValueOf(impl)
	implType := implValue.Type()
	method, exists := implType.MethodByName(request.MethodName)
	if !exists {
		resp.Err = errors.New("method not found")
		log.Fatalf("method not found")
		return resp
	}

	// 校验参数
	paramsCount := method.Type.NumIn() - 1
	if len(request.Args) != paramsCount {
		resp.Err = errors.New("invalid number of parameters")
		log.Fatalf("invalid number of parameters")
		return resp
	}
	params := make([]reflect.Value, paramsCount)
	for i := 0; i < paramsCount; i++ {
		paramType := method.Type.In(i + 1)
		argValue := reflect.ValueOf(request.Args[i])
		if !argValue.IsValid() || !argValue.Type().AssignableTo(paramType) {
			resp.Err = errors.New("invalid argument")
			log.Fatalf("invalid argument")
			return resp
		}
		params[i] = argValue
	}
	// 调用method
	resultValues := method.Func.Call(append([]reflect.Value{implValue}, params...))
	// 处理返回值
	if len(resultValues) == 0 {
		resp.Err = errors.New("no result")
		log.Fatalf("no result")
		return resp
	}
	// 如果有错误返回值，检查并返回
	if len(resultValues) > 1 {
		lastValue := resultValues[len(resultValues)-1]
		if lastValue.Type().Name() == "error" && !lastValue.IsNil() {
			resp.Err = lastValue.Interface().(error)
			return resp
		}
		resp.Data = resultValues[0].Interface()
		resp.DataType = resultValues[0].Type().Name()
	}

	return resp
}

// doResponse 构造响应
func (h *RequestHandler) doResponse(resp *model.RpcResponse) {
	if resp.Err == nil {
		resp.Message = "success"
	}
	h.writer.Header().Set("Content-Type", "application/json")
	bytes, _ := h.serializer.Serialize(resp)
	log.Println("response: " + string(bytes))
	_, _ = h.writer.Write(bytes)
}
