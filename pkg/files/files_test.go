package files_test

import (
	"io/fs"

	"github.com/charmbracelet/log"
)

// TODO: Add file generator tests

type NullFileGenerator struct{}

func (nfg NullFileGenerator) WriteFile(name string, data []byte, perm fs.FileMode) error {
	log.Debug("Creating file", name)
	return nil
}

func (nfg NullFileGenerator) MkdirAll(path string, perm fs.FileMode) error {
	log.Debug("Creating directory", path)
	return nil
}

func (nfg NullFileGenerator) Remove(name string) error {
	log.Debug("Removing", name)
	return nil
}

func (nfg NullFileGenerator) RemoveAll(pathname string) error {
	log.Debug("Removin", pathname)
	return nil
}

func (NullFileGenerator) ReadFile(name string) ([]byte, error) {
	return []byte{}, nil
}

func (NullFileGenerator) ReadDir(name string) ([]fs.DirEntry, error) {
	return []fs.DirEntry{}, nil
}
