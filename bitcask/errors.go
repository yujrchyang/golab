package bitcask

import "errors"

var (
	ErrKeyIsEmpty             = errors.New("key is empty")
	ErrKeyNotFound            = errors.New("key not found")
	ErrIndexUpdateFailed      = errors.New("failed to update index")
	ErrDataFileNotFound       = errors.New("data file not found")
	ErrDataDirectoryCorrupted = errors.New("the database directory maybe corrupted")
)
