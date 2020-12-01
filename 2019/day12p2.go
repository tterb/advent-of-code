package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type moon struct {
    pos coord
    velo coord
}

type coord struct {
    x int
    y int
    z int
}

func (m *moon) applyVelocity() {
    m.pos.x += m.velo.x
    m.pos.y += m.velo.y
    m.pos.z += m.velo.z
}

func main() {
    input := []string{}
    scanner := bufio.NewScanner(os.Stdin)
    // Read and process input from stdin
    for scanner.Scan() {
        input = append(input, scanner.Text())
    }
    step := 0
    moons := resolveMoons(input)
    period := make(map[string]int)
    positions := make(map[string][]int)
    lcm := -1
    for true {
        applyGravity(&moons)
        applyVelocity(&moons, &positions)
        if step > 1 && step % 100 == 0 {
            lcm = resolveLCM(&positions, &period, len(moons))
            if lcm > 0 {
                break
            }
        }
        step++
    }
    fmt.Printf("\n%d steps until repeated\n", lcm)
}

func resolveMoons(input []string) []moon {
    moons := []moon{}
    for _, line := range input {
        // Remove brackets
        line = line[1:]
        line = line[:len(line)-1]
        vals := []int{}
        for _, val := range strings.Split(line, ", ") {
            // Get position values
            temp := strings.Split(val, "=")
            n, _ := strconv.Atoi(temp[1])
            vals = append(vals, n)
        }
        m := moon{coord{vals[0], vals[1], vals[2]}, coord{0,0,0}}
        moons = append(moons, m)
    }
    return moons
}


func applyGravity(moons *[]moon) {
    visited := make(map[moon]bool)
    for i, _ := range (*moons) {
        m := (*moons)[i]
        visited[m] = true
        for j, _ := range (*moons) {
            temp := (*moons)[j]
            if !visited[temp] {
                if m.pos.x < temp.pos.x {
                    m.velo.x += 1
                } else if m.pos.x > temp.pos.x {
                    m.velo.x -= 1
                }
                if m.pos.y < temp.pos.y {
                    m.velo.y += 1
                } else if m.pos.y > temp.pos.y {
                    m.velo.y -= 1
                }
                if m.pos.z < temp.pos.z {
                    m.velo.z += 1
                } else if m.pos.z > temp.pos.z {
                    m.velo.z -= 1
                }
            }
            (*moons)[j] = temp
        }
        (*moons)[i] = m
    }
}

func applyVelocity(moons *[]moon, positions *map[string][]int) {
    for i, m := range (*moons) {
        (*moons)[i].applyVelocity()
        (*positions)["x"] = append((*positions)["x"], m.pos.x)
        (*positions)["y"] = append((*positions)["y"], m.pos.y)
        (*positions)["z"] = append((*positions)["z"], m.pos.z)
    }
}

func resolveLCM(positions *map[string][]int, period *map[string]int, count int) int {
    for _, i := range [3]string{"x", "y", "z"} {
        if (*period)[i] <= 1 {
            temp := minPeriod((*positions)[i], count)
            if temp <= 1 {
                return -1
            } else {
                (*period)[i] = temp
            }
        }
    }
    // fmt.Printf("Periods: x: %d, y: %d, z: %d\n", (*period)["x"], (*period)["y"], (*period)["z"])
    return lcm((*period)["x"], (*period)["y"], (*period)["z"])
}

func lcm(a, b int, integers ...int) int {
    result := a * b / gcd(a, b)
    for i := 0; i < len(integers); i++ {
        result = LCM(result, integers[i])
    }
    return result
}

func gcd(a, b int) int {
    for b != 0 {
        t := b
        b = a % b
        a = t
    }
    return a
}

func minPeriod(arr []int, count int) int {
    i, j := 0, 1
    period := -1
    streak := false
    for j < len(arr) {
        if arr[i] == arr[j] {
            if !streak {
                period = j
            }
            streak = true
            i++
        } else {
            if streak {
                j--
            }
            i = 0
            streak = false
        }
        j++
    }
    if streak && len(arr) >= period*2 {
        return int(period/count)
    }
    return -1
}
