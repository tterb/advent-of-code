package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type bot struct {
    x int
    y int
    direction direct
}

type direct struct {
    x int
    y int
}

func main() {
    memory := []int{}
    dimension := 0
    scanner := bufio.NewScanner(os.Stdin)
    // Read and process input from stdin
    for scanner.Scan() {
        for _, val := range strings.Split(scanner.Text(), ",") {
            num, _ := strconv.Atoi(val)
            if num%10 == 4 {
                dimension++
            }
            memory = append(memory, num)
        }
    }
    dimension *= 3
    // Add buffer space to memory
    count := len(memory)
    for len(memory) < 50000 {
        memory = append(memory, 0)
    }
    
    panels := createPanels(dimension)
    robot := bot{len(panels)/2, len(panels)/2, direct{0, -1}}
    
    // Robot starts on a white panel
    panels[robot.y][robot.x] = true
    // Get number of painted panels
    painted := compute(memory, count, robot, panels)
    
    pCount := 0
    for _, val := range painted {
        if val {
            pCount++
        }
    }
    printer(panels, robot)
    fmt.Printf("%d panels painted\n", pCount)
}

func createPanels(dimension int) [][]bool {
    i, panels := 0, [][]bool{}
    for i < dimension {
        j, row := 0, []bool{}
        for j < dimension {
            row = append(row, false)
            j++
        }
        panels = append(panels, row)
        i++
    }
    return panels
}

func printer(panels [][]bool, robot bot) {
    fmt.Print("\n")
    for y, row := range panels {
        for x, j := range row {
            if robot.y == y && robot.x == x {
                if robot.direction.x == 0 && robot.direction.y == -1 {
                    fmt.Print("^")
                } else if robot.direction.x == -1 && robot.direction.y == 0 {
                    fmt.Print("<")
                } else if robot.direction.x == 1 && robot.direction.y == 0 {
                    fmt.Print(">")
                } else {
                    fmt.Print("v")
                }
            } else if j == true {
                fmt.Print("#")
            } else {
                fmt.Print(".")
            }
        }
        fmt.Print("\n")
    }
    fmt.Print("\n")
}

func paint(panels *[][]bool, robot bot, val int) {
    if val == 0 {
        (*panels)[robot.y][robot.x] = false
    } else if val == 1 {
        (*panels)[robot.y][robot.x] = true
    }
}

func turn(robot bot, val int) bot {
    directions := make(map[direct][]direct)
    up, left, down, right := direct{0, -1}, direct{-1, 0}, direct{0, 1}, direct{1, 0}
    
    directions[up] = []direct{left, right}
    directions[left] = []direct{down, up}
    directions[down] = []direct{right, left}
    directions[right] = []direct{up, down}
    
    robot.direction = directions[robot.direction][val]
    robot.x += robot.direction.x
    robot.y += robot.direction.y
    return robot
}

/*******************************************************
/* Intcode computer
********************************************************/

func compute(memory []int, instCount int, robot bot, panels [][]bool) map[direct]bool {
    // Array of instruction parameter counts
    paramCounts := []int{2, 4, 4, 2, 2, 3, 3, 4, 4, 2}
    painted := make(map[direct]bool)
    index, relBase := 0, 0
    output := []int{}
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
        params := resolveParams(memory, digits, index, relBase, paramCount)
        if op == 1 {         // addition instruction
            memory[params[2]] = params[0] + params[1]
        } else if op == 2 {  // multiplication instruction
            memory[params[2]] = params[0] * params[1]
        } else if op == 3 {  // input instruction
            inputVal := 1
            if panels[robot.y][robot.x] == false {
                inputVal = 0
            }
            memory[params[0]] = inputVal
        } else if op == 4 {  // output instruction
            temp := params[0]
            output = append(output, temp)
            if len(output) % 2 == 1 {
                painted[direct{robot.x, robot.y}] = true
                paint(&panels, robot, temp)
            } else {
                robot = turn(robot, temp)
            }
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
    return painted
}


func resolveParams(memory []int, digits []int, pointer int, relBase int, paramCount int) []int {
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
