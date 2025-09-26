package builtin

import (
	"reflect"
	"unsafe"
)

// String2Bytes 实现字符串到字节切片的零拷贝转换
// 注意：
//  1. 结果切片与原字符串共享内存空间
//  2. 禁止对结果切片进行写操作（会导致不可预知行为）
//  3. Go 1.20+ 建议使用 unsafe.StringData 替代
func String2Bytes(s string) []byte {
	// 步骤分解：
	// 1. 获取字符串头指针
	// 2. 转换为 StringHeader 获取底层数据指针
	// 3. 将数据指针转为 byte 类型指针
	// 4. 创建指定长度的切片
	return unsafe.Slice(
		(*byte)(unsafe.Pointer(
			(*reflect.StringHeader)(unsafe.Pointer(&s)).Data)), len(s))
}
