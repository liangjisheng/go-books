package pattern

import "sync"

type singleton struct {

}

var instance *singleton
var once *sync.Once

// GetInstance 懒汉模式, 只在使用的时候才加载
func GetInstance() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}
