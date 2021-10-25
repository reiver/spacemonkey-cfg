package main

import (
	"github.com/reiver/spacemonkey-cfg/arg"
	"github.com/reiver/spacemonkey-cfg/lib/ptcl/gmni"
	"github.com/reiver/spacemonkey-cfg/lib/ptcl/rook"
	"github.com/reiver/spacemonkey-cfg/srv/log"

	"fmt"
	"os"
)

func main() {

	var name string
	{
		if len(arg.Values) < 1 {
			os.Exit(1)
			return
		}

		name = arg.Values[0]
	}
	logsrv.Logf("name: %q", name)

	var value string
	{
		switch name {
		case "protocol/gemini/data-path":
			path, err := gmni.DataPath()
			if nil != err {
				logsrv.Log("error:", err)
				os.Exit(1)
				return
			}

			value = path
		case "protocol/rook/data-path":
			path, err := rook.DataPath()
			if nil != err {
				logsrv.Log("error:", err)
				os.Exit(1)
				return
			}

			value = path
		default:
			logsrv.Logf("unknown name: %q", name)
			os.Exit(1)
			return
		}
	}
	logsrv.Logf("value: %q", value)

	var trailer string
	{
		if arg.EOL {
			trailer = "\n"
		}
	}

	logsrv.Log("♜ shâh mât")
	fmt.Fprintf(os.Stdout, "%s%s", value, trailer)
}
