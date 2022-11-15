package main

import "fmt"

func main() {
	var son int
	fmt.Printf("Son kiriting ")
	fmt.Scanln(&son)
	fibonachi(son)
}

func fibonachi(son int) {
	var first_num int = 0
	var second_num int = 1
	fmt.Println(first_num)
	fmt.Println(second_num)
	first_p := &first_num
	second_p := &second_num
	for second_num <= son {
		add := first_num + second_num
		fmt.Println(add)
		*first_p = second_num
		*second_p = add
	}
}
