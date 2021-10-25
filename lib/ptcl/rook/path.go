package rook

import (
	"github.com/reiver/spacemonkey-cfg/lib/datadir"

	"errors"
	"fmt"
	"path/filepath"
)

const (
	ptcldir = "rook-protocol"
)

func DataPath() (string, error) {

	var basedatapath string
	{
		var err error

		basedatapath, err = datadir.Path()
		if nil != err {
			return "", err
		}
	}
	basedatapath = filepath.Clean(basedatapath)
	if !filepath.IsAbs(basedatapath) {
		return "", errors.New("the user's base data-path is not an absolute path")
	}

	var datapath string
	{
		datapath = filepath.Join(basedatapath, ptcldir)
	}
	datapath = filepath.Clean(datapath)
	if !filepath.IsAbs(datapath) {
		return "", fmt.Errorf("the user-application's data-path for %s for the user is not an absolute path", ptcldir)
	}

	if "" == datapath {
		return "", fmt.Errorf("could not figure out the user-application's data-path for %s for the user", ptcldir)
	}

	return datapath, nil
}
