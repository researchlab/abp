package code_03_StackQueueConvert

//队列 只能从一端插入，另一端删除的线性结构
//从队首删除，队尾入数据
type Queue struct {
	data []interface{}
	len  int
}

func New() *Queue {
	return &Queue{
		data: make([]interface{}, 0),
		len:  0,
	}
}

//获取队列头部元素值
func (p *Queue) Peek() interface{} {
	if p.len == 0 {
		return nil
	}
	return p.data[0]
}

func (p *Queue) Push(data interface{}) {
	p.data = append(p.data, data)
	p.len++
}

//出队列 队首弹出元素
func (p *Queue) Poll() (el interface{}) {
	if p.len == 0 {
		return nil
	}
	el, p.data = p.data[0], p.data[1:]
	p.len--
	return el
}

func (p *Queue) IsEmpty() bool {
	if p.len == 0 {
		return true
	}
	return false
}

func (p *Queue) Size() int {
	return p.len
}
