package main

import (
    "fmt"
)

func main() {
    var input int
    fmt.Println("Welcome to the Fibonacci program!")
    fmt.Println("Select an option:")
    fmt.Println("1. Find the nth Fibonacci number")
    fmt.Println("2. Find the position of a Fibonacci number")
    _, err := fmt.Scanln(&input)
    if err != nil {
        fmt.Println("Invalid input: Enter a number")
        return
    }

    switch input {
    case 1:
        var n int
        fmt.Println("Enter a non-negative integer:")
        _, err := fmt.Scanln(&n)
        if err != nil || n < 0 {
            fmt.Println("Invalid input: Enter a non-negative integer")
            return
        }
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

// Recursive function to calculate the nth Fibonacci number
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

// Function to find the position of a given Fibonacci number
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
