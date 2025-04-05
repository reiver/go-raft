package raft

import (
	"strconv"
)

func gatherBytes(p []byte, filename string, bytes []byte) []byte {

	p = append(p, filename...)
	p = append(p, '\n')

	var length string = strconv.FormatUint(uint64(len(bytes)), 10)

	p = append(p, length...)
	p = append(p, '\n')

	p = append(p, bytes...)
	p = append(p, '\n')

	p = append(p, '\n')

	return p
}
