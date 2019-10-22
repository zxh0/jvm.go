package vmutils

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"path/filepath"
)

type JModFile struct {
	absPath string
	r       *zip.Reader
}

func NewJModFile(path string) (*JModFile, error) {
	if absPath, err := filepath.Abs(path); err != nil {
		return nil, err
	} else {
		return &JModFile{absPath: absPath}, nil
	}
}

func (mf *JModFile) AbsPath() string {
	return mf.absPath
}
func (mf *JModFile) IsOpen() bool {
	return mf.r != nil
}

func (mf *JModFile) Open() error {
	r, err := OpenJModReader(mf.absPath)
	if err == nil {
		mf.r = r
	}
	return err
}

func (mf *JModFile) ReadFile(filename string) ([]byte, error) {
	return readFileInZip(mf.r, filename)
}

func OpenJModReader(filename string) (*zip.Reader, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	data = data[4:] // skip 0x4a4d0100
	return zip.NewReader(bytes.NewReader(data), int64(len(data)))
}
