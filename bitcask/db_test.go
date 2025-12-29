package bitcask

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yujrchyang/golab/bitcask/utils"
)

func destroyDB(db *DB) {
	if db != nil {
		if err := db.Close(); err != nil {
			panic(err)
		}
		err := os.RemoveAll(db.options.DirPath)
		if err != nil {
			panic(err)
		}
	}
}

func TestDB_ListKeys(t *testing.T) {
	opts := DefaultOptions
	dir, _ := os.MkdirTemp("", "bitcask-go-listkeys")
	opts.DirPath = dir
	db, err := Open(opts)
	defer destroyDB(db)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	// 数据库为空
	keys := db.ListKeys()
	assert.Equal(t, 0, len(keys))

	// 只有一条数据
	err = db.Put(utils.GetTestKey(11), utils.RandomValue(20))
	assert.Nil(t, err)
	keys2 := db.ListKeys()
	assert.Equal(t, 1, len(keys2))
}

func TestDB_Fold(t *testing.T) {
	opts := DefaultOptions
	dir, _ := os.MkdirTemp("", "bitcask-go-fold")
	opts.DirPath = dir
	db, err := Open(opts)
	defer destroyDB(db)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	err = db.Put(utils.GetTestKey(11), utils.RandomValue(20))
	assert.Nil(t, err)

	err = db.Fold(func(key, value []byte) bool {
		t.Log(string(key))
		t.Log(string(value))
		return true
	})
	assert.Nil(t, err)
}
