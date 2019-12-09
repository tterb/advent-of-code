package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Node struct {
    id string
    orbit []string
    depth int
}

func (node *Node) addOrbit(id string) []string {
    node.orbit = append(node.orbit, id)
    return node.orbit
}

func main() {
    orbits := make(map[string]Node)
    // Read and process input from stdin
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        vals := strings.Split(scanner.Text(), ")")
        node := orbits[vals[0]]
        if _, exists := orbits[vals[0]]; !exists {
            node = Node{id: vals[0], orbit: []string{}, depth: 0}
        }
        if _, exists := orbits[vals[1]]; !exists {
            orbits[vals[1]] = Node{id: vals[1], orbit: []string{}, depth: 0}
        }
        node.addOrbit(vals[1])
        orbits[vals[0]] = node
    }
    depths := traverseOrbit(orbits, "COM", []int{}, 0)
    count := 0
    // Add collected depths
    for _, i := range depths {
        count += i
    }
    fmt.Printf("orbit count: %d", count)
}

func traverseOrbit(orbits map[string]Node, id string, depths []int, depth int) []int {
    node := orbits[id]
    node.depth = depth
    for _, n := range node.orbit {
        depths = traverseOrbit(orbits, n, depths, depth+1)
    }
    return append(depths, depth)
}
