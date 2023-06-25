package main

import "fmt"

func main() {
	var a, b = 2, 3
	var c, d float64 = 2.3, 1.1
	// 编译器自动推算出类型
	fmt.Println(getMaxValue(a, b))
	// 自己显示指定类型
	fmt.Println(getMaxValue[float64](c, d))
	fmt.Println(getComparable("hello", "hello"))
}

func getMaxValue[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}
func getComparable[T comparable](a, b T) bool {
	// comparable 仅支持 == 和 !=
	if a == b {
		return true
	}
	return false
}
func getCusMaxValue[T CusNumber](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type CusNumber interface {
	// 用~表示底层数据类型或者说是float64以及其衍生类型
	// | 表示取并集
	// 多行表示取交集
	uint8 | uint16 | uint32 | uint64 | float32 | ~float64
}

// MyFloat64 衍生类型与float64是不同的类型
type MyFloat64 float64

// MyInt64 别名，与int64s同一种类型
type MyInt64 = int64

func sameType[T int | float64, K int | float64](a T, b K) bool {
	return false
}
