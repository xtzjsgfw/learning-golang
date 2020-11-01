package singleton

import (
	"fmt"
	"testing"
)

func TestLazyOnce(t *testing.T) {
	instance1 := lazyOnce()
	instance1.Name = "第一次赋值单例模式"
	fmt.Println(instance1.Name)

	instance2 := lazyOnce()
	fmt.Println(instance2.Name)
}

func TestDoubleCheckLock(t *testing.T) {
	s1 := DoubleCheckLock()
	fmt.Println(s1.Name)

	s2 := DoubleCheckLock()
	fmt.Println(s2.Name)
}

func TestHungerOnce(t *testing.T) {
	fmt.Println(hungerInstance.Name)
	if hungerInstance == nil {
		t.Error("单例模式错误")
	}
}

func TestGoOnce(t *testing.T) {
	s1 := GoOnce()
	fmt.Println(s1.Name)

	s2 := GoOnce()
	fmt.Println(s2.Name)
}
