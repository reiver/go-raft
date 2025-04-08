package raft

import (
	"io"

	"github.com/reiver/go-rcv"
)

func readFileName(reader io.Reader) (fileName string, err error) {

	_, err = rcv.Freadln(reader, &fileName)
	if nil != err {
		var nada string
		return nada, err
	}

	return fileName, nil
}
