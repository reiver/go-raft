package raft

import (
	"io"

	"github.com/reiver/go-erorr"
)

// writeString is similar to io.WriteString(), except that if the number of bytes written (by io.WriteString()) is not 'len(str)' then it returns an error.
func writeString(writer io.Writer, str string) error {

	n, err := io.WriteString(writer, str)
	if nil != err {
		return erorr.Errorf("problem writing string: %w", err)
	}
	if expected, actual := len(str), n ; expected != actual {
		return erorr.Errorf("expected to write %d bytes but actually wrote %d bytes", expected, actual)
	}

	return nil
}
