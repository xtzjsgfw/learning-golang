###单例模式

*保证只有一个实例*

####懒汉式
存在线程安全问题
```go
var instance *singleton

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
```


####饿汉式
```go
var hungerInstance *singleton

func init() {
	hungerInstance = NewHungerOnce()
}

func NewHungerOnce() *singleton {
	return &singleton{
		Name: "饿汉模式：在包加载时通过init()创建单例，相较于懒汉式更加安全，但是当程序不需要时浪费空间",
	}
}
```
饿汉模式：在包加载时通过init()创建单例，相较于懒汉式更加安全，但是当程序不需要时浪费空间


####双重锁检查机制
```go
var mux sync.Mutex
// 改善lazyOnce
func lockLazyOnce() *singleton {
	mux.Lock()
	defer mux.Unlock()
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}
```
每次执行时都加锁，开销大

```go
var mux sync.Mutex
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
```
只有确实没有instance的情况才加锁

####sync.Once原子操作
```go
var goInstance *singleton
var once sync.Once
func GoOnce() *singleton {
	once.Do(func() {
		goInstance = &singleton{
			"使用sync.Once",
		}
	})
	return goInstance
}
```