package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strings"
)

type coord struct {
    x int
    y int
}

func main() {
    mapp := [][]bool{}
    asteroids := make(map[coord]bool)
    // sightCount := make(map[coord]int)
    scanner := bufio.NewScanner(os.Stdin)
    // Read and process input from stdin
    for scanner.Scan() {
        row := []bool{}
        for _, val := range strings.Split(scanner.Text(), "") {
            if val == "#" {
                // temp := coord{len(mapp), len(row)}
                // asteroids = append(asteroids, coord{len(row), len(mapp)})
                asteroids[coord{len(row), len(mapp)}] = true
            }
            row = append(row, (val == "#"))
        }
        mapp = append(mapp, row)
    }
    printMap(mapp)
    
    // Resolve line-of-sights
    sights := make(map[coord][]coord)
    for node, val := range asteroids {
        if val {
            sights[node] = resolveSight(asteroids, node, len(mapp))
        }
    }

    max := 0
    station := coord{-1, -1}
    // Get max line-of-sight
    for node, _ := range sights {
        // fmt.Printf("(%d, %d): ", node.x, node.y)
        // for _, i := range sights[node] {
            // fmt.Printf("(%d, %d), ", i.x, i.y)
        // }
        // fmt.Print("\n")
        if max < len(sights[node]) {
            max = len(sights[node])
            station = node
        }
    }
    // printer(asteroids, sightCount, len(mapp), len(mapp[0]))
    printer(asteroids, sights, len(mapp), len(mapp[0]))
    fmt.Printf("station: (%d, %d) = %d", station.x, station.y, max)
}

func resolveSight(asteroids map[coord]bool, node coord, lim int) []coord {
    sight := make(map[coord]bool)
    for k, v := range asteroids {
        if k != node {
            sight[k] = v
        }
    }
    for n, val := range sight {
        intersect := between(node, n)
        if val && asteroids[intersect] {
            sight[n] = false
            fmt.Printf("(%d, %d) =/=> (%d, %d)\n", node.x, node.y, n.x, n.y)
            for _, i := range obscured(n, intersect, lim) {
                if sight[i] && i != intersect {
                    fmt.Printf("(%d, %d) =/=> (%d, %d)\n", node.x, node.y, i.x, i.y)
                    sight[i] = false
                }
            }
        }
    }
    // Convert map[coord]bool to slice
    slice := []coord{}
    for n, val := range sight {
        if val {
            slice = append(slice, n)
        }
    }
    // fmt.Printf("(%d, %d): ", node.x, node.y)
    // for _, i := range slice {
        // fmt.Printf("(%d, %d), ", i.x, i.y)
    // }
    // fmt.Print("\n")
    return slice
}

func between(a coord, b coord) coord {
    x, y := (a.x + b.x)/2, (a.y + b.y)/2
    temp := coord{x, y}
    eq := math.Hypot((float64(a.x)-float64(x)), (float64(a.y)-float64(y))) == math.Hypot((float64(b.x)-float64(x)), (float64(b.y)-float64(y)))
    // above, left := (b.y < a.y), (b.x < a.x)
    // sx, sy := 0, 0
    // if above {
    //     sy := -(math.Abs(a.y - b.y))
    // } else {
    //     sy := math.Abs(a.y - b.y)
    // }
    // if left {
    //     sx := -(math.Abs(a.x - b.x))
    // } else {
    //     sx := math.Abs(a.x - b.x)
    // }
    // // slope := coord{sx, sy}
    // x, y := a.x+sx, a.y+sy
    // for (x >= 0 && x < lim-1) || (y >= 0 && y < lim-1) {
    //     if 
    //     x += sx
    //     y += sy
    // }
    if eq {
        // fmt.Printf("(%d, %d), (%d, %d) => (%d, %d)\n", a.x, a.y, b.x, b.y, x, y)
        return temp
    }
    return coord{-1, -1}
    // return temp
}

// func obscured(asteroids map[coord]bool, a, b coord, lim int) []coord {
//     arr := []coord
//     above, left := (b.y < a.y), (b.x < a.x)
//     sx, sy := math.Abs(a.x - b.x), math.Abs(a.y - b.y)
//     if above {
//         sy *= -1
//     }
//     if left {
//         sx *= -1
//     }
//     x, y := a.x+sx, a.y+sy
//     count := 0
//     for (x >= 0 && x < lim-1) || (y >= 0 && y < lim-1) {
//         temp := coord{x, y}
//         if asteroids[temp] {
//             if count != 0 {
//                 arr = append(arr, temp)
//             } else {
//                 count++
//             }
//         }
//         x += sx
//         y += sy
//     }
//     return arr
// }

func obscured(a, b coord, lim int) []coord {
    arr := []coord{}
    s1 := coord{(a.x - b.x), (a.y - b.y)}
    // fmt.Printf("(%d, %d), (%d, %d) = %d/%d\n", a.x, a.y, b.x, b.y, s1.x, s1.y)
    x, y := b.x, b.y
    for x >= 0 && x < lim && x != a.x+s1.x && y >= 0 && y < lim && y != a.y+s1.y {
        x += s1.x
        y += s1.y
        arr = append(arr, coord{x, y})
    }
    s2 := coord{(b.x - a.x), (b.y - a.y)}
    x, y = a.x, a.y
    for x >= 0 && x < lim && x != b.x+s2.x && y >= 0 && y < lim && y != b.y+s2.y {
        x += s2.x
        y += s2.y
        arr = append(arr, coord{x, y})
    }
    return arr
}

func printer(asteroids map[coord]bool, sight map[coord][]coord, rows, cols int) {
    i, j := 0, 0
    c := coord{-1, -1}
    // fmt.Printf("%d, %d\n", rows, cols)
    fmt.Print("\n")
    for i < rows {
        for j < cols {
            c = coord{j, i}
            if len(sight[c]) == 0 {
                fmt.Print(". ")
            } else {
                fmt.Printf("%d ", len(sight[c]))
            }
            j++
        }
        j = 0
        fmt.Print("\n")
        i++
    }
    fmt.Print("\n")
}

func printMap(mapp [][]bool) {
    for _, row := range mapp {
        for _, i := range row {
            if i {
                fmt.Print("# ")
            } else {
                fmt.Print(". ")
            }
        }
        fmt.Print("\n")
     }
     fmt.Print("\n")
}
