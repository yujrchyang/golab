package index

import (
	"bytes"

	"github.com/google/btree"
	"github.com/yujrchyang/golab/bitcask/data"
)

// 抽象索引接口，后续如果想要接入其他的数据结构，则直接实现这个接口即可
type Indexer interface {
	// 向索引中存储 key 对应的数据位置信息
	Put(key []byte, pos *data.LogRecordPos) bool

	// 根据 key 取出对应的索引位置信息
	Get(key []byte) *data.LogRecordPos

	// 根据 key 删除对应的索引位置信息
	Delete(key []byte) bool

	// 索引中的数据量
	Size() int

	// 索引迭代器
	Iterator(reverse bool) Iterator
}

type IndexType = int8

const (
	// Btree 索引
	Btree IndexType = iota + 1

	// 自适应基数树
	ART
)

func NewIndexer(typ IndexType) Indexer {
	switch typ {
	case Btree:
		return NewBTree()
	case ART:
		return nil
	default:
		panic("unsupport index type")
	}
}

type Item struct {
	key []byte
	pos *data.LogRecordPos
}

func (ai *Item) Less(bi btree.Item) bool {
	return bytes.Compare(ai.key, bi.(*Item).key) == -1
}

// 通用索引迭代器
type Iterator interface {
	// 重新回到迭代器的起点，即第一个数据
	Rewind()

	// 根据传入的 key 查找到第一个大于（或小于）等于的目标 key，从这个 key 开始便利
	Seek(key []byte)

	// 跳转到下一个 key
	Next()

	// 是否有效，即是否已经遍历完所有的 key，用于退出遍历
	Valid() bool

	// 当前遍历位置的 key 数据
	Key() []byte

	// 当前遍历位置的 value 数据
	Value() *data.LogRecordPos

	// 关闭迭代器，释放相应资源
	Close()
}
