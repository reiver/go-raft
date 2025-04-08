package raft

import (
	"io"
)

func drain(reader io.Reader) {
	if nil == reader {
		return
	}

	var buffer [1]byte
	var p []byte = buffer[:]

	var err error

	for nil == err {
		_, err = reader.Read(p)
	}
}
