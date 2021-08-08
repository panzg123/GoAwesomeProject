package main

import (
	"encoding/json"
	"fmt"
)

// Test 按照对齐原则，Test 变量占用 16 字节的内存
type Test struct {
	F1 uint64
	F2 uint32
	F3 byte
}


// https://mp.weixin.qq.com/s/28ScdoPWrQ2t870GtgLX1Q
// 自定义序列化方式
func main() {
	t := Test{F1: 0x1234, F2: 0x4567, F3: 12,}

	// 测试序列化
	bs, err := json.Marshal(&t)
	if err != nil {
		panic("")
	}
	fmt.Printf("t -> []byte\t: %v\n", bs)

	// 测试反序列化
	t1 := Test{}
	err = json.Unmarshal(bs, &t1)
	if err != nil {
		panic("")
	}
	fmt.Printf("[]byte -> t1\t: %v\n", t1)
}