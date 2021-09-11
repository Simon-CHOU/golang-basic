package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 100; i++ {
		fmt.Println(i, "你好 golang")
		if i == 5 {
			break
		}
	}

	var k = 10
	for k > 0 {
		if k%3 == 0 {
			fmt.Println(k, "拖拽上传")
		}
		k--
		fmt.Println(k)
	}

	str1, str3 := "我是", "Gopher"
	fmt.Println(str1 + str3)
	fmt.Printf("%s\t\\%%%s\n", str1, str3)

	var area float64 = 2.0 / 6
	fmt.Printf("%f\n", area)
	fmt.Printf("%.2f\n", area)
	fmt.Printf("%9.2f\n", area)
	fmt.Printf("%b\n", area)
	fmt.Printf("%E\n", area)
	fmt.Printf("%e\n", area)
	fmt.Printf("%G\n", area)
	fmt.Printf("%g\n", area)

	//计算两个圆的面积并输出，精确到小数点后三位
	var r int = 23
	const pi float32 = 3.14159
	fmt.Printf("%.3f", float32(r*r)*pi)

}
