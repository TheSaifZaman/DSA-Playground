# Fibonacci Series

### Short Description:

Fibonacci is a sequence of numbers where each number is the sum of the previous two numbers in the sequence. The
sequence starts with 0 and 1, and each subsequent number is the sum of the previous two numbers. The sequence goes 0, 1,
1, 2, 3, 5, 8, 13, 21, 34, and so on. This sequence has numerous applications in mathematics, science, and computer
science, such as in coding theory, number theory, and algorithm design.

### Problem Statement:

Given a positive integer n, find the nth Fibonacci number, or given a Fibonacci number f, find the position or index of
f in the Fibonacci sequence.

### Solution (in Go):

```go
package main

import (
    "fmt"
)

// Define the main function
func main() {
    var input int
    fmt.Println("Welcome to the Fibonacci program!")
    fmt.Println("Select an option:")
    fmt.Println("1. Find the nth Fibonacci number")
    fmt.Println("2. Find the position of a Fibonacci number")

    // Read user input and check for errors
    _, err := fmt.Scanln(&input)
    if err != nil {
        fmt.Println("Invalid input: Enter a number")
        return
    }

    // Use a switch statement to handle the user's choice
    switch input {
    case 1:
        var n int
        fmt.Println("Enter a non-negative integer:")
        _, err := fmt.Scanln(&n)
        if err != nil || n < 0 {
            fmt.Println("Invalid input: Enter a non-negative integer")
            return
        }
        // Call the fibonacci function and print the result
        fib := fibonacci(n)
        fmt.Printf("The %dth Fibonacci number is %d\n", n, fib)
    case 2:
        var f int
        fmt.Println("Enter a Fibonacci number:")
        _, err := fmt.Scanln(&f)
        if err != nil || f < 0 {
            fmt.Println("Invalid input: Enter a non-negative integer")
            return
        }
        // Call the findFibonacciPosition function and print the result
        pos := findFibonacciPosition(f)
        if pos == -1 {
            fmt.Printf("No Fibonacci number exists with the value %d\n", f)
        } else {
            fmt.Printf("The Fibonacci number %d is at position %d\n", f, pos)
        }
    default:
        fmt.Println("Invalid option: Enter 1 or 2")
    }
}

// Define the fibonacci function to recursively calculate the nth Fibonacci number
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

// Define the findFibonacciPosition function to iteratively find the position of a given Fibonacci number
func findFibonacciPosition(f int) int {
    a := 0
    b := 1
    pos := 1
    for b < f {
        c := a + b
        a = b
        b = c
        pos++
    }
    if b == f {
        return pos
    } else {
        return -1
    }
}
```

### Result:

Here are some example usages of the previous code:

- Example 1:

```
Welcome to the Fibonacci program!
Select an option:
1. Find the nth Fibonacci number
2. Find the position of a Fibonacci number
   1
   Enter a non-negative integer:
   7
   The 7th Fibonacci number is 13
```

- Example 2:

```
Welcome to the Fibonacci program!
Select an option:
1. Find the nth Fibonacci number
2. Find the position of a Fibonacci number
   2
   Enter a Fibonacci number:
   21
   The Fibonacci number 21 is at position 8
```   

- Example 3:

```
Welcome to the Fibonacci program!
Select an option:
1. Find the nth Fibonacci number
2. Find the position of a Fibonacci number
   3
   Invalid option: Enter 1 or 2
```