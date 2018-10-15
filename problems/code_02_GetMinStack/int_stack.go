package code_02_GetMinStack

type IntStack struct {
	stack []int
	len   int
}

func NewStack() *IntStack {
	return &IntStack{
		stack: make([]int, 0),
		len:   0,
	}
}

func (p *IntStack) Push(data int) {
	pre := make([]int, 1)
	pre[0] = data
	p.stack = append(pre, p.stack...)
	p.len++
}

func (p *IntStack) Pop() (el int, ok bool) {
	if p.len == 0 {
		return 0, false
	}
	p.len--
	el, p.stack = p.stack[0], p.stack[1:]

	return el, true
}

func (p *IntStack) Peek() (el int, ok bool) {
	if p.len == 0 {
		return 0, false
	}
	return p.stack[0], true
}
