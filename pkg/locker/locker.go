package locker

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/satori/go.uuid"
)

// Locker is a type that holds all data about a container/Locker
type Locker struct {
	Name       string
	Env        []string
	PID        int
	ID         string
	Command    []string
	Process    *exec.Cmd
	Filesystem *Filesystem
}

// NewLocker is a factory for lockers, it should be used like this:
// 	locker = NewLocker(name, args[1:]}
//
// We really only need a command to make a process run in a locker. Everything else can be generated.
func NewLocker(name string, command []string) Locker {
	if name == "" {
		name = "gnarly_narwhal"
	}
	id := uuid.NewV4().String()

	process := exec.Command(command[0], command[1:]...)

	process.Stdin = os.Stdin
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr
	process.Env = []string{"PATH=/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin"}

	process.Run()
	pid := process.Process.Pid

	fs := NewFilesystem(id)

	return Locker{
		Name:       name,
		Env:        []string{},
		PID:        pid,
		ID:         id,
		Command:    command,
		Process:    process,
		Filesystem: fs,
	}
}

// Run takes a locker and runs the contents of args in an isolated environment
func (locker *Locker) Run() {
	fmt.Printf("Command: %v\n", strings.Join(locker.Command, " "))

	if err := locker.Process.Run(); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	locker.PID = locker.Process.Process.Pid
	fmt.Println("Container ran as PID: %i", locker.PID)
}
