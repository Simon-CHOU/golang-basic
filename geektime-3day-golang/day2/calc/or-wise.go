package main

import "fmt"

func xor() {
	arr := []int{2, 2, 3, 5, 3}
	var output int
	for i, item := range arr {
		if i == 0 {
			output = item
			continue
		}
		output = output ^ item
	}
	fmt.Println("不同的位数：", output)
}
