package raft

import (
	"io"

	"sourcecode.social/reiver/go-erorr"
)

// copyN is similar to io.CopyN(), except that if the number of bytes written (by io.CopyN()) is not 'n' then it returns an error.
func copyN(writer io.Writer, reader io.Reader, n int64) error {

	numWritten, err := io.CopyN(writer, reader, n)
	if nil != err {
		return erorr.Errorf("problem copying %d bytes from reader to writer: %w", n, err)
	}
	if expected, actual := n, numWritten ; expected != actual {
		return erorr.Errorf("expected to write %d bytes but actually wrote %d bytes", expected, actual)
	}

	return nil
}
