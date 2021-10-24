package lgr

import (
	"fmt"
	"io"
	"strings"
)

type Logger struct {
	Writer io.Writer
}

func (receiver Logger) Log(a ...interface{}) {
	var w io.Writer = receiver.Writer
	{
		if nil == w {
			return
		}
	}

	var storage strings.Builder
	{
		storage.WriteString("♞ ")
		fmt.Fprintln(&storage, a...)
	}

	io.WriteString(w, storage.String())
}

func (receiver Logger) Logf(format string, a ...interface{}) {
	var w io.Writer = receiver.Writer
	{
		if nil == w {
			return
		}
	}

	var storage strings.Builder
	{
		storage.WriteString("♞ ")
		fmt.Fprintf(&storage, format, a...)
		fmt.Fprintln(&storage)
	}

	io.WriteString(w, storage.String())
}
