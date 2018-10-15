package code_03_StackQueueConvert

type Stack struct {
	data []interface{}
	len  int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]interface{}, 0),
		len:  0,
	}
}
func (p *Stack) Peek() interface{} {
	if p.len == 0 {
		return nil
	}
	return p.data[0]
}

func (p *Stack) Push(data interface{}) {
	pre := make([]interface{}, 1)
	pre[0] = data
	p.data = append(pre, p.data...)
	p.len++
}

func (p *Stack) Pop() interface{} {
	if p.len == 0 {
		return nil
	}
	var el interface{}
	el, p.data = p.data[0], p.data[1:]
	p.len--
	return el
}

func (p *Stack) IsEmpty() bool {
	return p.len == 0
}

func (p *Stack) Size() int {
	return p.len
}
