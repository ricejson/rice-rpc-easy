package model

type RpcResponse struct {
	Data     any    `json:"data"`
	DataType any    `json:"dataType"`
	Message  string `json:"message"`
	Err      error  `json:"err"`
}

func NewRpcResponse(data any, dataType any, message string, err error) *RpcResponse {
	return &RpcResponse{
		Data:     data,
		DataType: dataType,
		Message:  message,
		Err:      err,
	}
}
