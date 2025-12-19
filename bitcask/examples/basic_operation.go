package main

import (
	"fmt"

	"github.com/yujrchyang/golab/bitcask"
)

func main() {
	opts := bitcask.DefaultOptions
	opts.DirPath = "/tmp/bitcask"
	db, err := bitcask.Open(opts)
	if err != nil {
		panic(err)
	}

	err = db.Put([]byte("hello"), []byte("world"))
	if err != nil {
		panic(err)
	}

	val, err := db.Get([]byte("hello"))
	if err != nil {
		panic(err)
	}
	fmt.Println("val = ", string(val))

	err = db.Delete([]byte("hello"))
	if err != nil {
		panic(err)
	}
}
