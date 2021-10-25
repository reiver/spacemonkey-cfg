package datadir

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func Path() (string, error) {

	var datapath string
	{
		if "" == datapath {
			datapath = xdg_data_home()
		}

		if "" == datapath {
			datapath = xdg_data_dir()
		}

		if "" == datapath {
			datapath = default_datapath()
		}

	}

	if "" == datapath {
		return "", errors.New("could not figure out the path to the user's base data-dir")
	}

	return datapath, nil
}

func xdg_data_home() string {
	const name string = "XDG_DATA_HOME"

	var datapath string
	{
		datapath = os.Getenv(name)
	}
	datapath = filepath.Clean(datapath)
	if !filepath.IsAbs(datapath) {
		return ""
	}

	return datapath
}

func xdg_data_dir() string {
	const name string = "XDG_DATA_DIRS"

	paths := os.Getenv(name)
	if "" == paths {
		return ""
	}

	index := strings.Index(paths, ":")
	if index < 0 {
		return paths
	}

	var datapath string
	{
		datapath = paths[:index]
	}
	datapath = filepath.Clean(datapath)
	if !filepath.IsAbs(datapath) {
		return ""
	}

	return datapath
}

func default_datapath() string {
	var homepath string
	{
		path, err := os.UserHomeDir()
		if nil != err {
			// could not get the user's home-path when trying to figure out the user's default base data-path
			return ""
		}

		homepath = path
	}
	if "" == homepath {
		// could not get the user's home-path when trying to figure out the user's default base data-path
		return ""
	}

	var datapath string
	{
		const datadir string = ".local/share"

		datapath = filepath.Join(homepath, datadir)
	}
	datapath = filepath.Clean(datapath)
	if !filepath.IsAbs(datapath) {
		return ""
	}

	return datapath

}
