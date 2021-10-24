package arg

import (
	"flag"
)

var (
	Values []string
)

var (
	EOL    bool
	Verbose bool
)

func init() {
	flag.BoolVar(&EOL,     "eol", false, "output trailing end-of-line character(s)")
	flag.BoolVar(&Verbose, "v",   false, "verbose logs outputted")

	flag.Parse()

	Values = flag.Args()
}
