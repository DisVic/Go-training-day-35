package main

import "fmt"

func main() {
	powByTwoNumber()
	powByTwoNumber()
	powByTwoNumber()
	//fn := powByTwoNumber()
	/*fn()
	fn()
	fn()*/
}

func powByTwoNumber() func() {
	number := 5
	fmt.Println(number)
	return func() {
		number *= 2
		fmt.Println(number)
	}

}
