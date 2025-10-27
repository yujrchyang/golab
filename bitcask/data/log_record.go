package data

type LogRecordType = byte

const (
	LogRecordNormal LogRecordType = iota
	LogRecordDeleted
)

// 写入到数据文件的记录
// 之所以叫日志，是因为数据文件中的数据是追加写入的，类似日志格式
type LogRecord struct {
	Key   []byte
	Value []byte
	Type  LogRecordType
}

type LogRecordPos struct {
	Fid    uint32 // 文件 ID，表示将数据存储到了那个文件当中
	Offset uint64 // 偏移，表示将数据存储到了数据文件中的哪个位置
}

// 对 LogRecord 进行编码，返回字节数组及长度
func EncodeLogRecord(logRecord *LogRecord) ([]byte, uint64) {
	return nil, 0
}
