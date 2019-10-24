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

func OpenJModFile(path string) (*JModFile, error) {
	if jmodFile, err := NewJModFile(path); err != nil {
		return nil, err
	} else if err := jmodFile.Open(); err != nil {
		return nil, err
	} else {
		return jmodFile, nil
	}
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

func (mf *JModFile) Close() {
	mf.r = nil
}

func (mf *JModFile) Open() error {
	data, err := ioutil.ReadFile(mf.absPath)
	if err != nil {
		return err
	}

	data = data[4:] // skip 0x4a4d0100
	r, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err == nil {
		mf.r = r
	}
	return err
}

func (mf *JModFile) ReadFile(filename string) ([]byte, error) {
	return readFileInZip(mf.r, filename)
}

func (mf *JModFile) ListFiles() []string {
	files := make([]string, 0, 100)
	for _, f := range mf.r.File {
		files = append(files, f.Name)
	}
	return files
}
