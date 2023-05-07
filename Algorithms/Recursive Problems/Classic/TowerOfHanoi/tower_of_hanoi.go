package main

import "fmt"

func towerOfHanoi(n, source, destination int, towers []int) {
    if n == 0 {
        return
    }
    auxiliary := getAuxiliary(source, destination, towers)
    k := len(towers)
    towerOfHanoi(n-1, source, auxiliary, towers)
    fmt.Printf("Move disk %d from Tower %d to Tower %d\n", n, source, destination)
    towers[(destination-1+k)%k] = n
    towerOfHanoi(n-1, auxiliary, destination, towers)
}

func getAuxiliary(source, destination int, towers []int) int {
    for i, t := range towers {
        if t != source && t != destination {
            return i+1
        }
    }
    return -1
}

func main() {
    var n, m int
    fmt.Print("Enter the number of disks: ")
    fmt.Scan(&n)
    fmt.Print("Enter the number of towers: ")
    fmt.Scan(&m)
    towers := make([]int, m)
    towerOfHanoi(n, 1, m, towers)
}
