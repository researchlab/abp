package stack

type Node struct {
	data interface{}
	next *Node
}

type Stack struct {
	head *Node
}

func New() *Stack {
	return &Stack{head: &Node{data: nil, next: nil}}
}

func (p *Stack) Push(data interface{}) {
	n := &Node{data: data, next: p.head}
	p.head = n
}

func (p *Stack) Pop() (interface{}, bool) {
	n := p.head
	if p.head == nil {
		return nil, false
	}
	p.head = p.head.next
	return n.data, true
}

//返回栈顶元素
func (p *Stack) Peek() (interface{}, bool) {
	if p.head == nil {
		return nil, false
	}
	return p.head.data, true
}
