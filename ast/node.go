package ast

type Node struct {
	Type  string
	Raw   string
	Left  *Node
	Right *Node
}
