package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    row := []int{}
    layer := [][]int{}
    layers := [][][]int{}
    width, height := 25, 6
    zeroes, ones, twos := 0, 0, 0
    min, res := 151, 0
    i, j := 0, 0
    scanner := bufio.NewScanner(os.Stdin)
    // Read and process input from stdin
    for scanner.Scan() {
        for _, val := range strings.Split(scanner.Text(), "") {
            if len(row) == width {
                layer = append(layer, row)
                row = []int{}
                i++
                if len(layer) == height {
                    layers = append(layers, layer)
                    layer = [][]int{}
                    if zeroes < min {
                        min = zeroes
                        res = ones*twos
                    }
                    zeroes, ones, twos = 0, 0, 0
                    j++
                }
            }
            num, _ := strconv.Atoi(val)
            // Count occurrences
            if num == 0 {
                zeroes++
            } else if num == 1 {
                ones++
            } else if num == 2 {
                twos++
            }
            row = append(row, num)
        }
    }
    fmt.Printf("output: %d\n", res)
}