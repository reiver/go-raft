package raft

import (
	"github.com/reiver/go-erorr"
)

func putBytes(writer Writer, bytes []byte) error {

	if nil == writer {
		return errNilWriter
	}

	var length int = len(bytes)

	if length <= 0 {
		return nil
	}

	{
		n, err := writer.Write(bytes)
		if nil != err {
			return erorr.Errorf("raft: problem writing %d bytes: %s", length, err)
		}
		if n != length {
			return erorr.Errorf("raft: expected to write %d bytes but actually wrote %d bytes", length, n)
		}
	}

	{
		err := writer.Sync()
		if nil != err {
			return erorr.Errorf("raft: problem sync'ing the appended %d bytes: %s", length, err)
		}
	}

	return nil
}
