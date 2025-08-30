package serializer

import "encoding/json"

// NativeSerializer Go原生序列化器
type NativeSerializer struct {
}

func NewNativeSerializer() *NativeSerializer {
	return &NativeSerializer{}
}

func (s *NativeSerializer) Serialize(obj any) ([]byte, error) {
	return json.Marshal(obj)
}

func (s *NativeSerializer) Deserialize(bytes []byte, target any) error {
	return json.Unmarshal(bytes, target)
}
