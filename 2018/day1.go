package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    var inp []string
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        inp = append(inp, scanner.Text())
    }
    
    total := 0
    for _, i := range inp {
        val, _ := strconv.Atoi(i)
        total += val
    }
    fmt.Printf("%d ", total)
}