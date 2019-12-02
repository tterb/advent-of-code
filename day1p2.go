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
        for num > 0 {
            num = int(num/3) - 2
            if num > 0 {
                total += num
            }
        }
    }
    fmt.Printf("%d ", total)
}
