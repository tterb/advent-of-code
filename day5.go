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
        for _, val := range strings.Split(scanner.Text(), ",") {
            num, _ := strconv.Atoi(val)
            opcodes = append(opcodes, num)
        }
    }
    
    index, output := 0, 0
    conditionerID := 1
    instrCount := 2
    for index < len(opcodes)-1 {
        inst := opcodes[index]
        var digits []int
        // process instruction
        digits = intToSlice(inst, digits)
        // append leading zeroes to instruction
        for len(digits) < 5 {
            digits = append([]int{0}, digits...)
        }
        op := digits[len(digits)-1]
        modes := digits[:len(digits)-2]
        if op <= 2 {
            instrCount = 4
            var params []int
            for i := 1; i < 3; i++ {
                // Get param values based on modes
                if modes[len(modes)-i] == 0 {
                    if len(opcodes) < opcodes[index+i] {
                        fmt.Printf("diagnostic code: %d\n", output)
                        os.Exit(0)
                    }
                    params = append(params, opcodes[opcodes[index+i]])
                } else {
                    params = append(params, opcodes[index+i])
                }
            }
            if op == 1 {  // addition instruction
                opcodes[opcodes[index+3]] = params[0] + params[1]
            } else if op == 2 {  // multiplication instruction
                opcodes[opcodes[index+3]] = params[0] * params[1]
            }
        } else if op == 3 {  // input instruction
            instrCount = 2
            // Feed air conditioner ID as input
            opcodes[opcodes[index+1]] = conditionerID
        } else if op == 4 {  // output instruction
            instrCount = 2
            output = opcodes[opcodes[index+1]]
        }
        index += instrCount
    }
    fmt.Printf("diagnostic code: %d\n", output)
}

func intToSlice(n int, digits []int) []int {
    if n != 0 {
        i := n % 10
        digits = append([]int{i}, digits...)
        return intToSlice(n/10, digits)
    }
    return digits
}
