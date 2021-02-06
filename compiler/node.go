package compiler

//type Node struct {
//	Type     string
//	Raw      string
//	Left     *Node
//	Right    *Node
//	Children []*Node
//}

type Node interface {
	Label() string
}

type Number struct {
	token *Token
}

func NewNumber(token *Token) *Number {
	return &Number{token: token}
}

type BinOpNode struct {
	left  *Number
	op    *Token
	right *Number
}
