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

func OpenZipFile(path string) (*ZipFile, error) {
	if zipFile, err := NewZipFile(path); err != nil {
		return nil, err
	} else if err := zipFile.Open(); err != nil {
		return nil, err
	} else {
		return zipFile, nil
	}
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

func (zf *ZipFile) HasFile(filename string) bool {
	for _, f := range zf.rc.File {
		if f.Name == filename {
			return true
		}
	}
	return false
}

func (zf *ZipFile) ReadFile(filename string) ([]byte, error) {
	return readFileInZip(&zf.rc.Reader, filename)
}

func readFileInZip(r *zip.Reader, filename string) ([]byte, error) {
	for _, f := range r.File {
		if f.Name == filename {
			return readFileInZip0(f)
		}
	}
	return nil, io.EOF // TODO
}

func readFileInZip0(file *zip.File) ([]byte, error) {
	if rc, err := file.Open(); err != nil {
		return nil, err
	} else {
		defer rc.Close()
		return ioutil.ReadAll(rc)
	}
}
