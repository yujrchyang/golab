package index

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yujrchyang/golab/bitcask/data"
)

func TestBTree_PGD(t *testing.T) {
	bt := NewBTree()

	res := bt.Put(nil, &data.LogRecordPos{Fid: 1, Offset: 100})
	assert.True(t, res)

	res = bt.Put([]byte("a"), &data.LogRecordPos{Fid: 1, Offset: 100})
	assert.True(t, res)
	res = bt.Put([]byte("a"), &data.LogRecordPos{Fid: 1, Offset: 300})
	assert.True(t, res)

	rv := bt.Get([]byte("a"))
	assert.Equal(t, int32(1), rv.Fid)
	assert.Equal(t, int64(300), rv.Offset)

	res = bt.Delete([]byte("a"))
	assert.True(t, res)
	res = bt.Delete([]byte("a"))
	assert.False(t, res)
}

func TestBTree_Iterator(t *testing.T) {
	bt1 := NewBTree()
	iter1 := bt1.Iterator(false)
	assert.Equal(t, false, iter1.Valid())

	// 2. BTree 有数据的情况
	bt1.Put([]byte("ccde"), &data.LogRecordPos{Fid: 1, Offset: 10})
	iter2 := bt1.Iterator(false)
	assert.Equal(t, true, iter2.Valid())

	// 3. 多条数据
	bt1.Put([]byte("aaa"), &data.LogRecordPos{Fid: 1, Offset: 10})
	bt1.Put([]byte("bbb"), &data.LogRecordPos{Fid: 1, Offset: 10})
	bt1.Put([]byte("ccc"), &data.LogRecordPos{Fid: 1, Offset: 10})
	iter3 := bt1.Iterator(false)
	for iter3.Rewind(); iter3.Valid(); iter3.Next() {
		t.Log("key = ", string(iter3.Key()))
	}

	// 4. 测试 seek
	iter4 := bt1.Iterator(false)
	for iter4.Seek([]byte("cc")); iter4.Valid(); iter4.Next() {
		t.Log("key = ", string(iter4.Key()))
	}

	// 5. 反向遍历
	iter5 := bt1.Iterator(true)
	for iter5.Seek([]byte("bb")); iter5.Valid(); iter5.Next() {
		t.Log("key = ", string(iter5.Key()))
	}
}
