package main

import "fmt"
import "rsc.io/quote"

//variable
// var : Can be used inside and outside of functions & Variable declaration and value assignment can be done separately
// :=  : Can only be used inside functions &	Variable declaration and value assignment cannot be done separately (must be done in the same line)

var name = "Fiqri Pratama"
var name2 string

// konstanta in golang
const PI = 3.14

func main() {
	fmt.Println(quote.Go())
	fmt.Println(conditional())
	name2 = "Muhammad Zayn"
	name3 := "bundede"

	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(PI)

	output()
	array()
	slices()
	switchCase()
	looping()

	greetings("Fiqri Pratama", 23)
	fmt.Println(sum(3, 5))
	fmt.Println(multiplication(6, 8))
	fmt.Println(substraction(9, 2))
}

func conditional() string {
	text := "overCond"
	num := 5

	if num >= 3 {
		text = "afterCond"
	}
	return text
}

func output() {
	a := "test"
	b := 10

	fmt.Printf("a has value : %v and type %T", a, a)
	fmt.Println()
	fmt.Printf("b has value : %v and type %T", b, b)
	fmt.Println()
}

var familys = [...]string{"Fiqri", "Elda", "Zayn"}

func array() {
	fmt.Println(familys)

	fruits := [...]string{"Watermelon", "Banana", "Orange"}
	fmt.Println(fruits)

	// initialize with specific elements <index:value>
	arr1 := [5]int{1: 10, 2: 40}

	fmt.Println(arr1)

	// get length of arrays
	println(len(familys))
}

func slices() {
	// create slice
	cars := []string{"BMW", "Mazda", "Volvo", "Hyundai", "Brio"}

	// create slice from array
	sliceCars := cars[1:2]

	nums := []int{1, 2, 3, 4, 6, 7, 8, 10}

	// create slice with make: can assign value with set by index or copy from another array value
	numbers := make([]int, 5)

	// copy value array from variable nums to variable numbers
	copy(numbers, nums)

	fmt.Println(sliceCars)
	fmt.Println(numbers)

}

func switchCase() {
	day := 4

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}
}

func looping() {
	// looping even numbers
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	animals := []string{"Bear", "Girrafe", "Mouse"}
	//To only show the value or the index, you can omit the other output using an underscore (_)

	for _, val := range animals {
		fmt.Println(val)
	}
}

func greetings(name string, age int) {
	fmt.Println("Hello My Name Is", name, "And My Age Is", age)
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func multiplication(num1 int, num2 int) (result int) {
	result = num1 * num2
	return
}

func substraction(num1 int, num2 int) (result int, text string) {
	result = num1 - num2
	text = "Successss"
	return
}
