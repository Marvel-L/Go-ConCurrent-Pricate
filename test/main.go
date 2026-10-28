package main

import "fmt"

type Node struct {
	Val int
}

func changePointer(n *Node) {
	n = &Node{Val: 100}
}

func main() {
	x := Node{Val: 1}
	changePointer(&x)
	fmt.Println(x.Val) // 1
}
