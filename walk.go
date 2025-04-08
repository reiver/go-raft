package raft

import (
	"io"

	"github.com/reiver/go-erorr"
)

const (
	errNilWalkFunc = erorr.Error("raft: nil walk-func")
)

type WalkFunc func(filename string, reader io.Reader)error

func Walk(reader io.Reader, fn WalkFunc) error {

	if nil == fn {
		return errNilWalkFunc
	}

	for i:=uint64(0); true ; i++ {

		fileName, err := readFileName(reader)
		if erorr.Is(err, io.EOF) {
			return nil
		}
		if nil != err {
			return erorr.Errorf("raft: problem reading file-name for file #%d: %w", i, err)
		}
		if "" == fileName {
	/////////////// CONTINUE
			continue
		}
		if 0 == i && magic == fileName {
	/////////////// CONTINUE
			continue
		}

		fileLength, err := readFileLength(reader)
		if nil != err {
			return erorr.Errorf("raft: problem reading file-length for file #%d (%q): %w", i, fileName, err)
		}

		var limitedReader io.Reader = io.LimitReader(reader, fileLength)

		err = fn(fileName, limitedReader)
		if nil != err {
			return err
		}

		// We attempt to drain the limited-reader just in walk the WalkFunc didn't fully read from the reader.
		// We need to do this so that the file-cursor moved after this embedded-file in the raft-file.
		drain(limitedReader)
	}

	return nil
}
