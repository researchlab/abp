package code_02_GetMinStack

type MyStack struct {
	data *IntStack
	min  *IntStack
}

func New() *MyStack {
	return &MyStack{
		data: NewStack(),
		min:  NewStack(),
	}
}
func (p *MyStack) Push(data int) {
	if v, ok := p.min.Peek(); !ok || (ok && v > data) {
		p.min.Push(data)
	} else {
		p.min.Push(v)
	}
	p.data.Push(data)
}

func (p *MyStack) Pop() (int, bool) {
	p.min.Pop()
	return p.data.Pop()
}

func (p *MyStack) Peek() (int, bool) {
	return p.data.Peek()
}

func (p *MyStack) GetMin() (int, bool) {
	return p.min.Peek()
}
