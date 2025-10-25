package fio

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFileIOManager(t *testing.T) {
	fio, err := NewFileIOManager(filepath.Join("/tmp/bitcask", "a.data"))
	assert.NotNil(t, fio)
	assert.Nil(t, err)
}

func TestFileIO_Write(t *testing.T) {
	path := filepath.Join("/tmp/bitcask", "a.data")

	fio, err := NewFileIOManager(path)
	assert.NotNil(t, fio)
	assert.Nil(t, err)

	n, err := fio.Write([]byte(""))
	assert.Nil(t, err)
	assert.Equal(t, 0, n)

	n, err = fio.Write([]byte("bitcask"))
	assert.Nil(t, err)
	assert.Equal(t, 7, n)

	os.Remove(path)
}

func TestFileIO_Read(t *testing.T) {
	path := filepath.Join("/tmp/bitcask", "a.data")

	fio, err := NewFileIOManager(path)
	assert.NotNil(t, fio)
	assert.Nil(t, err)

	n, err := fio.Write([]byte("key-a"))
	assert.Nil(t, err)
	assert.Equal(t, 5, n)

	n, err = fio.Write([]byte("key-b"))
	assert.Nil(t, err)
	assert.Equal(t, 5, n)

	b := make([]byte, 5)
	n, err = fio.Read(b, 0)
	assert.Nil(t, err)
	assert.Equal(t, 5, n)
	assert.Equal(t, []byte("key-a"), b)

	b1 := make([]byte, 5)
	n, err = fio.Read(b1, 5)
	assert.Nil(t, err)
	assert.Equal(t, 5, n)
	assert.Equal(t, []byte("key-b"), b1)

	os.Remove(path)
}

func TestFileIO_Sync(t *testing.T) {
	path := filepath.Join("/tmp/bitcask", "a.data")

	fio, err := NewFileIOManager(path)
	assert.NotNil(t, fio)
	assert.Nil(t, err)

	n, err := fio.Write([]byte("key-a"))
	assert.Nil(t, err)
	assert.Equal(t, 5, n)

	err = fio.Sync()
	assert.Nil(t, err)

	os.Remove(path)
}

func TestFileIO_Close(t *testing.T) {
	path := filepath.Join("/tmp/bitcask", "a.data")

	fio, err := NewFileIOManager(path)
	assert.NotNil(t, fio)
	assert.Nil(t, err)

	n, err := fio.Write([]byte("key-a"))
	assert.Nil(t, err)
	assert.Equal(t, 5, n)

	err = fio.Close()
	assert.Nil(t, err)

	os.Remove(path)
}
