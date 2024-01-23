package raft

import (
	"io/fs"
)

// fileSize returns the file-size of an fs.File.
func fileSize(file fs.File) (int64, error) {

	if nil == file {
		return 0, errNilFile
	}

	var fileinfo fs.FileInfo
	{
		var err error

		fileinfo, err = fileInfo(file)
		if nil != err {
			return 0, err
		}
	}

	return fileinfo.Size(), nil
}
