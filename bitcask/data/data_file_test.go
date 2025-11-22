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
