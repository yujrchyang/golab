package sys

// ExtentStats 存储文件的碎片统计信息
type ExtentStats struct {
	ExtentCount uint32 // 物理碎片
}

// FileExtentReader 定义获取碎片信息的接口
type FileExtentReader interface {
	GetExtentStats(fd uintptr) (*ExtentStats, error)
}
