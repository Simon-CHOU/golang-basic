package main

import "fmt"

func main() {
	var left, right int = 2, 3
	// var value int = left + right
	// Go也支持短变量声明（就是带冒号的），编译器可以推断类型（通过后面的赋值自动推断出类型），不用我们显式声明了
	value := left + right
	fmt.Println(value)        //5
	fmt.Println(left / right) //0
	var leftF, rightF float32 = 2, 3
	fmt.Println(leftF / rightF) //0.6666667
	fmt.Println(left % right)   //2
	//fmt.Println(leftF % rightF) // 浮点数不支持取模 invalid operation: operator % not defined for leftF (variable of type float32)

	//比较运算符
	fmt.Println(true && false)
	var k, f bool = true, false
	fmt.Println(k || f)
	fmt.Println(!k)
	if k == true {
		fmt.Println("k is good.")
	}
	fmt.Println(left*right == 6)
	fmt.Println(left*right < 6)
	fmt.Println(left*right <= 6)
	fmt.Println(left*right > 6)
	fmt.Println(left*right >= 6)

	//位运算符
	xor() //异或
}
