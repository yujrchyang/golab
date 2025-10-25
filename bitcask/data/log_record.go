package data

type LogRecordPos struct {
	Fid    uint32 // 文件 ID，表示将数据存储到了那个文件当中
	Offset uint64 // 偏移，表示将数据存储到了数据文件中的哪个位置
}
