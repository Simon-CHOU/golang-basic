package main

import (
	"fmt"
	"reflect"
)

func main() {
	var hello string = "hello,golang"
	fmt.Println(hello)
	var int2 int = 33
	fmt.Println(int2)
	var float3 float64 = 1.234
	fmt.Println(float3)
	var struct4 struct{} = struct{}{}
	fmt.Println(struct4)
	var map5 map[string]string = map[string]string{}
	fmt.Println(map5)

	var hello1, hello2 string = "你好", "你大爷的"
	fmt.Println(hello1, hello2)
	var (
		hello3 = "3" // 简洁定义，省掉类型，可以根据字面量自动推断类型
		hello4 = "3kda"
		hello5 = "5ll"
		hello6 = "6lldkl3"
		var7   = 777
	)
	fmt.Println(hello3, hello4, hello5, hello6, var7)
	//还可以省掉 var，:= 符号
	studentName := "Alice"
	fmt.Println(studentName)

	//变量名可以为中文
	var 横 int = 1
	var 纵 int = 2
	var 面积 = 横 * 纵
	fmt.Println(面积)

	//变量初始值
	var a int
	var b float64
	var str string    //空字符串
	fmt.Println(a, b) //0 0
	fmt.Println(str)  //空行
	fmt.Println("---------------")
	var mapI map[string]string
	fmt.Println(mapI) //空键值对
	fmt.Println("---------------")

	// var str1, str2 = "a", "b"
	var str3, int4 = "c", 4
	fmt.Println(str3, int4)
	fmt.Println(reflect.TypeOf(str3))
	fmt.Println(reflect.TypeOf(int4)) //使用反射包，获取变量类型

	//类型转换的代价（精度损失）
	fmt.Println("-----类型转换的代价----------")

}
