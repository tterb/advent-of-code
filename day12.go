package main

import (
    "bufio"
    "fmt"
    "math"
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

func (m *moon) energy() int {
    pot := int(math.Abs(float64(m.pos.x))+math.Abs(float64(m.pos.y))+math.Abs(float64(m.pos.z)))
    kin := int(math.Abs(float64(m.velo.x))+math.Abs(float64(m.velo.y))+math.Abs(float64(m.velo.z)))
    total := pot * kin
    fmt.Printf("pot: %d + %d + %d = %d;  kin: %d + %d + %d = %d;  total: %d\n", int(math.Abs(float64(m.pos.x))), int(math.Abs(float64(m.pos.y))), int(math.Abs(float64(m.pos.z))), pot, int(math.Abs(float64(m.velo.x))), int(math.Abs(float64(m.velo.y))), int(math.Abs(float64(m.velo.z))), kin, total)
    return total
}

func main() {
    input := []string{}
    scanner := bufio.NewScanner(os.Stdin)
    // Read and process input from stdin
    for scanner.Scan() {
        input = append(input, scanner.Text())
    }
    step := 0
    limit := 1000
    moons := resolveMoons(input)
    for step < limit {
        if step % 100 == 0 {
            printer(moons, step)
        } 
        applyGravity(&moons)
        applyVelocity(&moons)
        step++
    }
    printer(moons, step)
    energy := calcTotalEnergy(&moons)
    fmt.Printf("\nTotal energy: %d\n", energy)
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

func applyVelocity(moons *[]moon) {
    for i, _ := range (*moons) {
        (*moons)[i].applyVelocity()
    }
}

func calcTotalEnergy(moons *[]moon) int {
    total := 0
    fmt.Print("\n")
    for _, m := range (*moons) {
        total += m.energy()
    }
    return total
}

func printer(moons []moon, step int) {
    fmt.Printf("Step #%d\n", step)
    for j, i := range moons {
        fmt.Printf("%d: position: (x: %d, y: %d, z: %d), velocity: (x: %d, y: %d, z: %d)\n", j, i.pos.x, i.pos.y, i.pos.z, i.velo.x, i.velo.y, i.velo.z)
    }
}
