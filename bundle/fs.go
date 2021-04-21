package bundle

import (
	"errors"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const suffix = ".xml"

type Dir string

// mapDirOpenError maps the provided non-nil error from opening name
// to a possibly better non-nil error. In particular, it turns OS-specific errors
// about opening files in non-directories into os.ErrNotExist. See Issue 18984.
func mapDirOpenError(originalErr error, name string) error {
	if os.IsNotExist(originalErr) || os.IsPermission(originalErr) {
		return originalErr
	}
	
	parts := strings.Split(name, string(filepath.Separator))
	for i := range parts {
		if parts[i] == "" {
			continue
		}
		fi, err := os.Stat(strings.Join(parts[:i+1], string(filepath.Separator)))
		if err != nil {
			return originalErr
		}
		if !fi.IsDir() {
			return os.ErrNotExist
		}
	}
	return originalErr
}

// Open implements FileSystem using os.Open, opening files for reading rooted
// and relative to the directory d.
func (p Dir) Open(name string) (http.File, error) {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) {
		return nil, errors.New("http: invalid character in file path")
	}
	dir := string(p)
	if dir == "" {
		dir = "."
	}
	fullName := filepath.Join(dir, filepath.FromSlash(path.Clean("/"+name)))
	f, err := os.Open(fullName)
	if err != nil {
		return nil, mapDirOpenError(err, fullName)
	}
	return &File{f: f}, nil
}

type File struct {
	f http.File
}

func (p File) Close() error {
	return p.f.Close()
}

func (p File) Read(len []byte) (int, error) {
	return p.f.Read(len)
}

func (p File) Seek(offset int64, whence int) (int64, error) {
	return p.f.Seek(offset, whence)
}

func (p File) Readdir(count int) ([]os.FileInfo, error) {
	infos, err := p.f.Readdir(count)
	if err != nil {
		return infos, err
	}
	var res []os.FileInfo
	for _, v := range infos {
		if v.IsDir() || strings.HasSuffix(v.Name(), suffix) {
			res = append(res, v)
		}
	}
	return res, nil
}

func (p File) Stat() (os.FileInfo, error) {
	return p.f.Stat()
}
