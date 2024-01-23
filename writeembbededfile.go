package raft

import (
	"io"
	"strconv"

	"sourcecode.social/reiver/go-erorr"
)

// writeEmbeddedFile embeds a(nother) file into a RAFT file using the provided file-name, file-size, and file-content.
//
// Note that this adds an additional "\n" after the file-content.
// Technically it isn't necessary.
// But it is added to make the resulting raft file easier to read (if a person was to try to read it).
func writeEmbeddedFile(writer io.Writer, filename string, filesize int64, filecontent io.Reader) error {

	if nil == writer {
		return errNilWriter
	}

	if nil == filecontent {
		return errNilReader
	}

	// file name
	{
		var value string = filename

		{
			err := writeString(writer, value)
			if nil != err {
				return erorr.Errorf("raft: problem writing embedded-file's file-name (i.e., %q): %w", value, err)
			}
		}

		{
			err := writeString(writer, eol)
			if nil != err {
				return erorr.Errorf("raft: problem writing end-of-line (i.e., %q) after writing embedded-file's file-name (i.e., %q): %w", eol, value, err)
			}
		}
	}

	// file size
	{
		var value string = strconv.FormatInt(filesize, 10)

		{
			err := writeString(writer, value)
			if nil != err {
				return erorr.Errorf("raft: problem writing embedded-file's file-size — %q: %w", value, err)
			}
		}

		{
			err := writeString(writer, eol)
			if nil != err {
				return erorr.Errorf("raft: problem writing end-of-line (i.e., %q) after writing embedded-file's file-size (i.e., %q): %w", eol, value, err)
			}
		}
	}

	// file content
	{
		{
			err := copyN(writer, filecontent, filesize)
			if nil != err {
				return erorr.Errorf("raft: problem writing embedded-file's file-content: %w", err)
			}
		}

		{
			var value string = eoleol // <---- Note this is a 'eoleol' rather than just a single 'eol'. This is because the 2nd 'eol' add a blank line that makes it easier for a programmer to read the file.

			err := writeString(writer, value)
			if nil != err {
				return erorr.Errorf("raft: problem writing 2 end-of-lines (i.e., %q) after writing file-content: %w", value, err)
			}
		}
	}

	return nil
}
