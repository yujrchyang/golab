package data

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenDataFile(t *testing.T) {
	dbDir := "/tmp/bitcask"

	fileID := int32(0)
	dataFile1, err := OpenDataFile(dbDir, fileID)
	assert.Nil(t, err)
	assert.NotNil(t, dataFile1)
	os.Remove(filepath.Join(dbDir, fmt.Sprintf("%09d", fileID)+DataFileNameSuffix))

	fileID = int32(111)
	dataFile2, err := OpenDataFile(dbDir, fileID)
	assert.Nil(t, err)
	assert.NotNil(t, dataFile2)

	fileID = int32(111)
	dataFile3, err := OpenDataFile(dbDir, fileID)
	assert.Nil(t, err)
	assert.NotNil(t, dataFile3)
	os.Remove(filepath.Join(dbDir, fmt.Sprintf("%09d", fileID)+DataFileNameSuffix))
}

func TestDataFile_Write(t *testing.T) {
	dbDir := "/tmp/bitcask"

	fileID := int32(0)
	dataFile, err := OpenDataFile(dbDir, fileID)
	assert.Nil(t, err)
	assert.NotNil(t, dataFile)

	err = dataFile.Write([]byte("aaa"))
	assert.Nil(t, err)

	os.Remove(filepath.Join(dbDir, fmt.Sprintf("%09d", fileID)+DataFileNameSuffix))
}

func TestDataFile_Close(t *testing.T) {
	dbDir := "/tmp/bitcask"

	fileID := int32(0)
	dataFile, err := OpenDataFile(dbDir, fileID)
	assert.Nil(t, err)
	assert.NotNil(t, dataFile)

	err = dataFile.Write([]byte("aaa"))
	assert.Nil(t, err)

	err = dataFile.Close()
	assert.Nil(t, err)

	os.Remove(filepath.Join(dbDir, fmt.Sprintf("%09d", fileID)+DataFileNameSuffix))
}

func TestDataFile_Sync(t *testing.T) {
	dbDir := "/tmp/bitcask"

	fileID := int32(0)
	dataFile, err := OpenDataFile(dbDir, fileID)
	assert.Nil(t, err)
	assert.NotNil(t, dataFile)

	err = dataFile.Write([]byte("aaa"))
	assert.Nil(t, err)

	err = dataFile.Sync()
	assert.Nil(t, err)

	err = dataFile.Close()
	assert.Nil(t, err)

	os.Remove(filepath.Join(dbDir, fmt.Sprintf("%09d", fileID)+DataFileNameSuffix))
}

func TestDataFile_ReadLogRecord(t *testing.T) {
	dbDir := "/tmp/bitcask"

	fileID := int32(0)
	dataFile, err := OpenDataFile(dbDir, fileID)
	assert.Nil(t, err)
	assert.NotNil(t, dataFile)

	// 只有一条 LogRecord
	rec1 := &LogRecord{
		Key:   []byte("name"),
		Value: []byte("bitcask kv go"),
	}
	res1, size1 := EncodeLogRecord(rec1)
	err = dataFile.Write(res1)
	assert.Nil(t, err)

	readRec1, readSize1, err := dataFile.ReadLogRecord(0)
	assert.Nil(t, err)
	assert.Equal(t, rec1, readRec1)
	assert.Equal(t, size1, readSize1)

	// 多条 LogRecord 从不同的位置读取
	rec2 := &LogRecord{
		Key:   []byte("name"),
		Value: []byte("new value"),
	}
	res2, size2 := EncodeLogRecord(rec2)
	err = dataFile.Write(res2)
	assert.Nil(t, err)

	readRec2, readSize2, err := dataFile.ReadLogRecord(readSize1)
	assert.Nil(t, err)
	assert.Equal(t, rec2, readRec2)
	assert.Equal(t, size2, readSize2)

	// 被删除的数据子啊数据文件末尾
	rec3 := &LogRecord{
		Key:   []byte("name"),
		Value: []byte(""),
		Type:  LogRecordDeleted,
	}
	res3, size3 := EncodeLogRecord(rec3)
	err = dataFile.Write(res3)
	assert.Nil(t, err)

	readRec3, readSize3, err := dataFile.ReadLogRecord(readSize1 + readSize2)
	assert.Nil(t, err)
	assert.Equal(t, rec3, readRec3)
	assert.Equal(t, size3, readSize3)

	os.Remove(filepath.Join(dbDir, fmt.Sprintf("%09d", fileID)+DataFileNameSuffix))
}
