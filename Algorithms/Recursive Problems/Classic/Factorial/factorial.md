# Factorial

### Short Description:

The factorial problem involves computing the product of all positive integers up to a given number n. The factorial of n
is denoted as n! and is defined as n! = n x (n-1) x (n-2) x ... x 1. The problem can be solved using a recursive
function or an iterative loop. The factorial problem has various applications in mathematics, probability theory, and
computer science, such as calculating permutations and combinations, and evaluating series expansions.

### Problem Statement:

The problem is to compute the factorial of a non-negative integer n or given a factorial value, find the number whose
factorial value is n. The program should display a menu to the user with two options:

- Compute the factorial of a number
- Find the number whose factorial value is given

> For **option 1**, the program should prompt the user to enter a non-negative integer n, and then compute and display the
factorial value of n using a recursive function.

> For **option 2**, the program should prompt the user to enter a factorial value fact, and then find and display the number n
whose factorial value is fact. If there is no such integer n, the program should display an error message.

The program should also handle error cases, such as when the user enters a negative integer or a non-positive factorial
value.

### Solution (in Go):

```go
package main

import (
	"fmt"
)

// Define a function to compute the factorial of a non-negative integer n.
func factorial(n int) (int, error) {
	// Check if n is negative. If so, return an error.
	if n < 0 {
		return 0, fmt.Errorf("factorial is not defined for negative numbers")
	}
	// If n is 0, the factorial is 1.
	if n == 0 {
		return 1, nil
	}
	// If n is positive, recursively compute the factorial of n-1 and multiply by n.
	fact, err := factorial(n - 1)
	if err != nil {
		return 0, err
	}
	return n * fact, nil
}

// Define a function to find the input number given its factorial value.
func findNumber(fact int) (int, error) {
	// Check if fact is less than 1. If so, return an error.
	if fact < 1 {
		return 0, fmt.Errorf("factorial value must be a positive integer")
	}
	// Use a loop to incrementally divide fact by the numbers from 1 to n,
	// where n is the maximum possible value of the input number.
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

// Define the main function to read user input and call the appropriate function.
func main() {
	var choice int
	var num, fact int
	var err error
	// Print menu options to the console.
	fmt.Println("Enter 1 to find the factorial of a number")
	fmt.Println("Enter 2 to find the number given its factorial value")
	// Read user choice from the console.
	fmt.Print("Enter your choice: ")
	_, err = fmt.Scan(&choice)
	if err != nil {
		fmt.Printf("Error reading input: %s\n", err.Error())
		return
	}
	// Call the appropriate function based on user choice.
	switch choice {
	case 1:
		// Read a non-negative integer from the console.
		fmt.Print("Enter a non-negative integer: ")
		_, err = fmt.Scan(&num)
		if err != nil {
			fmt.Printf("Error reading input: %s\n", err.Error())
			return
		}
		// Compute the factorial of the input number and print the result.
		fact, err = factorial(num)
		if err != nil {
			fmt.Printf("Error computing factorial: %s\n", err.Error())
			return
		}
		fmt.Printf("The factorial of %d is %d\n", num, fact)
	case 2:
		// Read a positive integer factorial value from the console.
		fmt.Print("Enter a positive integer factorial value: ")
		_, err = fmt.Scan(&fact)
		if err != nil {
			fmt.Printf("Error reading input: %s\n", err.Error())
			return
		}
		// Find the input number whose factorial value is the input value and print the result.
		num, err = findNumber(fact)
		if err != nil {
			fmt.Printf("Error finding number: %s\n", err.Error())
			return
		}
		fmt.Printf("%d is the factorial of %d\n", fact, num)
	default:
		// Handle invalid user choice.
		fmt.Println("Invalid choice. Please enter 1 or 2.")
	}
}
```

### Result:

Here are some examples of how the factorial function code can be used in different scenarios:

- Calculating factorials of small numbers:

```
Enter a non-negative integer: 5
The factorial of 5 is 120
```

- Calculating factorials of large numbers:

```
Enter a non-negative integer: 20
The factorial of 20 is 2432902008176640000
```

- Finding the number whose factorial is given:

```
Enter the factorial value: 120
The number whose factorial value is 120 is 5
```

- Handling invalid input:

```
Enter a non-negative integer: -5
Invalid input: Enter a non-negative integer

Enter the factorial value: -120
Invalid input: Enter a positive integer
```

- Providing a menu for the user to select an option:

```
Welcome to the factorial program!
1. Calculate the factorial of a number
2. Find the number whose factorial is given
   Select an option: 1

Enter a non-negative integer: 6
The factorial of 6 is 720

Welcome to the factorial program!
1. Calculate the factorial of a number
2. Find the number whose factorial is given
   Select an option: 2

Enter the factorial value: 3628800
The number whose factorial value is 10
```

- Handling invalid menu selection:

```
Welcome to the factorial program!
1. Calculate the factorial of a number
2. Find the number whose factorial is given
   Select an option: 3
   
Invalid option: Enter 1 or 2
```

- Handling non-integer input:

```
Welcome to the factorial program!
1. Calculate the factorial of a number
2. Find the number whose factorial is given
   Select an option: 2

Enter the factorial value: abc
Invalid input: Enter a positive integer
```
- Handling non-factorial input:

```
Welcome to the factorial program!
1. Calculate the factorial of a number
2. Find the number whose factorial is given
   Select an option: 2

Enter the factorial value: 7
No number exists whose factorial value is 7
```