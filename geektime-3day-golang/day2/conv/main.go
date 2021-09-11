package main

import (
	"fmt"
	"math"
)

//类型转换
func main() {
	//精度损失
	var f1 float64 = 1.23456
	var i1 int64 = int64(f1) //小数部分丢失
	fmt.Println(f1, i1)      // 1.23456 1

	//有符号数与无符号数的转化，原码反码补码
	var ui1 uint64 = math.MaxUint64
	var i2 int64 = int64(ui1)
	fmt.Println(ui1, i2) //18446744073709551615 -1

	var i3 int64 = -100
	var ui2 uint64 = uint64(i3)
	fmt.Println(i3, ui2) //-100 18446744073709551516

	// 大数
	var ibig int64 = math.MaxInt64
	var ismall int32 = math.MaxInt32
	var ismall2 int32 = int32(ibig)
	fmt.Println(ibig, ismall, ismall2) //9223372036854775807 2147483647 -1
	var uismall3 uint32 = uint32(ibig)
	fmt.Println(ibig, ismall, uismall3) //9223372036854775807 2147483647 4294967295

	const pi = 3.14
	// 常量不能被赋值 pi = 3.1415 //cannot assign to pi (untyped float constant 3.14)
	// 常量不用不会报 error，只会报warn

}
