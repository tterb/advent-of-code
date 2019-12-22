package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    memory := []int{}
    scanner := bufio.NewScanner(os.Stdin)
    // Read and process input from stdin
    for scanner.Scan() {
        for _, val := range strings.Split(scanner.Text(), ",") {
            num, _ := strconv.Atoi(val)
            memory = append(memory, num)
        }
    }
    count := len(memory)
    for len(memory) < 50000 {
        memory = append(memory, 0)
    }
    out, output := 0, []int{}
    // Test-mode = 1
    inputVal := 1
    output = compute(memory, count, inputVal, out)
    fmt.Printf("BOOST keycode: %d\n", output[0])
}


func compute(memory []int, instCount, inputVal, out int) []int {
    var output []int
    // Array of instruction parameter counts
    paramCounts := []int{2, 4, 4, 2, 2, 3, 3, 4, 4, 2}
    index, relBase := 0, 0
    for index < instCount-2 {
        op := 0
        digits := intToSlice(memory[index], []int{})
        // Get operation from instruction
        if (len(digits) > 1) {
            op = (digits[len(digits)-2]*10) + (digits[len(digits)-1])
        } else {
            op = digits[len(digits)-1]
        }
        // Exit on 99
        if op == 99 {
            break
        }
        
        // Get instruction parameter count
        paramCount := paramCounts[op]
        // Get parameters by mode
        params := resolveParams(memory, digits, index, out, relBase, paramCount)
        if op == 1 {         // addition instruction
            memory[params[2]] = params[0] + params[1]
        } else if op == 2 {  // multiplication instruction
            memory[params[2]] = params[0] * params[1]
        } else if op == 3 {  // input instruction
            memory[params[0]] = inputVal
        } else if op == 4 {  // output instruction
            output = append(output, params[0])
        } else if op == 5 {  // jump-if-true instruction
            if params[0] != 0 {
                index = params[1] - paramCount
            }
        } else if op == 6 {  // jump-if-false instruction
            if params[0] == 0 {
                index = params[1] - paramCount
            }
        } else if op == 7 {  // less-than instruction
            if params[0] < params[1] {
                memory[params[2]] = 1
            } else {
                memory[params[2]] = 0
            }
        } else if op == 8 {  // equals instruction
            if params[0] == params[1] {
                memory[params[2]] = 1
            } else {
                memory[params[2]] = 0
            }
        } else if op == 9 {  // increment relative base
            relBase += params[0]
		}
        index += paramCount
    }
    return output
}


func resolveParams(memory, digits []int, pointer, output, relBase, paramCount int) []int {
    writeParam := []int{-1, 3, 3, 1, -1, -1, -1, 3, 3, -1}
    op := digits[len(digits)-1]
    // append leading zeroes to instruction
    for len(digits) <= paramCount {
        digits = append([]int{0}, digits...)
    }
    // get param modes from instruction
    modes := digits[:len(digits)-2]
    var params []int
    i := 1
    for i < paramCount {
        mode := modes[len(modes)-i]
        // Get param values based on parameter modes
        if mode == 0 {         // position mode
            if i == writeParam[op] {
                params = append(params, memory[pointer+i])
            } else {
                params = append(params, memory[memory[pointer+i]])
            }
        } else if mode == 1 {  // immediate mode
			params = append(params, memory[pointer+i])
        } else if mode == 2 {  // relative mode
            if i == writeParam[op] {
                params = append(params, relBase+memory[pointer+i])
            } else {
                params = append(params, memory[relBase+memory[pointer+i]])
            }
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
