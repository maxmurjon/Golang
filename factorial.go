package main

import "fmt"

func main() {
	var son int
	fmt.Printf("Son kiriting ")
	fmt.Scanln(&son)
	fmt.Println(factorial(son))

	//fibonachi(21)
}
func factorial(son int) int {
	if son > 0 {
		result := 1
		for i := 1; i != son+1; i++ {
			result *= i
		}
		return result
	} else if son == 0 {
		return 1
	} else {
		result2 := 1
		for i := -1; i != (son - 1); i-- {
			result2 *= i
		}
		return result2
	}

}
