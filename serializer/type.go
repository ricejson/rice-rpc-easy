package serializer

// Serializer 序列化接口
type Serializer interface {
	// serialize 序列化
	serialize(obj any) ([]byte, error)
	// deserialize 反序列化
	deserialize(bytes []byte, target any) error
}
