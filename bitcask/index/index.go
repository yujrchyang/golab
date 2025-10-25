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
}

type Item struct {
	key []byte
	pos *data.LogRecordPos
}

func (ai *Item) Less(bi btree.Item) bool {
	return bytes.Compare(ai.key, bi.(*Item).key) == -1
}
