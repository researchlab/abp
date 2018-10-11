package quicksort

type Node struct {
	Data *Point
	Next *Node
}

type Point struct {
	Less int
	More int
}
type Stack struct {
	head *Node
}

func NewStack() *Stack {
	return &Stack{}
}

func (p *Stack) Push(po *Point) {
	p.head = &Node{
		Data: po,
		Next: p.head,
	}
}

func (p *Stack) Pop() *Point {
	if p.head == nil {
		return nil
	}
	po := p.head.Data
	p.head = p.head.Next
	return po
}

func (p *Stack) IsEmpty() bool {
	if p.head == nil {
		return true
	}
	return false
}
