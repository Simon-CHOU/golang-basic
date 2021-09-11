package main

import "fmt"

func main() {
	var buyFruit string = "六个苹果"
	var see_watermelon bool = false
	// var see_watermelon bool = true
	var banna bool = true
	if see_watermelon {
		buyFruit = "一个苹果"
	} else if banna == true {
		buyFruit = "10个苹果"
	} else {
		buyFruit = "两个橘子"
	}
	fmt.Println(buyFruit) //10个苹果
}
