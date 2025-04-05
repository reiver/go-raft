package raft

// AppendBytes embeds a file in a raft-file with the provided file-name and the contents of the file provided as []byte.
//
// Use this if you need the write to be atomic.
func AppendBytes(writer Writer, filename string, content []byte) error {

	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = gatherBytes(p, filename, content)

	return putBytes(writer, p)
}
