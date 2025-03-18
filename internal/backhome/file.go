package backhome

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/fernandogiovanini/backhome/internal/logger"
	"github.com/fernandogiovanini/backhome/internal/utils"
)

// File represents a file to be copied
// The path is the absolute path of the file
type File struct {
	path string
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

func NewFile(filename string) (*File, error) {
	path, err := utils.ResolvePath(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve %s: %w", filename, err)
	}

	return &File{
		path: path,
	}, nil
}

// CopyTo copies the file to the destination
// The destination is created based on the [Local.path]
func (file *File) CopyTo(local *Local) error {
	if local == nil {
		return errors.New("nil pointer: Local is not initialized")
	}

	destination, err := file.NewDestination(local)
	if err != nil {
		return fmt.Errorf("failed to create destination for %s: %w", file.path, err)
	}

	srcFile, err := os.Open(file.path)
	if err != nil {
		return fmt.Errorf("failed to open %s for reading: %v", file.path, err)
	}
	defer srcFile.Close()

	fileinfo, err := os.Stat(file.path)
	if err != nil {
		return fmt.Errorf("failed to read source %s stats: %w", file.path, err)
	}

	dstFile, err := os.OpenFile(destination.path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, fileinfo.Mode())
	if err != nil {
		return fmt.Errorf("failed to open %s for writing: %w", destination.path, err)
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return fmt.Errorf("failed to copy %s to %s: %w", srcFile.Name(), dstFile.Name(), err)
	}

	if er := dstFile.Sync(); er != nil {
		return fmt.Errorf("failed to sync %s: %w", dstFile.Name(), er)
	}

	logger.Debug("copied %s to %s", file.path, destination.path)

	return nil
}

func (file *File) NewDestination(local *Local) (*Destination, error) {
	if local == nil {
		return nil, errors.New("nil pointer: Local is not initialized")
	}

	if local.GetPath() == "" {
		return nil, errors.New("invalid destination: local.BasePath is empty")
	}

	if file.path == "" {
		return nil, errors.New("invalid destination: item.Path is empty")
	}

	filename, err := filepath.Abs(filepath.Join(local.GetPath(), file.path))
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for destination files: %w", err)
	}

	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create destination directory %s: %w", dir, err)
	}

	return &Destination{
		path: filename,
	}, nil
}

func NewFileList(filenames []string) (*FileList, error) {
	fileList := &FileList{
		Files: make([]*File, 0),
	}
	for _, filename := range filenames {
		file, err := NewFile(filename)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve path of file %s: %w", filename, err)
		}
		fileList.Files = append(fileList.Files, file)
	}

	return fileList, nil
}

func (fileList *FileList) CopyTo(local *Local) error {
	if local == nil {
		return errors.New("nil pointer: Local is not initialized")
	}

	for i, file := range fileList.Files {
		fmt.Printf("%3d/%-3d %-50s\t", i+1, fileList.Count(), file.path)
		if err := file.CopyTo(local); err != nil {
			fmt.Println("FAILED")
			return fmt.Errorf("failed to copy %s to %s: %w", file.path, local.GetPath(), err)
		}
		fmt.Printf("OK (%.2f%%)\n", float64(i+1)/float64(fileList.Count())*100)
	}

	logger.Debug("copied %d files to %s", len(fileList.Files), local.GetPath())

	return nil
}
func (fileList *FileList) Count() int {
	return len(fileList.Files)
}
