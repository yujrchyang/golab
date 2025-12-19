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
