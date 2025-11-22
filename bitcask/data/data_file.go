package data

import (
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"path/filepath"

	"github.com/yujrchyang/golab/bitcask/fio"
)

var ErrInvalidCRC = errors.New("invalid crc value, log record maybe corrupted")

const DataFileNameSuffix = ".data"

// 数据文件
type DataFile struct {
	FileId    int32         // 文件 id
	WriteOff  int64         // 文件偏移
	IOManager fio.IOManager // IO 读写管理
}

// 打开新的数据文件
func OpenDataFile(dirPath string, fileId int32) (*DataFile, error) {
	fileName := filepath.Join(dirPath, fmt.Sprintf("%09d", fileId)+DataFileNameSuffix)
	ioManager, err := fio.NewIOManager(fileName)
	if err != nil {
		return nil, err
	}
	return &DataFile{
		FileId:    fileId,
		WriteOff:  0,
		IOManager: ioManager,
	}, nil
}

// 根据 offset 从文件中读取 LogRecord
func (df *DataFile) ReadLogRecord(offset int64) (*LogRecord, int64, error) {
	fileSize, err := df.IOManager.Size()
	if err != nil {
		return nil, 0, err
	}
	var headerBytes int64 = maxLogRecordHeaderSize
	if offset+headerBytes > fileSize {
		headerBytes = fileSize - offset
	}

	// 读取 header 信息
	headerBuf, err := df.readNBytes(headerBytes, offset)
	if err != nil {
		return nil, 0, err
	}
	header, headerSize := decodeLogRecordHeader(headerBuf)
	// 表示读取到了文件末尾
	if header == nil {
		return nil, 0, io.EOF
	}
	if header.crc == 0 && header.keySize == 0 && header.valSize == 0 {
		return nil, 0, io.EOF
	}

	// 取出对应的 key 和 value 的长度
	keySize, valSize := header.keySize, header.valSize
	recordSize := headerSize + int64(keySize) + int64(valSize)
	logRecord := &LogRecord{Type: header.recordType}

	// 开始读取用户实际存储的 key/value 数据
	if keySize > 0 || valSize > 0 {
		kvBuf, err := df.readNBytes(int64(keySize)+int64(valSize), offset+headerSize)
		if err != nil {
			return nil, 0, err
		}

		// 解出 key 和 value
		logRecord.Key = kvBuf[:keySize]
		logRecord.Value = kvBuf[keySize:]
	}

	// 校验数据的有效性
	crc := getLogRecordCRC(logRecord, headerBuf[crc32.Size:headerSize])
	if crc != header.crc {
		return nil, 0, ErrInvalidCRC
	}

	return logRecord, recordSize, nil
}

func (df *DataFile) Write(buf []byte) error {
	n, err := df.IOManager.Write(buf)
	if err != nil {
		return err
	}
	df.WriteOff += int64(n)

	return nil
}

func (df *DataFile) Sync() error {
	return df.IOManager.Sync()
}

func (df *DataFile) Close() error {
	return df.IOManager.Close()
}

func (df *DataFile) readNBytes(n int64, offset int64) (b []byte, err error) {
	b = make([]byte, n)
	_, err = df.IOManager.Read(b, offset)
	return
}
