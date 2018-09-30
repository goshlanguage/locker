package locker

import (
	"fmt"
	"os"
	"testing"

	"github.com/satori/go.uuid"
)

var id = uuid.NewV4().String()
var expected = "/usr/local/locker/filesystems/" + id + "/rootfs"

func TestNewFilesystem(t *testing.T) {
	fs := NewFilesystem(id)
	if fs.Path != expected {
		fmt.Printf(
			"Failed to get new filesystem. Tried to create and expected: %s\nGot: %s\n",
			expected,
			fs.Path,
		)
		t.Fail()
	}
}

func TestRemoveFilesystem(t *testing.T) {
	fs := NewFilesystem(id)
	fs.RemoveFilesystem()

	catch := false
	if _, err := os.Stat(expected); os.IsNotExist(err) {
		catch = true
	} else if err != nil {
		t.Fail()
	}
	if !catch {
		t.Fail()
	}
}
