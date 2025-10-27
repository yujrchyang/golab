package data

import "github.com/yujrchyang/golab/bitcask/fio"

// 数据文件
type DataFile struct {
	FileId    uint32        // 文件 id
	WriteOff  uint64        // 文件偏移
	IOManager fio.IOManager // IO 读写管理
}

// 打开新的数据文件
func OpenDataFile(dirPath string, fileId uint32) (*DataFile, error) {
	return nil, nil
}

func (df *DataFile) ReadLogRecord(offset uint64) (*LogRecord, error) {
	return nil, nil
}

func (df *DataFile) Write(buf []byte) error {
	return nil
}

func (df *DataFile) Sync() error {
	return nil
}
