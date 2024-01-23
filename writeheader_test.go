package raft

import (
	"testing"

	"strings"
)

func TestWriteHeader(t *testing.T) {

	var expected string = "RAFT/1\n\n"

	var actual string
	{
		var buffer strings.Builder


		err := writeHeader(&buffer)
		if nil != err {
			t.Errorf("Did not expect an error but actually got one.")
			t.Logf("ERROR: %q", err)
			return
		}

		actual = buffer.String()
	}

	{
		if expected != actual {
			t.Errorf("The actual written value is not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			return
		}
	}
}
