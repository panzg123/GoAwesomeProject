package main

import (
	"fmt"
	"sync"
	"time"
)

// Container 测试数据结构
type Container struct {
	sync.Mutex
	counters map[string]int
	ID int
	Name string
}

// incWrong 错误的接收方式
// c 値接受者
// Container其中包含lock，不能值传递，报错 fatal error: concurrent map writes
func (c Container) incWrong(name string) {
	c.Lock()
	defer c.Unlock()
	c.counters[name]++
}

// inc c指针接收器，正确的做法
func (c *Container) inc(name string) {
	c.Lock()
	defer c.Unlock()
	c.counters[name]++
}

func (c Container) byValMethod() {
	fmt.Printf("byValMethod got &c=%p, &(c.s)=%p\n", &c, &(c.Name))
}

func (c *Container) byPtrMethod() {
	fmt.Printf("byPtrMethod got &c=%p, &(c.s)=%p\n", c, &(c.Name))
}

// https://mp.weixin.qq.com/s/CvHxFrIWlA70TYElPE3-cg
// c是値接收器，非指针接收器，拷贝mutex
func main() {
	c := Container{counters: map[string]int{"a": 0, "b": 0}}
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
	}
	go doIncrement("a", 100000)
	go doIncrement("a", 100000)

	time.Sleep(1*time.Second)
	fmt.Println(c.counters)

	// 打印下两者的地址的，指针传递是一样的地址，值传递只是拷贝
	// in main &c=0xc0000783c0, &(c.s)=0xc0000783d8
	// byValMethod got &c=0xc000188090, &(c.s)=0xc0001880a8
	// byPtrMethod got &c=0xc0000783c0, &(c.s)=0xc0000783d8
	fmt.Printf("in main &c=%p, &(c.s)=%p\n", &c, &(c.Name))
	c.byValMethod()
	c.byPtrMethod()
}