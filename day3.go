package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type coord struct {
    x int
    y int
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    distance := make(map[coord]int)
    // Max integer value
    minDist := int(^uint(0) >> 1)
    wireIndex := 1
    // Read and process input from stdin
    for scanner.Scan() {
        wire := strings.Split(scanner.Text(), ",")
        // Starting position
        pos := coord{0, 0}
        for _, p := range wire {
            incX, incY := getDirection(p[0:1])
            amt, _ := strconv.Atoi(p[1:])
            // Iterate path and check for intersections
            for j := 0; j < amt; j++ {
                pos = coord{pos.x+incX, pos.y+incY}
                dist := distance[pos]
                if wireIndex > 1 && dist != 0 {
                    if dist < minDist {
                        minDist = dist
                    }
                } else if wireIndex == 1 {
                    distance[pos] = abs(pos.x) + abs(pos.y)
                }
            }
        }
        wireIndex += 1
    }
    fmt.Printf("%d ", minDist)
}

func getDirection(c string) (int, int) {
    if c == "U" {
        return 0, -1
    } else if c == "D" {
        return 0, 1
    } else if c == "R" {
        return 1, 0
    } else {
       return -1, 0
   }
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
