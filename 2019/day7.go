package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    origin := []int{}
    scanner := bufio.NewScanner(os.Stdin)
    // Read and process input from stdin
    for scanner.Scan() {
        for _, val := range strings.Split(scanner.Text(), ",") {
            num, _ := strconv.Atoi(val)
            origin = append(origin, num)
        }
    }
    
    opcodes := make([]int, len(origin))
    copy(opcodes, origin)
    max := 0
    sequences := getPermutations([]int{0,1,2,3,4})
    for _, seq := range sequences {
        out := 0
        for _, phase := range seq {
            // Restore original memory values
            copy(opcodes, origin)
            out = compute(opcodes, phase, out)
            if max < out {
                max = out
            }
        }
    }
    fmt.Printf("diagnostic code: %d\n", max)
}


func compute(opcodes []int, phase, out int) int {
    // Array of instruction parameter counts
    paramCounts := []int{0, 4, 4, 2, 2, 3, 3, 4, 4}
    instrCount := 2
    index, outIndex := 0, 0
    for index < len(opcodes)-1 {
        op := 0
        if opcodes[index] < 0 {
            op -= (opcodes[index]*-1)%10
        } else {
            op = opcodes[index]%10
        }
        // Get instruction parameter count
        instrCount = paramCounts[op]
        // Get parameters by mode
        params := getInstructionParams(opcodes, index, out, instrCount)
        if op == 1 {  // addition instruction
            opcodes[opcodes[index+3]] = params[0] + params[1]
        } else if op == 2 {  // multiplication instruction
            opcodes[opcodes[index+3]] = params[0] * params[1]
        } else if op == 3 {  // input instruction
            // Feed radiator controller ID as input
            var input int
            if outIndex == 0 {
                input = phase
            } else {
                input = out
            }
            if params[0] == 0 {
                opcodes[opcodes[index+1]] = input
            } else {
                opcodes[index+1] = input
            }
            outIndex++
        } else if op == 4 {  // output instruction
            out = opcodes[opcodes[index+1]]
            break
        } else if op == 5 {  // jump-if-true instruction
            if params[0] != 0 {
                // set instruction pointer to second parameter, minus increment
                index = params[1] - instrCount
            }
        } else if op == 6 {  // jump-if-false instruction
            if params[0] == 0 {
                // set instruction pointer to second parameter, minus increment
                index = params[1] - instrCount
            }
        } else if op == 7 {  // less-than instruction
            // store 1 or 0 at the index of the third parameter, based on if 
            // the first parameter is less-than the second
            if params[0] < params[1] {
                opcodes[opcodes[index+3]] = 1
            } else {
                opcodes[opcodes[index+3]] = 0
            }
        } else if op == 8 {  // equals instruction
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
    return out
}


func getInstructionParams(opcodes []int, index, out, instrCount int) []int {
    var digits []int
    digits = intToSlice(opcodes[index], digits)
    // append leading zeroes to instruction
    for len(digits) <= instrCount {
        digits = append([]int{0}, digits...)
    }
    // get param modes from instruction
    modes := digits[:len(digits)-2]
    var params []int
    i := 1
    for i < instrCount && index+i < len(opcodes) {
        // Get param values based on modes
        if modes[len(modes)-i] == 1 {
            params = append(params, opcodes[index+i])
        } else {
            if opcodes[index+i] > len(opcodes)-1 {
                return params
            }
            params = append(params, opcodes[opcodes[index+i]])
        }
        i++
    }
    return params
}

func getPermutations(arr []int) [][]int {
    var helper func([]int, int)
    res := [][]int{}
    helper = func(arr []int, n int){
        if n == 1{
            tmp := make([]int, len(arr))
            copy(tmp, arr)
            res = append(res, tmp)
        } else {
            for i := 0; i < n; i++{
                helper(arr, n - 1)
                if n % 2 == 1{
                    tmp := arr[i]
                    arr[i] = arr[n - 1]
                    arr[n - 1] = tmp
                } else {
                    tmp := arr[0]
                    arr[0] = arr[n - 1]
                    arr[n - 1] = tmp
                }
            }
        }
    }
    helper(arr, len(arr))
    return res
}

func intToSlice(n int, digits []int) []int {
    if n != 0 {
        i := n % 10
        digits = append([]int{i}, digits...)
        return intToSlice(n/10, digits)
    }
    return digits
}
