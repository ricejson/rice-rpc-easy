package registry

import "sync"

// LocalRegistry 本地服务注册器
type LocalRegistry struct {
	services sync.Map
}

var localRegistry *LocalRegistry
var once sync.Once

// GetInstance 获取唯一实例
func GetInstance() *LocalRegistry {
	once.Do(func() {
		localRegistry = &LocalRegistry{}
	})
	return localRegistry
}

// Register 注册服务
func (r *LocalRegistry) Register(serviceName string, impl any) {
	r.services.Store(serviceName, impl)
}

// Get 获取服务
func (r *LocalRegistry) Get(serviceName string) (any, bool) {
	return r.services.Load(serviceName)
}

// Remove 删除服务
func (r *LocalRegistry) Remove(serviceName string) {
	r.services.Delete(serviceName)
}
