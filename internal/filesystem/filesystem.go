//go:generate mockery --all --case snake

package filesystem

import "os"

type FileSystem interface {
	Open(name string) (*os.File, error)
	Stat(name string) (os.FileInfo, error)
	OpenFile(name string, flag int, perm os.FileMode) (*os.File, error)
	MkdirAll(path string, perm os.FileMode) error
	RemoveAll(path string) error
	IsNotExist(err error) bool
	IsPermission(err error) bool
}

type OSFileSystem struct{}

func NewFileSystem() *OSFileSystem {
	return new(OSFileSystem)
}

func (l *OSFileSystem) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (l *OSFileSystem) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

func (l *OSFileSystem) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (l *OSFileSystem) RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func (l *OSFileSystem) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

func (l *OSFileSystem) IsPermission(err error) bool {
	return os.IsPermission(err)
}

func (l *OSFileSystem) Open(name string) (*os.File, error) {
	return os.Open(name)
}
