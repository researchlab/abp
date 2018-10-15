package code_03_StackQueueConvert

//两个队列实现一个栈
type TwoQueuesStack struct {
	data *Queue
	help *Queue
}

func NewTQS() *TwoQueuesStack {
	return &TwoQueuesStack{
		data: New(),
		help: New(),
	}
}

//直接进入data queue
func (p *TwoQueuesStack) Push(data interface{}) {
	p.data.Push(data)
}

func (p *TwoQueuesStack) Peek() interface{} {
	if p.data.IsEmpty() {
		return nil
	}
	for p.data.Size() != 1 {
		p.help.Push(p.data.Poll())
	}
	res := p.data.Poll()
	p.help.Push(res)
	p.data, p.help = p.help, p.data
	return res
}

func (p *TwoQueuesStack) Pop() interface{} {
	if p.data.IsEmpty() {
		return nil
	}
	for p.data.Size() != 1 {
		p.help.Push(p.data.Poll())
	}
	res := p.data.Poll()
	p.data, p.help = p.help, p.data
	return res
}

//两个栈实现一个队列
//当pop栈不为空时 push栈不动 直接对pop 栈进行pop 和peek操作 这样就不需要交换pop和Push栈了
type TwoStacksQueue struct {
	push *Stack
	pop  *Stack
}

func NewTSQ() *TwoStacksQueue {
	return &TwoStacksQueue{
		push: NewStack(),
		pop:  NewStack(),
	}
}
func (p *TwoStacksQueue) Peek() interface{} {
	if p.push.IsEmpty() && p.pop.IsEmpty() {
		return nil
	}
	if p.pop.IsEmpty() {
		for !p.push.IsEmpty() {
			p.pop.Push(p.push.Pop())
		}
	}
	return p.pop.Peek()
}

func (p *TwoStacksQueue) Push(data interface{}) {
	p.push.Push(data)
}

func (p *TwoStacksQueue) Poll() interface{} {
	if p.pop.IsEmpty() && p.push.IsEmpty() {
		return nil
	}
	if p.pop.IsEmpty() {
		for !p.push.IsEmpty() {
			p.pop.Push(p.push.Pop())
		}
	}
	return p.pop.Pop()
}
