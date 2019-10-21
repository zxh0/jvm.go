package vmutils

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"path/filepath"
)

type ZipFile struct {
	absPath string
	rc      *zip.ReadCloser
}

func NewZipFile(path string) (*ZipFile, error) {
	if absPath, err := filepath.Abs(path); err != nil {
		return nil, err
	} else {
		return &ZipFile{absPath: absPath}, nil
	}
}

func (zf *ZipFile) AbsPath() string {
	return zf.absPath
}
func (zf *ZipFile) IsOpen() bool {
	return zf.rc != nil
}

func (zf *ZipFile) Open() error {
	rc, err := zip.OpenReader(zf.absPath)
	if err == nil {
		zf.rc = rc
	}
	return err
}

func (zf *ZipFile) Close() error {
	if zf.rc != nil {
		return zf.rc.Close()
	}
	return nil
}

func (zf *ZipFile) ReadFile(filename string) ([]byte, error) {
	for _, f := range zf.rc.File {
		if f.Name == filename {
			return readFileInZip(f)
		}
	}
	return nil, io.EOF // TODO
}

func readFileInZip(file *zip.File) ([]byte, error) {
	if rc, err := file.Open(); err != nil {
		return nil, err
	} else {
		defer rc.Close()
		return ioutil.ReadAll(rc)
	}
}
