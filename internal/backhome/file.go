//go:generate mockery --all --case snake

package backhome

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/fernandogiovanini/backhome/internal/filesystem"
	"github.com/fernandogiovanini/backhome/internal/utils"
)

type FileOperations interface {
	CopyTo(local *Local) error
	NewDestination(local *Local) (*Destination, error)
	Path() string
}

// File represents a file to be copied
// The path is the absolute path of the file
type File struct {
	filesystem filesystem.FileSystem
	path       string
}

type FileListOperations interface {
	CopyTo(local *Local) error
	Count() int
}

// FileList represents a list of files to be copied
type FileList struct {
	Files []*File
}

// Destination represents the destination of a file copy
// The path is based on [Local.path]
type Destination struct {
	path string
}

type DestinationOperations interface {
	Path() string
}

func NewFile(filename string, filesystem filesystem.FileSystem) (*File, error) {
	path, err := utils.ResolvePath(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve %s: %w", filename, err)
	}

	return &File{
		path:       path,
		filesystem: filesystem,
	}, nil
}

// CopyTo copies the file to the destination
// The destination is created based on the [Local.path]
func (f File) CopyTo(local *Local) error {
	destination, err := f.NewDestination(local)
	if err != nil {
		return fmt.Errorf("failed to create destination for %s: %w", f.Path(), err)
	}

	return f.copyFile(*destination)
}

func (f File) copyFile(destination Destination) error {
	src, err := f.filesystem.Open(f.Path())
	if err != nil {
		return fmt.Errorf("failed to open %s for reading: %v", f.Path(), err)
	}
	defer src.Close()

	fileinfo, err := f.filesystem.Stat(f.Path())
	if err != nil {
		return fmt.Errorf("failed to read source %s stats: %w", f.Path(), err)
	}

	dst, err := f.filesystem.OpenFile(destination.Path(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, fileinfo.Mode())
	if err != nil {
		return fmt.Errorf("failed to open %s for writing: %w", destination.Path(), err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return fmt.Errorf("failed to copy %s to %s: %w", src.Name(), dst.Name(), err)
	}

	return dst.Sync()
}

func (f File) Path() string {
	return f.path
}

func (f File) NewDestination(local *Local) (*Destination, error) {
	if local.Path() == "" {
		return nil, errors.New("invalid destination: local.BasePath is empty")
	}

	if f.Path() == "" {
		return nil, errors.New("invalid destination: file.path is empty")
	}

	filename, err := filepath.Abs(filepath.Join(local.Path(), f.Path()))
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for destination files: %w", err)
	}

	dir := filepath.Dir(filename)
	if err := f.filesystem.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create destination directory %s: %w", dir, err)
	}

	return &Destination{
		path: filename,
	}, nil
}

func NewFileList(filenames []string, filesystem filesystem.FileSystem) (*FileList, error) {
	fileList := &FileList{
		Files: make([]*File, 0),
	}
	for _, filename := range filenames {
		file, err := NewFile(filename, filesystem)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve path of file %s: %w", filename, err)
		}
		fileList.Files = append(fileList.Files, file)
	}

	return fileList, nil
}

func (fl FileList) CopyTo(local *Local, writer io.Writer) error {
	for i, file := range fl.Files {
		fmt.Fprintf(writer, "%3d/%-3d %-50s\t", i+1, fl.Count(), file.Path())
		if err := file.CopyTo(local); err != nil {
			fmt.Fprintf(writer, "FAILED")
			return fmt.Errorf("failed to copy %s to %s: %w", file.Path(), local.Path(), err)
		}
		fmt.Fprintf(writer, "OK (%.2f%%)\n", float64(i+1)/float64(fl.Count())*100)
	}

	return nil
}

func (fl FileList) Count() int {
	return len(fl.Files)
}

func (d *Destination) Path() string {
	return d.path
}
