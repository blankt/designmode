package singleton

import "sync"

// Singleton 使用single接口防止返回的对象为一个包私有对象？
type Singleton interface {
	foo()
}

type singleton struct{}

func (s singleton) foo() {
}

var (
	instance *singleton
	once     sync.Once
)

func NewInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}
