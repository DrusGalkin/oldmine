package pkg

import (
	"os"
)

func MustLoadMkDir(path string) {
	if path == "" {
		path = getDefaultPath()
	}

	os.Mkdir(path, os.ModeDir)
}

func getDefaultPath() string {
	return "/var/lib/skin-system/uploads"
}
