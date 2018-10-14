package queue

type Node struct {
	data interface{}
	next *Node
}

type Queue struct {
	head *Node
	end  *Node
}

func NewQueue() *Queue {
	return &Queue{nil, nil}
}

func (p *Queue) Push(data interface{}) {
	n := &Node{data: data, next: nil}

	if p.end == nil {
		p.head = n
		p.end = n
		return
	}
	p.end.next = n
	p.end = n
}

func (p *Queue) Pop() (interface{}, bool) {
	if p.head == nil {
		return nil, false
	}
	data := p.head.data
	p.head = p.head.next
	if p.head == nil {
		p.end = nil
	}
	return data, true
}
