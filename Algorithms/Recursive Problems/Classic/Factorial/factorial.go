package main

import (
	"fmt"
)

func factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("factorial is not defined for negative numbers")
	} else if n == 0 {
		return 1, nil
	} else {
		fact, err := factorial(n - 1)
		if err != nil {
			return 0, err
		}
		return n * fact, nil
	}
}

func findNumber(fact int) (int, error) {
	if fact < 1 {
		return 0, fmt.Errorf("factorial value must be a positive integer")
	}
	n := 0
	for i := 1; fact > 0; i++ {
		if fact%i != 0 {
			return 0, fmt.Errorf("%d is not a valid factorial value", fact)
		}
		n++
		fact /= i
	}
	return n, nil
}

func main() {
	var choice int
	var num, fact int
	var err error
	fmt.Println("Enter 1 to find the factorial of a number")
	fmt.Println("Enter 2 to find the number given its factorial value")
	fmt.Print("Enter your choice: ")
	_, err = fmt.Scan(&choice)
	if err != nil {
		fmt.Printf("Error reading input: %s\n", err.Error())
		return
	}
	switch choice {
	case 1:
		fmt.Print("Enter a non-negative integer: ")
		_, err = fmt.Scan(&num)
		if err != nil {
			fmt.Printf("Error reading input: %s\n", err.Error())
			return
		}
		fact, err = factorial(num)
		if err != nil {
			fmt.Printf("Error computing factorial: %s\n", err.Error())
			return
		}
		fmt.Printf("The factorial of %d is %d\n", num, fact)
	case 2:
		fmt.Print("Enter a positive integer factorial value: ")
		_, err = fmt.Scan(&fact)
		if err != nil {
			fmt.Printf("Error reading input: %s\n", err.Error())
			return
		}
		num, err = findNumber(fact)
		if err != nil {
			fmt.Printf("Error finding number: %s\n", err.Error())
			return
		}
		fmt.Printf("%d is the factorial of %d\n", fact, num)
	default:
		fmt.Println("Invalid choice. Please enter 1 or 2.")
	}
}
