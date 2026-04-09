//go:build linux
// +build linux

package sys

/*
#define _GNU_SOURCE
#include <linux/fs.h>
#include <linux/fiemap.h>
#include <sys/ioctl.h>

static unsigned long get_fiemap_ioc() {
    return FS_IOC_FIEMAP;
}
*/
import "C"
import (
	"syscall"
	"unsafe"
)

var FS_IOC_FIEMAP = uintptr(C.get_fiemap_ioc())

type Fiemap struct {
	Start         uint64
	Length        uint64
	Flags         uint32
	MappedExtents uint32
	ExtentCount   uint32
	Reserved      uint32
}

type LinuxExtentReader struct{}

func (r LinuxExtentReader) GetExtentStats(fd uintptr) (*ExtentStats, error) {
	fiemap := Fiemap{
		Start:  0,
		Length: ^uint64(0), // 扫描全文件
	}

	_, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		fd,
		FS_IOC_FIEMAP,
		uintptr(unsafe.Pointer(&fiemap)),
	)
	if errno != 0 {
		return nil, errno
	}

	return &ExtentStats{
		ExtentCount: fiemap.MappedExtents,
	}, nil
}

var DefaultReader FileExtentReader = LinuxExtentReader{}
