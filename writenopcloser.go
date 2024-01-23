package raft

import (
	"io"
)

type internalWriteNopCloser struct {
	writer io.Writer
}

func (receiver internalWriteNopCloser) Write(p []byte) (n int, err error) {
	return receiver.writer.Write(p)
}

func (internalWriteNopCloser) Close() error {
	return nil
}
