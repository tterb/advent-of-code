package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

// Elected to use a BST for improved lookup efficiency

type Node struct {
    data   int
    left   *Node
    right  *Node
    parent *Node
}

type Tree struct {
    root *Node
}

func newNode(data int) *Node {
	var n Node
	n.data = data
	return &n
}

func (t *Tree) insert(data int) {
	node := newNode(data)
	if t.root == nil {
		t.root = node
		return
	}
	// find position
	var pos *Node = nil
	temp := t.root
	for temp != nil {
	    pos = temp
	    if node.data < temp.data {
		    temp = temp.left
	    } else {
		    temp = temp.right
	    }
	}
	// insert node
	if node.data < pos.data {
		pos.left = node
	} else {
		pos.right = node
	}
	node.parent = pos
}

func (t *Tree) contains(data int) bool {
	return t.root.elem(data)
}

func (n *Node) elem(data int) bool {
	if n == nil {
		return false
	}
	if data == n.data {
		return true
	} else if data < n.data {
		return n.left.elem(data)
	}
	return n.right.elem(data)
}


func main() {
  var arr []string
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
      arr = append(arr, scanner.Text())
  }
  total := 0
	found := false
  tree := Tree{nil}
	for found != true {
	  for _, i := range arr {
	    val, _ := strconv.Atoi(i)
	    total += val
	    if tree.contains(total) {
	      fmt.Printf("%d ", total)
				found = true
				break
	    }
	    tree.insert(total)
	  }
	}
  // fmt.Printf("%d ", total)
}