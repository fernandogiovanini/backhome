package backhome

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/fernandogiovanini/backhome/logger"
	"github.com/fernandogiovanini/backhome/utils"
)

// Item represents a file to be copied
// The filename is the absolute path of the file
type Item struct {
	Filename string
}

// Destination represents the destination of a file copy
// The filename and dir are created based on [Local.BasePath]
type Destination struct {
	Filename string
	Dir      string
}

func NewItem(filename string) (*Item, error) {
	path, err := utils.ResolvePath(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve path of target %s: %w", filename, err)
	}

	return &Item{
		Filename: path,
	}, nil
}

// CopyTo copies the file to the destination
// The destination is created based on the [Local.BasePath]
func (item *Item) CopyTo(local *Local) error {
	if item == nil {
		return errors.New("nil pointer: Item is not initialized")
	}

	if local == nil {
		return errors.New("nil pointer: Local is not initialized")
	}

	destination, err := item.NewDestination(local)
	if err != nil {
		return fmt.Errorf("failed to create destination for %s: %w", item.Filename, err)
	}

	srcFile, err := os.Open(item.Filename)
	if err != nil {
		return fmt.Errorf("failed to open %s for reading: %w", item.Filename, err)
	}
	defer srcFile.Close()

	fileinfo, err := os.Stat(item.Filename)
	if err != nil {
		return fmt.Errorf("failed to read source %s stats: %w", item.Filename, err)
	}

	dstFile, err := os.OpenFile(destination.Filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, fileinfo.Mode())
	if err != nil {
		return fmt.Errorf("failed to open %s for writing: %w", destination.Filename, err)
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return fmt.Errorf("failed to copy %s to %s: %w", srcFile.Name(), dstFile.Name(), err)
	}
	dstFile.Sync()

	logger.Debug("copied %s to %s", item.Filename, destination.Filename)

	return nil
}

func (item *Item) NewDestination(local *Local) (*Destination, error) {
	if item == nil {
		return nil, errors.New("nil pointer: Item is not initialized")
	}

	if local == nil {
		return nil, errors.New("nil pointer: Local is not initialized")
	}

	if local.BasePath == "" {
		return nil, errors.New("invalid destination: local.BasePath is empty")
	}

	if item.Filename == "" {
		return nil, errors.New("invalid destination: item.Path is empty")
	}

	filename, err := filepath.Abs(filepath.Join(local.BasePath, item.Filename))
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for destination files: %w", err)
	}

	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create destination directory %s: %w", dir, err)
	}

	return &Destination{
		Filename: filename,
		Dir:      filepath.Dir(filename),
	}, nil
}

// ItemList represents a list of files to be copied
type ItemList struct {
	Item []*Item
}

func NewItemList(filenames []string) (*ItemList, error) {
	itemList := &ItemList{
		Item: make([]*Item, 0),
	}
	for _, filename := range filenames {
		item, err := NewItem(filename)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve path of target %s: %w", filename, err)
		}
		itemList.Item = append(itemList.Item, item)
	}
	return itemList, nil
}

func (itemList *ItemList) CopyTo(local *Local) error {
	if itemList == nil {
		return errors.New("nil pointer: ItemList is not initialized")
	}

	if local == nil {
		return errors.New("nil pointer: Local is not initialized")
	}

	for _, item := range itemList.Item {
		if err := item.CopyTo(local); err != nil {
			return fmt.Errorf("failed to copy %s to %s: %w", item.Filename, local.BasePath, err)
		}
	}

	logger.Debug("copied %d items to %s", len(itemList.Item), local.BasePath)

	return nil
}
