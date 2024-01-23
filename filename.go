package raft

import (
	"io/fs"
)

// fileName returns the file-name of an fs.File.
func fileName(file fs.File) (string, error) {

	if nil == file {
		return "", errNilFile
	}

	var fileinfo fs.FileInfo
	{
		var err error

		fileinfo, err = fileInfo(file)
		if nil != err {
			return "", err
		}
	}

	return fileinfo.Name(), nil
}
