package stackv1

import "sync"

type Stack struct {
	stack []interface{}
	lock  sync.Mutex
}

func New() *Stack {
	return &Stack{
		stack: make([]interface{}, 0),
	}
}

func (p *Stack) Len() int {
	p.lock.Lock()
	defer p.lock.Unlock()
	return len(p.stack)
}

func (p *Stack) Pop() (interface{}, bool) {
	p.lock.Lock()
	defer p.lock.Unlock()
	if len(p.stack) <= 0 {
		return nil, false
	}
	el := p.stack[0]
	if len(p.stack) == 1 {
		p.stack = make([]interface{}, 0)
	} else {
		p.stack = p.stack[1:]
	}
	return el, true
}

func (p *Stack) Push(el interface{}) {
	p.lock.Lock()
	defer p.lock.Unlock()
	prepend := make([]interface{}, 1)
	prepend[0] = el
	p.stack = append(prepend, p.stack...)
}
