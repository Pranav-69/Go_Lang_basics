package main

import "fmt"

// func sum(number int) int {
// 	if number <= 0 {
// 		return 0
// 	} else {
// 		return number + sum(number-1)

// 	}
// }

func findSquare(num int) int {
	square := num * num
	return square
}

var sum = 0

func main() {
	// var num int
	// fmt.Println("Enter the value of num :")
	// fmt.Scanf("%d", &num)
	// if num < 0 {
	// 	fmt.Println("Invalid Input!! Please try again")
	// 	return
	// }

	// result := sum(num)
	// // var result = sum(num)
	// fmt.Println("Sum:", result)
	// var greet = func() {
	// 	fmt.Println("Function without name")
	// }
	// // welcome := greet
	// var welcome = greet
	// greet()
	// welcome()
	// var sum = func(n1, n2 int) {
	// 	s := n1 + n2
	// 	fmt.Println("Sum is : ", s)

	// }
	// fmt.Printf("%T\n", sum)
	// sum(5, 5)

	// var sum = func(n1, n2 int) int {
	// 	sum := n1 + n2
	// 	return sum
	// }
	// result := sum(3, 1)
	// fmt.Println("Sum is :", result)

	sum := func(n1, n2 int) int {
		return n1 + n2
	}
	
	result := findSquare(sum(6, 9))
	fmt.Println("Result is ", result)

}
