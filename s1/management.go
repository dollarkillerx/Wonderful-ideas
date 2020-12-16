package plugin

import (
	"fmt"
	"sync"
)

type RegistryMar struct {
	persistence map[string]Plugin
	mu          sync.Mutex
}

func New() *RegistryMar {
	return &RegistryMar{persistence: map[string]Plugin{}}
}

// 插件注册
func (r *RegistryMar) Registry(key string, plugin Plugin) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.persistence[key] = plugin
}

// 初始化插件
func (r *RegistryMar) GetPlugin(key string) (plugin Plugin, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	plugin, ex := r.persistence[key]
	if !ex {
		return nil, fmt.Errorf("plugin not found")
	}
	return plugin, nil
}
