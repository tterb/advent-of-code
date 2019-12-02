package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    total := 0
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        num, _ := strconv.Atoi(scanner.Text())
        total += int(num/3) - 2
    }
    fmt.Printf("%d ", total)
}
