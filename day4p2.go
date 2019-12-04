package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    var min, max int
    scanner := bufio.NewScanner(os.Stdin)
    // Read and process input from stdin
    for scanner.Scan() {
        vals := strings.Split(scanner.Text(), "-")
        min, _ = strconv.Atoi(vals[0])
        max, _ = strconv.Atoi(vals[1])
    }
    
    passwords := 0
    // Increment through and check values for valid passwords
    for min < max {
        if isValid(min) {
            passwords += 1
        }
        min++
    }
    fmt.Printf("%d passwords", passwords)
}

func isValid(num int) bool {
    digits := intToSlice(num, []int{})
    hasAdjacent := false
    i := 0
    for i < len(digits)-1 {
        if digits[i] > digits[i+1] {
            return false
        } else if digits[i] == digits[i+1] {
            // Check that equal adjacent values aren't part of a larger group
            if (i == 0 || digits[i-1] != digits[i]) && (i == len(digits)-2 || digits[i+2] != digits[i]) {
                hasAdjacent = true
            }
        }
        i++
    }
    return hasAdjacent
}

func intToSlice(n int, digits []int) []int {
    if n != 0 {
        i := n % 10
        digits = append([]int{i}, digits...)
        return intToSlice(n/10, digits)
    }
    return digits
}
