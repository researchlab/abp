package heap

type Heap struct {
	Size  int
	Elems []int
}

func NewHeap(MaxSize int) *Heap {
	return &Heap{
		Size:  -1,
		Elems: make([]int, MaxSize, MaxSize),
	}
}

func (p *Heap) Push(data int) {
	p.Size++
	i := p.Size // data数据待存放的下标
	// 默认新插入的data 是存放在最末尾的
	// 最小堆调整时, 就需要和起父节点比较
	for i > 0 {
		parent := (i - 1) / 2
		// 如果父节点值小于要插入的节点值， 则可以退出
		if p.Elems[parent] <= data {
			break
		}
		// 否则交换位置, 先将父节点值存储到当前data位置， 然后更新当前data应该存储的索引位置为父节点位置
		// 然后再次比较
		p.Elems[i] = p.Elems[parent]
		i = parent
	}
	p.Elems[i] = data
}

func (p *Heap) Pop() int {
	if p.Size < 0 {
		return -1
	}
	ret := p.Elems[0]
	// 将最后一个节点的值调换到堆顶，然后进行堆调整
	// 记录堆最后一个元素
	node := p.Elems[p.Size]
	i := 0 // 最后一个元素初始位置应该时堆顶，因为堆顶原生刚才pop 出去了，ret
	for {
		// 作为堆顶元素， 要调整就是需要和左右节点比较
		l := 2*i + 1
		r := 2*i + 2
		// 如果左子树大于堆大小，表示当前节点i 已经是叶子节点了, 因为它没有左子树
		if l >= p.Size {
			break
		}

		// 寻找左右最小节点
		if r < p.Size && p.Elems[r] <= p.Elems[l] {
			l = r
		}
		// 比较父节点与最小子节点
		if node <= p.Elems[l] {
			break
		}
		// 父节点不是最小 则要调整更新
		// 此时只需要将最小节点复制给父节点索引位置即可
		p.Elems[i] = p.Elems[l]
		// 更新父节点存储的索引值
		i = l
	}
	// pop之后 size 减少1
	p.Size--
	// 将最后一个元素存放到调整的合适位置i
	p.Elems[i] = node
	return ret
}
