package singleton

import "sync"

var instance *singleton
var hungerInstance *singleton
var goInstance *singleton
var mux sync.Mutex
var once sync.Once

type singleton struct {
	Name string
}

func lazyOnce() *singleton {
	// 存在线程安全问题
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

// 改善lazyOnce
func lockLazyOnce() *singleton {
	mux.Lock()
	defer mux.Unlock()
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

// 改善lazyOnceLock
func DoubleCheckLock() *singleton {
	if instance == nil {
		mux.Lock()
		defer mux.Unlock()
		if instance == nil {
			instance = new(singleton)
		}
	}
	return instance
}

func init() {
	hungerInstance = NewHungerOnce()
}

func NewHungerOnce() *singleton {
	return &singleton{
		Name: "饿汉模式：在包加载时通过init()创建单例，相较于懒汉式更加安全，但是当程序不需要时浪费空间",
	}
}

func GoOnce() *singleton {
	once.Do(func() {
		goInstance = &singleton{
			"使用sync.Once",
		}
	})
	return goInstance
}
