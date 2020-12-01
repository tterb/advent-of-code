package main

import (
    "bufio"
    "fmt"
    // "io/ioutil"
    "math"
    "os"
    "strconv"
    "strings"
)

type Coord struct {
    x int
    y int
}

type Ball struct {
    pos Coord
    dir Coord
}

func (b *Ball) isFinished(dimension Coord) bool {
    return (b.pos.x < 0 || b.pos.x > dimension.x) || (b.pos.x < 0 || b.pos.x > dimension.x)
}

func (b *Ball) move(board *map[Coord]int) {
    next := Coord{b.pos.x+b.dir.x, b.pos.y+b.dir.x}
    tile := (*board)[next]
    if tile == 0 {         // empty tile
        (*board)[b.pos] = 0
        b.pos = next
        (*board)[next] = 4
    } else if tile == 1 {  // wall tile
        b.dir.y *= -1
    } else if tile == 2 {  // block tile
        (*board)[next] = 0
        b.dir.x *= -1
        b.dir.y *= -1
    } else if tile == 3 {  // paddle tile
        b.dir.x *= -1
        b.dir.y *= -1
    }
}

func main() {
    // Read and process input from file
    file, _ := os.Open("input/input13.txt")
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    memory := []int{}
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
    // for j, i := range memory {
    //     if j < count {
    //         fmt.Printf("%d,", i)
    //     } else {
    //         break
    //     }
    // }
    // fmt.Print("\n")
    
    // Set memory address 0 to 2 to play for free
    memory[0] = 2
    out, output := 0, []int{}
    input := 1
    output = compute(memory, count, input, out)
    
    // Use output to create game board
    board, _, dimension, score := createBoard(output)
    // for !ball.isFinished(dimension) {
    //     // printer(board, dimension)
    //     fmt.Printf("ball: (%d, %d) => ", ball.pos.x, ball.pos.y)
    //     ball.move(&board)
    //     fmt.Printf("(%d, %d)\n", ball.pos.x, ball.pos.y)
    // }
    printer(board, dimension, score)
    // memory[0] = 2
    input = getUserInput()
    output = []int{}
    output = compute(memory, count, input, out)
    
    // Use output to create game board
    board, _, dimension, score = createBoard(output)
    printer(board, dimension, score)
    
    blockCount := getBlocks(board)
    fmt.Printf("blocks: %d\n", blockCount)
}

func createBoard(input []int) (map[Coord]int, Ball, Coord, int) {
    board := make(map[Coord]int)
    max, min, i := Coord{0,0}, Coord{0,0}, 0
    score := 0
    // Get min and max indices
    for i+2 < len(input) {
        if input[i] == -1 && input[i+1] == 0 {
            score = input[i+2]
        } else {
            if max.x < input[i] {
                max.x = input[i]
            } else if max.y< input[i+1] {
                max.y = input[i+1]
            } else if min.x > input[i] {
                min.x = input[i]
            } else if min.y > input[i+1] {
                min.y = input[i+1]
            }
        }
        i += 3
    }
    offset := Coord{int(math.Abs(float64(min.x))), int(math.Abs(float64(min.y)))}
    max.x += offset.x+1
    max.y += offset.y+1
    i = 0
    ball := Ball{Coord{0,0}, Coord{1,1}}
    for i+2 < len(input) {
        temp := Coord{offset.x+input[i], offset.y+input[i+1]}
        board[temp] = input[i+2]
        if input[i+2] == 4 {
            ball.pos = temp
        }
        i += 3
    }
    return board, ball, max, score
}

func getBlocks(board map[Coord]int) int {
    total := 0
    for _, tile := range board {
        if tile == 2 {
            total += 1
        }
    }
    return total
}

func printer(board map[Coord]int, dimension Coord, score int) {
    y := 0
    fmt.Print("\n")
    for y < dimension.y {
        x := 0
        for x < dimension.x {
            tile := board[Coord{x, y}]
            if tile == 0 {
                fmt.Print(" ")
            } else if tile == 1 {
                fmt.Print("|")
            } else if tile == 2 {
                fmt.Print("X")
            } else if tile == 3 {
                fmt.Print("-")
            } else if tile == 4 {
                fmt.Print("o")
            }
            x++
        }
        fmt.Print("\n")
        y++
    }
    fmt.Printf("Score: %d\n", score)
}

func getUserInput() int {
    buffer := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a direction (none: 0, left: -1, right: 1): ")
    input, err := buffer.ReadString('\n')
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    val, _ := strconv.Atoi(input)
    return val
}


func compute(memory []int, instCount, input, out int) []int {
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
            memory[params[0]] = input
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
