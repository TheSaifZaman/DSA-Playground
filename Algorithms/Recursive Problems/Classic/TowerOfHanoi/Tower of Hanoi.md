# Tower of Hanoi

### Short Description:

The Tower of Hanoi problem falls under the category of classic recursive problems in computer science. It is a
well-known problem that involves solving a puzzle that consists of three rods and a number of disks of different sizes,
initially stacked in decreasing size order on one rod. The goal of the puzzle is to move the entire stack of disks to
another rod, subject to the constraint that only one disk can be moved at a time, and no disk can be placed on top of a
smaller disk. This problem has been studied extensively in computer science, mathematics, and other fields, and has many
applications in algorithm design, optimization, and game theory.

### Problem Statement:

The Tower of Hanoi is a classic mathematical puzzle that involves moving a stack of disks of different sizes from one
rod to another, subject to the following constraints:

- Only one disk can be moved at a time.
- A disk can only be placed on top of a larger disk or an empty rod.

The puzzle consists of three rods and a set of disks of different sizes, initially stacked in decreasing size order on
one of the rods. The objective is to move the entire stack of disks to another rod, while following the above
constraints.

The problem statement can be generalized to a version where there are more than n rods, and the goal is to move the
stack of disks to a specified destination rod, subject to the same constraints. In this generalized version, the number
of rods and the number and sizes of the disks can be specified as inputs to the problem.

### Solution (in Go):

Here's an example program in Go that solves the Tower of Hanoi problem with n disks and m towers:

```go
/*
This line declares the package of the program. 
In Go, the main package is a special package that is used to build standalone executable programs.
*/
package main

/*
This line imports the "fmt" package, which provides formatted I/O functions like "Scan" and "Printf".
*/
import "fmt"

/*
This line declares a function named "towerOfHanoi". 
The function takes four parameters: "n" represents the number of disks, "source" represents the source tower number, 
"destination" represents the destination tower number, and "towers" is a slice of integers that represents the state of the towers.
*/
func towerOfHanoi(n, source, destination int, towers []int) {
    /*
    This line checks if the number of disks is zero, which means there is nothing to move. 
    If so, the function simply returns without doing anything.
    */
    if n == 0 {
        return
    }
    
    /*
    This line calls the "getAuxiliary" function to get the tower number of an auxiliary tower that is not the source or destination tower. 
    The function returns the tower number as an integer, which is assigned to the "auxiliary" variable using the short variable declaration syntax ":=".
    */
    auxiliary := getAuxiliary(source, destination, towers)
    
    /*
    This line gets the length of the "towers" slice, which represents the number of towers in the puzzle. 
    The length is assigned to the "k" variable using the short variable declaration syntax.
    */
    k := len(towers)
    
    /*
    This line calls the "towerOfHanoi" function recursively with the top n-1 disks moved from the source tower to the auxiliary tower, 
    using the destination tower as the auxiliary tower.
    */
    towerOfHanoi(n-1, source, auxiliary, towers)
    
    /*
    This line prints a message to the console using the "Printf" function from the "fmt" package. 
    The message displays which disk is moved from which tower to which tower.
    */
    fmt.Printf("Move disk %d from Tower %d to Tower %d\n", n, source, destination)
    
    /*
    This line of code updates the state of the towers by moving the current disk (represented by the integer value "n") from the source tower to the destination tower. 
    The towers slice represents the state of all the towers in the game, where each index of the slice represents a tower number, and the value at that index represents the topmost disk on that tower.

    The expression (destination-1+k)%k calculates the index of the destination tower in the towers slice, taking into account the wrap-around behavior when the destination tower is the last tower. 
    Here's how it works:

    - Subtract 1 from the destination tower number to make it 0-indexed instead of 1-indexed, because slices in Go are 0-indexed.
    - Add k to the result to ensure that the index is non-negative.
    - Take the remainder of the result when divided by k, which ensures that the index wraps around to the beginning of the slice when it goes past the end.
    
    So the expression (destination-1+k)%k calculates the index of the destination tower in the towers slice, 
    and the line of code towers[(destination-1+k)%k] = n updates that index with the value of n, which represents the topmost disk on the source tower that is being moved to the destination tower. 
    By doing this, the function effectively moves the topmost disk from the source tower to the destination tower in the game.
    */
    towers[(destination-1+k)%k] = n
    
    /*
    This line calls the "towerOfHanoi" function recursively with the top n-1 disks moved from the auxiliary tower to the destination tower, using the source tower as the auxiliary tower.
    */
    towerOfHanoi(n-1, auxiliary, destination, towers)
}

/*
This lines declares a function named "getAuxiliary". 
The function takes three parameters: "source" represents the source tower number, "destination" represents the destination tower number, and "towers" is a slice of integers that represents the state of the towers. 
The function returns the tower number of an auxiliary tower that is not the source or destination tower.
*/
func getAuxiliary(source, destination int, towers []int) int {
    for i, t := range towers {
        if t != source && t != destination {
            return i+1
        }
    }
    return -1
}

/*
This is the main function where we take input of the disks count and number of towers and call the towerOfHanoi() function.
*/
func main() {
    var n, m int
    fmt.Print("Enter the number of disks: ")
    fmt.Scan(&n)
    fmt.Print("Enter the number of towers: ")
    fmt.Scan(&m)
    towers := make([]int, m)
    towerOfHanoi(n, 1, m, towers)
}
```

### Result:

```
For disks = 3, tower = 3

Result:

Move disk 1 from Tower 1 to Tower 1
Move disk 2 from Tower 1 to Tower 1
Move disk 1 from Tower 1 to Tower 1
Move disk 3 from Tower 1 to Tower 3
Move disk 1 from Tower 1 to Tower 2
Move disk 2 from Tower 1 to Tower 3
Move disk 1 from Tower 2 to Tower 3
```

