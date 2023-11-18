package main

import (
	"fmt"
)

/*
代码中将要删除元素的后续元素前移，然后缩容切片，实现切片中元素的删除。
这种实现的优势是不需要重新分配内存，因此效率较高。
删除元素的函数使用了泛型函数，可以处理任意类型的切片。
代码将要删除元素的后续元素前移，实现切片的缩容。
*/
func main() {
	nuSlice := []int{1, 2, 3, 4, 5}

	// 删除切片的第一个元素
	nuSlice = removeAt(nuSlice, 0)
	fmt.Println(nuSlice) // [2 3 4 5]

	// 删除切片的最后一个元素
	nuSlice = removeAt(nuSlice, len(nuSlice)-1)
	fmt.Println(nuSlice) // [2 3 4]

	// 删除切片的中间元素
	nuSlice = removeAt(nuSlice, 2)
	fmt.Println(nuSlice) // [2 4]
}

// 删除切片特定下标元素的方法
func removeAt[T any](nuSlice []T, index int) []T {
	// 判断下标是否越界
	if index < 0 || index >= len(nuSlice) {
		return nuSlice
	}

	// 将要删除元素的后续元素前移
	for i := index; i < len(nuSlice)-1; i++ {
		nuSlice[i] = nuSlice[i+1]
	}

	// 缩容切片
	nuSlice = nuSlice[:len(nuSlice)-1]

	return nuSlice
}
