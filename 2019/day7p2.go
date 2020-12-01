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
    results := getOutputSignal(origin, getPermutations([]int{0,1,2,3,4}), 0, 0)
    // var outs []int 
    // for _, out := range results {
    //     outs = getOutputSignal(origin, []int{5,6,7,8,9}, out)
    // }
    max := 0
    for _, i := range results {
        if max < i {
            max = i
        }
    }
    // max := 0
    // results := []int{}
    // sequences := getPermutations([]int{0,1,2,3,4})
    // for _, seq := range sequences {
    //     // Restore original memory values
    //     // seq := []int{0,1,2,3,4}
    //     // copy(opcodes, origin)
    //     fmt.Printf("\nphase: ")
    //     printer(seq)
    //     printer(opcodes)
    //     out := 0
    //     for _, phase := range seq {
    //         copy(opcodes, origin)
    //         out = compute(opcodes, phase, out)
    //         fmt.Printf("out: %d\n", out)
    //         fmt.Print("opcodes: ")
    //         printer(opcodes)
    //     }
    //     results = append(results, out)
    // }
    // copy(opcodes, origin)
    // sequences = getPermutations([]int{5,6,7,8,9})
    // for _, res := range results {
    //     for _, seq := range sequences {
    //         // Restore original memory values
    //         // seq := []int{0,1,2,3,4}
    //         // copy(opcodes, origin)
    //         // fmt.Printf("\nphase: ")
    //         // printer(seq)
    //         // printer(opcodes)
    //         out := res
    //         for _, phase := range seq {
    //             copy(opcodes, origin)
    //             out = compute(opcodes, phase, out)
    //             fmt.Printf("out: %d\n", out)
    //             // fmt.Print("opcodes: ")
    //             // printer(opcodes)
    //             if max < out {
    //               max = out
    //             }
    //         }
    //     }
    // }
    fmt.Print("\n")
    fmt.Printf("diagnostic code: %d\n", max)
}


func printer(opcodes []int) {
    for _, i := range opcodes {
      fmt.Printf("%d, ", i)
    }
    fmt.Printf("\n")
}

func getOutputSignal(origin []int, sequences [][]int, input, depth int) []int {
    // max := 0
    results := []int{}
    opcodes := make([]int, len(origin))
    copy(opcodes, origin)
    for _, seq := range sequences {
        out := input
        for _, phase := range seq {
            // Restore original memory values
            copy(opcodes, origin)
            out = compute(opcodes, phase, out)
        }
        // results = append(results, out)
        if depth == 0 {
            return getOutputSignal(opcodes, getPermutations([]int{5,6,7,8,9}), out, depth+1)
        } else {
            results = append(results, out)
        }
    }
    return results
}

func compute(opcodes []int, phase, out int) int {
    fmt.Printf("\n****************phase: %d****************\n", phase)
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
        // if opcodes[index]%100 > 9 || op == 0 {
        //     instrCount = 1
        //     fmt.Printf("Error: index %d\n", index)
        // } else {
        instrCount = paramCounts[op]
        // }
        // if op == 4 {  // output instruction
        //     out = opcodes[opcodes[index+1]]
        //     fmt.Printf("******out: %d******\n", out)
        //     index += instrCount
        //     continue
        // }
        // instrCount = paramCounts[op]
        params, err := getInstructionParams(opcodes, index, out, instrCount)
        if err != nil && op != 4 {
            return out
        }
        if op == 1 {  // addition instruction
            // process instruction modes and params
            // if index+instrCount > len(opcodes)-1 {
            //     return out
            // }
            opcodes[opcodes[index+3]] = params[0] + params[1]
            fmt.Printf(" - %d + %d => %d\n", params[0], params[1], opcodes[index+3])
        } else if op == 2 {  // multiplication instruction
            // if index+instrCount > len(opcodes)-1 {
            //     return out
            // }
            // fmt.Printf(" - %d * %d => %d\n", params[0], params[1], opcodes[opcodes[index+3]])
            opcodes[opcodes[index+3]] = params[0] * params[1]
            fmt.Printf(" - %d * %d => %d\n", params[0], params[1], opcodes[index+3])
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
            fmt.Print("\n")
            outIndex++
        } else if op == 4 {  // output instruction
            fmt.Printf("(%d) digits: %d\n", index, opcodes[index])
            out = opcodes[opcodes[index+1]]
            fmt.Printf("******out: %d******\n", out)
            fmt.Print("\n")
            break
            return out
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
        fmt.Printf("opcodes: ")
        printer(opcodes)
    }
    // fmt.Printf("\n")
    return out
}


func getInstructionParams(opcodes []int, index, out, instrCount int) ([]int, error) {
    var digits []int
    digits = intToSlice(opcodes[index], digits)
    // append leading zeroes to instruction
    for len(digits) <= instrCount {
        digits = append([]int{0}, digits...)
    }
    fmt.Printf("(%d) digits: ", index)
    for _, i := range digits {
        fmt.Printf("%d", i)
    }
    // fmt.Print("\n")
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
                return params, fmt.Errorf("Index out of range")
            }
            // fmt.Printf("%d\n", opcodes[index+i])
            // if opcodes[index+i] < 0 {
                // params = append(params, opcodes[len(opcodes)-1+opcodes[index+i]])
            // } else {
            //     params = append(params, opcodes[opcodes[index+i]])
            // }
            params = append(params, opcodes[opcodes[index+i]])
        }
        i++
    }
    // fmt.Printf("(%d) params: ", instrCount)
    // printer(params)
    // if len(params) < instrCount-1 {
    //     return params, fmt.Errorf("Index out of range")
    // }
    // fmt.Printf("params: ")
    // printer(params)
    return params, nil
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
