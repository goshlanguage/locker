package locker

import (
	"os"
	"path/filepath"
)

const basedir = "/usr/local/locker/filesystems/"

// Filesystem is a struct we use to keep track of a Filesystem's space on the host
type Filesystem struct {
	ID   string
	Path string
}

// NewFilesystem creates a blank dir in locker root under filesystems to provide FS isolation
func NewFilesystem(id string) *Filesystem {
	lockerPath := basedir + id + "/rootfs"

	if _, err := os.Stat(lockerPath); os.IsNotExist(err) {
		os.Mkdir(lockerPath, 0755)
	} else if err != nil {
		panic(err)
	}

	return &Filesystem{
		ID:   id,
		Path: lockerPath,
	}
}

// RemoveFilesystem cleans all the data off of the host
func (fs *Filesystem) RemoveFilesystem() (bool, error) {
	dir := basedir + fs.ID + "/rootfs"
	d, err := os.Open(dir)
	if err != nil {
		return false, err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return false, err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return false, err
		}
	}
	return true, nil

}
