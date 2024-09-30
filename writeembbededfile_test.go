package raft

import (
	"testing"

	"io/fs"
	"strings"
	"time"

	"github.com/reiver/go-strfs"
)

func TestWriteEmbbededFile(t *testing.T) {

	tests := []struct{
		Name string
		Content string
		Expected string
	}{
		{
			Name: "empty.txt",
			Content: "",
			Expected:
				"empty.txt" +"\n"+
				"0"         +"\n"+
				""+
				"\n\n",
		},



		{
			Name: "apple.txt",
			Content: "APPLE",
			Expected:
				"apple.txt" +"\n"+
				"5"         +"\n"+
				"APPLE"+
				"\n\n",
		},
		{
			Name: "banana.txt",
			Content: "BANANA\n",
			Expected:
				"banana.txt" +"\n"+
				"7"          +"\n"+
				"BANANA\n"+
				"\n\n",
		},
		{
			Name: "cherry.txt",
			Content: "CHERRY\n\n",
			Expected:
				"cherry.txt" +"\n"+
				"8"          +"\n"+
				"CHERRY\n\n"+
				"\n\n",
		},



		{
			Name: "README.md",
			Content: "Hello world!",
			Expected:
				"README.md"    +"\n"+
				"12"           +"\n"+
				"Hello world!"+
				"\n\n",
		},



		{
			Name: "SOMETHING.md",
			Content:
				"# Title" +"\n"+
				""        +"\n"+
				"Hello!"  +"\n"+
				""        +"\n",
			Expected:
				"SOMETHING.md" +"\n"+
				"17"           +"\n"+
				"# Title"      +"\n"+
				""             +"\n"+
				"Hello!"       +"\n"+
				""             +"\n"+
				"\n\n",
		},



		{
			Name: "article.txt",
			Content:
				"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Interdum velit laoreet id donec ultrices tincidunt arcu non sodales. Cras semper auctor neque vitae tempus quam pellentesque nec nam. Cursus turpis massa tincidunt dui ut. Diam vel quam elementum pulvinar etiam non quam. Gravida neque convallis a cras semper. Ornare massa eget egestas purus. Tempor id eu nisl nunc mi ipsum faucibus vitae aliquet. Fames ac turpis egestas maecenas pharetra. Arcu bibendum at varius vel pharetra vel turpis nunc. Integer quis auctor elit sed vulputate mi. Eget velit aliquet sagittis id consectetur purus ut faucibus. Sapien pellentesque habitant morbi tristique senectus."+
				"\n"+
				"\n"+
				"Lorem mollis aliquam ut porttitor leo a diam sollicitudin tempor. Quis commodo odio aenean sed adipiscing. Commodo quis imperdiet massa tincidunt nunc. Quam quisque id diam vel quam elementum pulvinar etiam non. Elit ut aliquam purus sit amet luctus venenatis lectus. Sit amet mauris commodo quis. Placerat vestibulum lectus mauris ultrices eros in. Tristique sollicitudin nibh sit amet commodo nulla facilisi nullam vehicula. Augue interdum velit euismod in. Tellus pellentesque eu tincidunt tortor. Commodo viverra maecenas accumsan lacus vel facilisis. Venenatis a condimentum vitae sapien pellentesque habitant morbi. Et ligula ullamcorper malesuada proin libero nunc consequat interdum varius. Tellus integer feugiat scelerisque varius. Bibendum enim facilisis gravida neque convallis. Nisl nisi scelerisque eu ultrices vitae auctor eu."+
				"\n",
			Expected:
				"article.txt" +"\n"+
				"1573"        +"\n"+
				"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Interdum velit laoreet id donec ultrices tincidunt arcu non sodales. Cras semper auctor neque vitae tempus quam pellentesque nec nam. Cursus turpis massa tincidunt dui ut. Diam vel quam elementum pulvinar etiam non quam. Gravida neque convallis a cras semper. Ornare massa eget egestas purus. Tempor id eu nisl nunc mi ipsum faucibus vitae aliquet. Fames ac turpis egestas maecenas pharetra. Arcu bibendum at varius vel pharetra vel turpis nunc. Integer quis auctor elit sed vulputate mi. Eget velit aliquet sagittis id consectetur purus ut faucibus. Sapien pellentesque habitant morbi tristique senectus."+
				"\n"+
				"\n"+
				"Lorem mollis aliquam ut porttitor leo a diam sollicitudin tempor. Quis commodo odio aenean sed adipiscing. Commodo quis imperdiet massa tincidunt nunc. Quam quisque id diam vel quam elementum pulvinar etiam non. Elit ut aliquam purus sit amet luctus venenatis lectus. Sit amet mauris commodo quis. Placerat vestibulum lectus mauris ultrices eros in. Tristique sollicitudin nibh sit amet commodo nulla facilisi nullam vehicula. Augue interdum velit euismod in. Tellus pellentesque eu tincidunt tortor. Commodo viverra maecenas accumsan lacus vel facilisis. Venenatis a condimentum vitae sapien pellentesque habitant morbi. Et ligula ullamcorper malesuada proin libero nunc consequat interdum varius. Tellus integer feugiat scelerisque varius. Bibendum enim facilisis gravida neque convallis. Nisl nisi scelerisque eu ultrices vitae auctor eu."+
				"\n"+
				"\n\n",
		},
	}

	for testNumber, test := range tests {

		var file fs.File
		{
			var content = strfs.CreateContent(test.Content)

			var regularfile = strfs.RegularFile{
				FileContent: content,
				FileName: test.Name,
				FileModTime: time.Date(2022, 12, 12, 10, 30, 14, 2, time.UTC),
			}

			file = &regularfile
		}

		var buffer strings.Builder

		err := writeEmbeddedFile(&buffer, test.Name, int64(len(test.Content)), file)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("FILE-NAME: %q", test.Name)
			t.Logf("FILE-CONTENT: %q", test.Content)
			t.Logf("EXPECTED-EMBEDDED-FILE:\n%q", test.Expected)
			t.Logf("EXPECTED-EMBEDDED-FILE:\n%s", test.Expected)
			continue
		}

		{
			actual := buffer.String()
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual written embedded file it not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("FILE-CONTENT: %q", test.Content)
				t.Logf("FILE-NAME: %q", test.Name)
				t.Logf("EXPECTED:\n%s", expected)
				t.Logf("ACTUAL:\n%s", actual)
				continue
			}
		}
	}
}
