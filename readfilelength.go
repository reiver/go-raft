package raft

import (
	"io"
	"strconv"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-rcv"
)

func readFileLength(reader io.Reader) (length int64, err error) {

	var fileLength string

	_, err = rcv.Freadln(reader, &fileLength)
	if nil != err {
		var nada int64
		return nada, err
	}

	n64, err := strconv.ParseInt(fileLength, 10, 64)
	if nil != err {
		var nada int64
		return nada, err
	}
	if n64 < 0 {
		var nada int64
		return nada, erorr.Errorf("raft: bad file-length (%q)", fileLength)
	}

	return n64, nil
}
