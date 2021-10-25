package logsrv

import (
	"github.com/reiver/spacemonkey-cfg/arg"
	"github.com/reiver/spacemonkey-cfg/lib/lgr"

	"os"
)

var (
	logger lgr.Logger = lgr.Logger{
		Writer: os.Stderr,
	}
)

func Log(a ...interface{}) {
	if !arg.Verbose {
		return
	}

	logger.Log(a...)
}

func Logf(format string, a ...interface{}) {
	if !arg.Verbose {
		return
	}

	logger.Logf(format, a...)
}
