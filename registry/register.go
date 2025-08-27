package registry

import "sync"

// LocalRegistry 本地服务注册器
type LocalRegistry struct {
	services sync.Map
}

func NewLocalRegistry() *LocalRegistry {
	return &LocalRegistry{}
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
