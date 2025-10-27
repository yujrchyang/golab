package bitcask

type Options struct {
	// 数据库目录
	DirPath string

	// 数据文件大小
	DataFileSize uint64

	// 每次写数据是否持久化
	SyncWrite bool
}
