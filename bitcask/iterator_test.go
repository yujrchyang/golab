package bitcask

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yujrchyang/golab/bitcask/utils"
)

func TestDB_Iterator(t *testing.T) {
	opts := DefaultOptions
	dir, _ := os.MkdirTemp("", "bitcask-go-iterator")
	opts.DirPath = dir
	db, err := Open(opts)
	defer destroyDB(db)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	err = db.Put(utils.GetTestKey(10), utils.RandomValue(10))
	assert.Nil(t, err)

	iterator := db.NewIterator(DefaultIteratorOptions)
	assert.NotNil(t, iterator)
	assert.Equal(t, iterator.Valid(), true)

	db.Put([]byte("annde"), utils.RandomValue(10))
	db.Put([]byte("cnedc"), utils.RandomValue(10))
	db.Put([]byte("aeeue"), utils.RandomValue(10))
	db.Put([]byte("esnue"), utils.RandomValue(10))
	db.Put([]byte("bnede"), utils.RandomValue(10))

	iter1 := db.NewIterator(DefaultIteratorOptions)
	for iter1.Rewind(); iter1.Valid(); iter1.Next() {
		t.Log("key = ", string(iter1.Key()))
	}
	t.Log("----------")

	// 反向迭代
	reserveIterOpt := DefaultIteratorOptions
	reserveIterOpt.Reverse = true
	iter2 := db.NewIterator(reserveIterOpt)
	for iter2.Rewind(); iter2.Valid(); iter2.Next() {
		t.Log("key = ", string(iter2.Key()))
	}
	t.Log("----------")

	// Prefix
	prefixIterOpt := DefaultIteratorOptions
	prefixIterOpt.Prefix = []byte("a")
	iter3 := db.NewIterator(prefixIterOpt)
	for iter3.Rewind(); iter3.Valid(); iter3.Next() {
		t.Log("key = ", string(iter3.Key()))
	}
}
