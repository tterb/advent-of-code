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
    controllerID := 5
    instrCount := 2
    for index < len(opcodes)-1 {
        op := opcodes[index]%10
        if opcodes[index]%100 > 9 {
            instrCount = 1
        } else if op == 1 {  // addition instruction
            instrCount = 4
            // process instruction modes and params
            params := getInstructionParams(opcodes, index, output, instrCount)
            opcodes[opcodes[index+3]] = params[0] + params[1]
        } else if op == 2 {  // multiplication instruction
            instrCount = 4
            params := getInstructionParams(opcodes, index, output, instrCount)
            opcodes[opcodes[index+3]] = params[0] * params[1]
        } else if op == 3 {  // input instruction
            instrCount = 2
            // Feed radiator controller ID as input
            params := getInstructionParams(opcodes, index, output, instrCount)
            if params[0] == 0 {
                opcodes[opcodes[index+1]] = controllerID
            } else {
                opcodes[index+1] = controllerID
            }
        } else if op == 4 {  // output instruction
            instrCount = 3
            if opcodes[index] > 9 {
                output = opcodes[index+1]
            } else {
                output = opcodes[opcodes[index+1]]
            }
            fmt.Printf("diagnostic code: %d\n", output)
            os.Exit(0)
        } else if op == 5 {  // jump-if-true instruction
            instrCount = 3
            params := getInstructionParams(opcodes, index, output, instrCount)
            if params[0] != 0 {
                // set instruction pointer to second parameter, minus increment
                index = params[1] - instrCount
            }
        } else if op == 6 {  // jump-if-false instruction
            instrCount = 3
            params := getInstructionParams(opcodes, index, output, instrCount)
            if params[0] == 0 {
                // set instruction pointer to second parameter, minus increment
                index = params[1] - instrCount
            }
        } else if op == 7 {  // less-than instruction
            instrCount = 4
            params := getInstructionParams(opcodes, index, output, instrCount)
            // store 1 or 0 at the index of the third parameter, based on if 
            // the first parameter is less-than the second
            if params[0] < params[1] {
                opcodes[opcodes[index+3]] = 1
            } else {
                opcodes[opcodes[index+3]] = 0
            }
        } else if op == 8 {  // equals instruction
            instrCount = 4
            params := getInstructionParams(opcodes, index, output, instrCount)
            // store 1 or 0 at the index of the third parameter, based on if 
            // the first parameter is equal to the second
            if params[0] == params[1] {
                opcodes[opcodes[index+3]] = 1
            } else {
                opcodes[opcodes[index+3]] = 0
            }
        }
        index += instrCount
    }
    fmt.Printf("diagnostic code: %d\n\n", output)
}

func getInstructionParams(opcodes []int, index, output, count int) []int {
    var digits []int
    digits = intToSlice(opcodes[index], digits)
    // append leading zeroes to instruction
    for len(digits) <= count {
        digits = append([]int{0}, digits...)
    }
    // get param modes from instruction
    modes := digits[:len(digits)-2]
    var params []int
    // for i := 1; i < count; i++ {
    i := 1
    for i < count && index+i < len(opcodes) {
        // Get param values based on modes
        if modes[len(modes)-i] == 1 {
            params = append(params, opcodes[index+i])
        } else {
            if opcodes[index+i] > len(opcodes)-1 {
                fmt.Printf("error - diagnostic code: %d\n", output)
                os.Exit(0)
            }
            params = append(params, opcodes[opcodes[index+i]])
        }
        i++
    }
    return params
}

func intToSlice(n int, digits []int) []int {
    if n != 0 {
        i := n % 10
        digits = append([]int{i}, digits...)
        return intToSlice(n/10, digits)
    }
    return digits
}
