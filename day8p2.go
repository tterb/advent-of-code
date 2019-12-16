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
    scanner := bufio.NewScanner(os.Stdin)
    // Read and process input from stdin
    for scanner.Scan() {
        for _, val := range strings.Split(scanner.Text(), "") {
            if len(row) == width {
                layer = append(layer, row)
                row = []int{}
                if len(layer) == height {
                    layers = append(layers, layer)
                    layer = [][]int{}
                }
            }
            num, _ := strconv.Atoi(val)
            row = append(row, num)
				}
		}
		if len(row) == width {
			layer = append(layer, row)
			if len(layer) == height {
				layers = append(layers, layer)
			}
		}
		image := combineLayers(layers)
		printImage(image)
}

func combineLayers(layers [][][]int) [][]int {
	// Combine image layers
	image := [][]int{}
	for i, row := range layers[0] {
		temp := []int{}
		for j, val := range row {
			// Pixel is transparent
			if val == 2 {
				index := 1
				// Find first layer that isn't transparent
				for index < len(layers) {
					if layers[index][i][j] != 2 {
						temp = append(temp, layers[index][i][j])
						break
					}
					index++
				}
			} else {
				temp = append(temp, val)
			}
		}
		image = append(image, temp)
	}
	return image
}


func printImage(image [][]int) {
	for _, row := range image {
		for _, val := range row {
			if val == 1 {
				fmt.Printf("%s", "█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
