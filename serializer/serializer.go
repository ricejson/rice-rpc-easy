package serializer

import "encoding/json"

// NativeSerializer Go原生序列化器
type NativeSerializer struct {
}

func (s *NativeSerializer) serialize(obj any) ([]byte, error) {
	return json.Marshal(obj)
}

func (s *NativeSerializer) deserialize(bytes []byte, target any) error {
	return json.Unmarshal(bytes, target)
}
