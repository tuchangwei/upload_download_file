package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Local struct {
	basePath string
}
//path is a relative path, basePath will be appended
func (l *Local) Save(path string, contents io.Reader) error {
	absPath := filepath.Join(l.basePath, path)

	// get the directory and make sure it exists
	d := filepath.Dir(absPath)
	fmt.Println("dir: ", d)
	err := os.MkdirAll(d, os.ModePerm)
	if err != nil {
		return fmt.Errorf("can't create dir, error: %q", err)
	}
	_, err = os.Stat(absPath)
	if err == nil {
		err = os.Remove(absPath)
		if err != nil {
			return fmt.Errorf("can't delete file, error: %q", err)
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("can't get file info, error: %q", err)
	}
	fmt.Println("error: ", err)
	f, err := os.Create(absPath)
	if err != nil {
		return fmt.Errorf("can't create file, error: %q", err)
	}
	defer f.Close()
	_, err = io.Copy(f, contents)
	if err != nil {
		return fmt.Errorf("can't write content to file, error: %q", err)
	}
	return nil
}

func NewLocal(basePath string) (*Local, error)  {
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}
	return &Local{ basePath: p}, nil
}
