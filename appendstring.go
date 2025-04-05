package raft

// AppendString embeds a file in a raft-file with the provided file-name and the contents of the file provided as a `string`.
//
// Use this if you need the write to be atomic.
func AppendString(writer Writer, filename string, content string) error {

	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = gatherString(p, filename, content)

	return putBytes(writer, p)
}
