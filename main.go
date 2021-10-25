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

	// If the user calls spacemonkey-cfg wit the --list flag then,
	// output all the (registered) configuration names,
	// and then exit.
	if arg.List {
		logsrv.Log("--list — will try to list all (register) configuration names")

		regsrv.ForEachName(func(name string){
			fmt.Fprintln(os.Stdout, name)
		})
		return
	}

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
