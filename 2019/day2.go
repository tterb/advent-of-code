package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    opcodes := []int{}
    scanner := bufio.NewScanner(os.Stdin)
    // Read and process input from stdin
    for scanner.Scan() {
        vals := strings.Split(scanner.Text(), ",")
        for _, i := range vals {
            num, _ := strconv.Atoi(i)
            opcodes = append(opcodes, num)
        }
    }
    // replace position 1 with the value 12 and
    // replace position 2 with the value 2
    opcodes[1], opcodes[2] = 12, 2
    index := 0
    for index < len(opcodes)-1 {
        var op, a, b, out int = opcodes[index], opcodes[index+1], opcodes[index+2], opcodes[index+3]
        if op == 1 {
            opcodes[out] = opcodes[a] + opcodes[b]
        } else if op == 2 {
            opcodes[out] = opcodes[a] * opcodes[b]
        }
        index += 4
    } 
    // Print all final values
    // for _, i := range opcodes {
    //     fmt.Printf("%d ", i)
    // }
    fmt.Printf("%d", opcodes[0])
}
