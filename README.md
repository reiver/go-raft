# go-raft

Package **raft** provides tools for reading and writing [RAFT files](https://github.com/reiver/raft-format), for the Go programming language.

The **raft format** is a very simple and easy to understand **archive format** and **container format** that can combine multiple files into a single aggregate file.

**Archive Formats** and **Container Formats** have many use-case:
* backups,
* eBooks,
* file-systems,
* image galleries,
* journals,
* music albums,
* photo albums,
* software packages,
* website archives,
* _etc_.

Basically, any use-case where you need to combine multiple files into a single aggregate file.

The **raft format** is similar to other **archive formats**, such as the  **ar format**, the **cpio format**, the **shar format**, the **tar format**, and the **WARC format** — but is designed to be easier to understand and implement than most (probably all) of the other **archive formats** and **container formats**.

The **raft format** is meant to be both programmer-legible and programmer-friendly.

For more information on the **raft format** see: https://github.com/reiver/raft-format

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-raft

[![GoDoc](https://godoc.org/github.com/reiver/go-raft?status.svg)](https://godoc.org/github.com/reiver/go-raft)

## Examples

Here is an example of writing a **raft** file:

```go
import "github.com/reiver/go-raft"

// ...

raftfile, err := raft.CreateFileWriter(filename)
if nil != er {
	return err
}
defer raftfile.Close()

err := raftfile.EmbedFromFile(file)
if nil != err{
	return err
}

err := raftfile.EmbedFromFile(file1)
if nil != err{
	return err
}

err := raftfile.EmbedFromFile(file2)
if nil != err{
	return err
}

err := raftfile.EmbedFromFileRenamed("newname.jpeg", file2)
if nil != err{
	return err
}

err := raftfile.EmbedFromString("something.txt", "Hello world!\n\nHow are you?\n")
if nil != err{
	return err
}

err := raftfile.EmbedFromBytes("something.txt", []byte{'H','i','!','\n'})
if nil != err{
	return err
}
```

## Import

To import package **raft** use `import` code like the follownig:
```
import "github.com/reiver/go-raft"
```

## Installation

To install package **raft** do the following:
```
GOPROXY=direct go get github.com/reiver/go-raft
```

## Author

Package **raft** was written by [Charles Iliya Krempeaux](http://reiver.link)
