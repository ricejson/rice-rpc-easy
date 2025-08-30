package serializer

// Serializer 序列化接口
type Serializer interface {
	// Serialize 序列化
	Serialize(obj any) ([]byte, error)
	// Deserialize 反序列化
	Deserialize(bytes []byte, target any) error
}
