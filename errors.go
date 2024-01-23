package raft

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errNilFile        = erorr.Error("raft: nil file")
	errNilFileInfo    = erorr.Error("raft: nil file-info")
	errNilReader      = erorr.Error("raft: nil reader")
	errNilReceiver    = erorr.Error("raft: nil receiver")
	errNilWriteCloser = erorr.Error("raft: nil write-closer")
	errNilWriter      = erorr.Error("raft: nil writer")
)
