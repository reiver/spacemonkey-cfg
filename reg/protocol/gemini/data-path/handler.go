package verboten

import (
	"github.com/reiver/spacemonkey-cfg/lib/hdl"
	"github.com/reiver/spacemonkey-cfg/lib/ptcl/gmni"
	"github.com/reiver/spacemonkey-cfg/srv/reg"
)

const (
	name = "protocol/gemini/data-path"
)

func init() {

	var handler hdl.Handler = hdl.HandlerFunc(gmni.DataPath)

	if err := regsrv.Register(handler, name); nil != err {
		panic(err)
	}

}


