package index

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yujrchyang/golab/bitcask/data"
)

func TestBTeee_PGD(t *testing.T) {
	bt := NewBTree()

	res := bt.Put(nil, &data.LogRecordPos{Fid: 1, Offset: 100})
	assert.True(t, res)

	res = bt.Put([]byte("a"), &data.LogRecordPos{Fid: 1, Offset: 100})
	assert.True(t, res)
	res = bt.Put([]byte("a"), &data.LogRecordPos{Fid: 1, Offset: 300})
	assert.True(t, res)

	rv := bt.Get([]byte("a"))
	assert.Equal(t, uint32(1), rv.Fid)
	assert.Equal(t, uint64(300), rv.Offset)

	res = bt.Delete([]byte("a"))
	assert.True(t, res)
	res = bt.Delete([]byte("a"))
	assert.False(t, res)
}
