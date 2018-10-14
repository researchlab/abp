package queue

type Node struct {
	data interface{}
	next *Node
}

type Queue struct {
	head *Node
	end  *Node
}

func New() *Queue {
	return &Queue{}
}

//入队列
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

//出队列
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

//返回队首元素值
func (p *Queue) Peek() (interface{}, bool) {
	if p.head == nil {
		return nil, false
	}
	return p.head.data, true
}
