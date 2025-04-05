package raft

import (
	"io"
)

type Writer interface {
	io.Writer
	Sync() error
}
