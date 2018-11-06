package cache

import "container/list"

const (
	DEFAULT_LRU_CACHE_SIZE = 10
)

//lru
//push front
//update and get trigger move to front
//delete the back elem
type LRUInterface interface {
	//set (k,v) to cache
	//update v if k exists
	//insert v if k not exists and check list length
	//delete the back element if list length > cache capacity
	Set(k, v interface{})
	//return (k)'s value and true from cache if k exists
	Get(k interface{}) (interface{}, bool)
	//return cache size
	Size() int
	//purge cache
	Purge()
}

type Node struct {
	Key, Value interface{}
}

type LRUCache struct {
	Capacity int
	dlist    *list.List
	cacheMap map[interface{}]*list.Element
}

func NewLRUCache(size int) *LRUCache {
	if size < 1 {
		size = DEFAULT_LRU_CACHE_SIZE
	}
	return &LRUCache{
		Capacity: size,
		dlist:    list.New(),
		cacheMap: make(map[interface{}]*list.Element),
	}
}

func (p *LRUCache) Set(k, v interface{}) {
	// update v if k exists
	//move k to front
	if elem, ok := p.cacheMap[k]; ok {
		elem.Value.(*Node).Value = v
		p.dlist.MoveToFront(elem)
		return
	}
	elem := p.dlist.PushFront(&Node{k, v})
	p.cacheMap[k] = elem

	if p.dlist.Len() > p.Capacity {
		lastElem := p.dlist.Back()
		if lastElem == nil {
			return
		}
		node := lastElem.Value.(*Node)
		delete(p.cacheMap, node.Key)
		p.dlist.Remove(lastElem)
	}
	return
}

func (p *LRUCache) Get(k interface{}) (interface{}, bool) {
	if p.cacheMap == nil {
		return nil, false
	}
	if elem, ok := p.cacheMap[k]; ok {
		p.dlist.MoveToFront(elem)
		return elem.Value.(*Node).Value, true
	}
	return nil, false
}

func (p *LRUCache) Size() int {
	return p.dlist.Len()
}

func (p *LRUCache) Purge() {
	p.cacheMap = make(map[interface{}]*list.Element)
	p.dlist.Init()
}
