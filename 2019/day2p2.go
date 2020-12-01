package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    target := 19690720
    origin := []int{}
    scanner := bufio.NewScanner(os.Stdin)
    // Read and process input from stdin
    for scanner.Scan() {
        vals := strings.Split(scanner.Text(), ",")
        for _, i := range vals {
            num, _ := strconv.Atoi(i)
            origin = append(origin, num)
        }
    }

    opcodes := make([]int, len(origin))
    limit := 99
    if len(origin) < 99 {
        limit = len(origin)
    }
    for noun := 0; noun <= limit; noun++ {
        for verb := 0; verb <= limit; verb++ {
            // Restore original memory values
            copy(opcodes, origin)
            opcodes[1], opcodes[2] = noun, verb
            output := calc_opcodes(opcodes)
            if output == target {
                fmt.Printf("%d", (100*noun)+verb)
                os.Exit(0)
            }
        }
    }
}

func calc_opcodes(opcodes []int) int {
    pointer := 0
    instruction_count := 4
    for pointer < len(opcodes)-1 {
        var op, a, b, out = opcodes[pointer], opcodes[pointer+1], opcodes[pointer+2], opcodes[pointer+3]
        if op == 99 {  // Halt
            break
        } else if op == 1 {  // Add
            opcodes[out] = opcodes[a] + opcodes[b]
            instruction_count = 4
        } else if op == 2 {  // Multiply
            opcodes[out] = opcodes[a] * opcodes[b]
            instruction_count = 4
        }
        pointer += instruction_count
    }
    return opcodes[0]
}
