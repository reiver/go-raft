package raft

import (
	"io/fs"

	"sourcecode.social/reiver/go-erorr"
)

// fileInfo is similar to fs.File.Stat(), except that if fs.File.Stat() returns 'nil' for 'fs.FileInfo' then it returns an error.
func fileInfo(file fs.File) (fs.FileInfo, error) {

	if nil == file {
		return nil, errNilFile
	}

	var fileinfo fs.FileInfo
	{
		var err error

		fileinfo, err = file.Stat()
		if nil != err {
			return nil, erorr.Errorf("raft: problem getting file-info: %w", err)
		}
		if nil == fileinfo {
			return nil, errNilFileInfo
		}
	}

	return fileinfo, nil
}
