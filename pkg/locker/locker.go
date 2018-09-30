package locker

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

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

	process := exec.Command("/proc/self/exe", append([]string{"child"}, command[1:]...)...)

	process.Stdin = os.Stdin
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr
	process.Env = []string{"PATH=/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin"}
	process.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

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

// SpawnChild should create an isolated process for us forked from our Locker process
func (locker *Locker) SpawnChild() {
	syscall.Mount(locker.Filesystem.Path, locker.Filesystem.Path, "", syscall.MS_BIND, "")
	os.MkdirAll(locker.Filesystem.Path, 0700)
	os.Chdir("/")

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
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
