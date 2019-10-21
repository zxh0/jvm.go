package vmutils

import (
	"io/ioutil"
	"path/filepath"
)

type Dir struct {
	absPath string
}

func NewDir(path string) (*Dir, error) {
	if absPath, err := filepath.Abs(path); err != nil {
		return nil, err
	} else {
		return &Dir{absPath: absPath}, nil
	}
}

func (dir *Dir) AbsPath() string {
	return dir.absPath
}

func (dir *Dir) ReadFile(filename string) ([]byte, error) {
	absFilename := filepath.Join(dir.absPath, filename)
	if data, err := ioutil.ReadFile(absFilename); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}
