//go:build !linux
// +build !linux

package sys

import "fmt"

type OtherExtentReader struct{}

func (r OtherExtentReader) GetExtentStats(fd uintptr) (*ExtentStats, error) {
	return nil, fmt.Errorf("file extent metrics not supported on this platform")
}

var DefaultReader FileExtentReader = OtherExtentReader{}
