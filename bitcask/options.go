package bitcask

type Options struct {
	// 数据库目录
	DirPath string

	// 数据文件大小
	DataFileSize int64

	// 每次写数据是否持久化
	SyncWrite bool

	// 索引类型
	IndexType IndexerType
}

// 索引迭代器配置项
type IteratorOptions struct {
	// 遍历前缀为指定值的 key，默认为空
	Prefix []byte
	// 是否反向遍历，默认 false 是正向
	Reverse bool
}

type WriteBatchOptions struct {
	// 一个批次中最大的数据量
	MaxBatchNum uint
	// 提交时是否 sync 持久化
	SyncWrites bool
}

type IndexerType = int8

const (
	BTree IndexerType = iota + 1
	ART
)

var DefaultOptions = Options{
	DirPath:      "/tmp/bitcask",
	DataFileSize: 256 * 1024 * 1024,
	SyncWrite:    false,
	IndexType:    BTree,
}

var DefaultIteratorOptions = IteratorOptions{
	Prefix:  nil,
	Reverse: false,
}

var DefaultWriteBatchOptions = WriteBatchOptions{
	MaxBatchNum: 10000,
	SyncWrites:  true,
}
