package main

import (
	"github.com/reiver/spacemonkey-cfg/arg"
	_ "github.com/reiver/spacemonkey-cfg/reg"
	"github.com/reiver/spacemonkey-cfg/srv/log"
	"github.com/reiver/spacemonkey-cfg/srv/reg"

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
		var found bool
		var err error

		value, found, err = regsrv.Get(name)
		if nil != err {
			logsrv.Log("error:", err)
			os.Exit(1)
			return
		}
		if !found {
			value = ""
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
