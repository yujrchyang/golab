package data

import "encoding/binary"

type LogRecordType = byte

const (
	LogRecordNormal LogRecordType = iota
	LogRecordDeleted
)

// crc type keySize valSize
const maxLogRecordHeaderSize = 4 + 1 + binary.MaxVarintLen32*2

// 写入到数据文件的记录
// 之所以叫日志，是因为数据文件中的数据是追加写入的，类似日志格式
type LogRecord struct {
	Key   []byte
	Value []byte
	Type  LogRecordType
}

type logRecordHeader struct {
	crc        uint32        // crc 校验值
	recordType LogRecordType // 表示 LogRecord 的类型
	keySize    uint32        // key 的长度
	valSize    uint32        // val 的长度
}

type LogRecordPos struct {
	Fid    int32 // 文件 ID，表示将数据存储到了那个文件当中
	Offset int64 // 偏移，表示将数据存储到了数据文件中的哪个位置
}

// 对 LogRecord 进行编码，返回字节数组及长度
func EncodeLogRecord(logRecord *LogRecord) ([]byte, int64) {
	return nil, 0
}

// 对 LogRecordHeader 进行解码
func decodeLogRecordHeader(buf []byte) (*logRecordHeader, int64) {
	return nil, 0
}

func getLogRecordCRC(lr *LogRecord, header []byte) uint32 {
	return 0
}
