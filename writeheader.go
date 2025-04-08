package raft

import (
	"io"

	"github.com/reiver/go-erorr"
)

// writeHeader writes the header of a RAFT file.
//
// It writes:
//
//	"RAFT/1\n\n"
//
// Technically, the 2nd "\n" isn't necessary.
// Technically, this is all that is necessary is "RAFT/1\n"
//
// The 2nd "\n" is added to make the resulting RAFT file easier to read (if a person was to look at it and try to read it).
func writeHeader(writer io.Writer) error {

	if nil == writer {
		return errNilWriter
	}

	{
		var header string = magic + eoleol

		err := writeString(writer, header)
		if nil != err {
			return erorr.Errorf("raft: problem writing first 2 lines of raft file — %q: %w", header, err)
		}
	}

	return nil
}

