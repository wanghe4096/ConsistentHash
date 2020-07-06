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

type Ring struct {
	buff         *treemap.Map
	hash         HashFunc
	virtualCount int // 虚拟节点的数量
}

/**
用你熟悉的编程语言实现一致性 hash 算法。
编写测试用例测试这个算法，测试 100 万 KV 数据，10 个服务器节点的情况下，计算这些 KV 数据在服务器上分布数量的标准差，以评估算法的存储负载不均衡性。
*/
func NewRing(hashFunc HashFunc) *Ring {
	return &Ring{buff: treemap.NewWithIntComparator(), hash: hashFunc}
}

func (r *Ring) Add(node Node) bool {
	// 添加实体节点
	r.buff.Put(r.hash(node.Name()), node)

	// 初始化虚拟节点
	for i := 0; i < r.virtualCount; i++ {
		r.buff.Put(r.hash(fmt.Sprint(node.Name(), i)), node)
	}

	return true
}

func (r *Ring) Get(key string) (node Node, found bool) {
	v, found := r.buff.Get(key)
	if found {
		node = v.(Node) // Go类型断言很尴尬
	}
	return
}
