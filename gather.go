package raft

import (
	"strconv"
)

func gatherBytes(p []byte, filename string, content []byte) []byte {

	p = append(p, filename...)
	p = append(p, '\n')

	var length string = strconv.FormatUint(uint64(len(content)), 10)

	p = append(p, length...)
	p = append(p, '\n')

	p = append(p, content...)
	p = append(p, '\n')

	p = append(p, '\n')

	return p
}

func gatherString(p []byte, filename string, content string) []byte {

	p = append(p, filename...)
	p = append(p, '\n')

	var length string = strconv.FormatUint(uint64(len(content)), 10)

	p = append(p, length...)
	p = append(p, '\n')

	p = append(p, content...)
	p = append(p, '\n')

	p = append(p, '\n')

	return p
}
