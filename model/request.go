package model

type RpcRequest struct {
	ServiceName string `json:"serviceName"`
	MethodName  string `json:"methodName"`
	ArgsType    []any  `json:"argsType"`
	Args        []any  `json:"args"`
}

func NewRpcRequest(serviceName string, methodName string, argsType []any, args []any) *RpcRequest {
	return &RpcRequest{
		ServiceName: serviceName,
		MethodName:  methodName,
		ArgsType:    argsType,
		Args:        args,
	}
}
