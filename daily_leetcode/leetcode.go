package main

import (
	"fmt"
	"reflect"
)

// firstUniqChar 找到第一个不重复的字符
func firstUniqChar(s string) int {
	MapCnt := make(map[uint8]uint32)
	for i := 0; i < len(s); i++ {
		MapCnt[s[i]] += 1
	}
	for i := 0; i < len(s); i++ {
		if MapCnt[s[i]] == 1 {
			return i
		}
	}
	return -1
}

// caseReflect 测试反射
func caseReflect() {
	var f = 1.0
	fmt.Println(reflect.TypeOf(f))
	v := reflect.ValueOf(f)
	fmt.Println(v.Kind())
	fmt.Println(v.Float())
}

func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// candy 135.分发糖果
func candy(ratings []int) int {
	val := make([]int, len(ratings))
	for i := range val {
		val[i] = 1
	}
	//从左往右
	for i := 1; i < len(val); i++ {
		if ratings[i] > ratings[i-1] {
			val[i] = IntMax(val[i-1]+1, val[i])
		}
	}
	//从右往左
	for i := len(val) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			val[i] = IntMax(val[i+1]+1, val[i])
		}
	}
	//累加
	var t = 0
	for _, v := range val {
		t += v
	}
	return t
}

// decodeString 394. 字符串解码
func decodeString(s string) string {
	return s
}

func main() {
	decodeString("3[a]2[bc]")
	data := map[string]string{"1":"2"}
	fmt.Println(data)
}
