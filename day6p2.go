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

func findPathDist(orbits map[string]Node, src string, dest string, dist map[string]int) int {
    queue := []Node{}
    visited := make(map[string]bool)
    depth := 0
    for k, _ := range orbits {
        visited[k] = false
        // max integer value
        dist[k] = int(^uint(0) >> 1)
    }
    queue = append(queue, orbits[src])
    visited[src] = true
    dist[src] = depth
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        for _, n := range node.orbit {
            if !visited[n] {
                visited[n] = true
                dist[n] = dist[node.id]+1
                queue = append(queue, orbits[n])
            }
        }
    }
    return dist[dest] - 2 
}

func main() {
    orbits := make(map[string]Node)
    // Read and process input from stdin
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        vals := strings.Split(scanner.Text(), ")")
        node := orbits[vals[0]]
        n := orbits[vals[1]]
        if _, exists := orbits[vals[0]]; !exists {
            node = Node{id: vals[0], orbit: []string{}, depth: 0}
        }
        if _, exists := orbits[vals[1]]; !exists {
            n = Node{id: vals[1], orbit: []string{}, depth: 0}
        }
        node.addOrbit(vals[1])
        n.addOrbit(vals[0])
        orbits[vals[0]] = node
        orbits[vals[1]] = n
    }
    dist := make(map[string]int)
    pathLength := findPathDist(orbits, "YOU", "SAN", dist)
    fmt.Printf("orbital transfers: %d", pathLength)
}
