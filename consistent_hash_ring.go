package ConsistentHashRing

import (
	"fmt"
	"github.com/emirpasic/gods/maps/treemap"
)

type Node interface {
	Name() string
	Index() int
}

// 计算哈稀算法
type HashFunc func(string) int

type ConsistentHash struct {
	circle       *treemap.Map
	hash         HashFunc
	virtualCount int // 虚拟节点的数量
}

func New(hashFunc HashFunc) *ConsistentHash {
	return &ConsistentHash{circle: treemap.NewWithIntComparator(), hash: hashFunc}
}

func (r *ConsistentHash) Add(node Node) bool {
	// 添加实体节点
	r.circle.Put(r.hash(node.Name()), node)

	// 初始化虚拟节点
	for i := 0; i < r.virtualCount; i++ {
		r.circle.Put(r.hash(fmt.Sprint(node.Name(), i)), node)
	}

	return true
}

func (r *ConsistentHash) Remove(node Node) {
	for i := 0; i < r.virtualCount; i++ {
		r.circle.Remove(r.hash(fmt.Sprint(node.Name(), i)))
	}

}

func (r *ConsistentHash) Get(key string) (node Node, found bool) {
	if r.circle.Empty() {
		return nil, false
	}

	v, found := r.circle.Get(key)
	if found {
		node = v.(Node) // Go类型断言很尴尬
		return
	}

	return
}
