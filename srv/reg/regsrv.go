package regsrv

import (
	"github.com/reiver/spacemonkey-cfg/lib/hdl"
	"github.com/reiver/spacemonkey-cfg/lib/reg"
)

var (
	registry reg.Registrar
)

func Register(handler hdl.Handler, name string) error {
	return registry.Register(handler, name)
}

func Get(name string) (string, bool, error) {
	return registry.Get(name)
}
