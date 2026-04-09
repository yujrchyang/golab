package sys

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetExtentStats(t *testing.T) {
	// 1. 创建临时文件
	tmpFile, err := os.CreateTemp("", "test_chunk_*")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// 2. 模拟空洞写入
	// 写入 4K，跳过 1M 再写入 4K
	data := make([]byte, 4096)
	_, _ = tmpFile.Write(data)
	_, _ = tmpFile.Seek(1024*1024, 1)
	_, _ = tmpFile.Write(data)

	// 3. 调用接口
	stats, err := DefaultReader.GetExtentStats(tmpFile.Fd())
	assert.NoError(t, err)
	assert.NotNil(t, stats)
	t.Logf("File: %s, Extents: %d", tmpFile.Name(), stats.ExtentCount)
	assert.GreaterOrEqual(t, int(stats.ExtentCount), 1)
}
