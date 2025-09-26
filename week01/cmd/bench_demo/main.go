package main

import (
	"fmt"
	"github.com/canteen_ns/go-container-lab/week01/pkg/builtin"
)

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := builtin.Array2Slice(&arr)
	fmt.Println("array->slice zero-copy:", slice)
	fmt.Printf("s  descriptor: %#v\n", slice) // 切片头: []int{len:4, cap:4, ptr:0xc0000102c0}
	fmt.Printf("s  data addr: %p\n", slice)   // 打印的是 **ptr字段** → 底层数组首元素地址
	fmt.Printf("&s:         %p\n", &slice)    // 打印的是切片头地址, 切片头大小固定，底层数据地址+数据长度+容量信息
}
