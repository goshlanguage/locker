package locker

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/satori/go.uuid"
)

// LockerOpts holds the configuration options to build a Locker
type LockerOpts struct {
	Name     string
	Env      []string
	Command  []string
	Hostname string
}

// Locker holds all data about a container/Locker
type Locker struct {
	PID        int
	ID         string
	Process    *exec.Cmd
	Filesystem *Filesystem
	Config     *LockerOpts
}

// NewLocker is a builder for lockers. If you're unfamiliar with the builder patterns, see:
//		https://github.com/tmrts/go-patterns/blob/master/creational/builder.md
//
// NewLocker spawns a new child process in order to apply namespaces and cgroups
// 		to the process. This is known as "fork() and exec()". SysProcAttr
//
//
func (l *LockerOpts) Build() Locker {

	// Really, we should be Forking a process to run our command in
	// However, then we don't have access to our Locker struct
	// process := exec.Command("/proc/self/exe", append([]string{"fork"}, l.Command[:]...)...)
	process := exec.Command(strings.Join(l.Command, " "))

	process.Stdin = os.Stdin
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr
	process.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	syscall.Sethostname([]byte(l.Hostname))

	id := uuid.NewV4().String()
	fs := NewFilesystem(id)

	syscall.Mount(fs.Path, "/", "", syscall.MS_BIND, "")
	os.MkdirAll(fs.Path, 0700)
	// When we have a mechanism for copying a base image, we can provide filesystem isolation
	// 	initially by Chrooting the filesystem. Until then, let's chdir into the directory
	// syscall.Chroot(fs.Path)
	syscall.Chdir(fs.Path)

	return Locker{
		PID:        0,
		ID:         id,
		Process:    process,
		Filesystem: fs,
		Config:     l,
	}
}

// Run takes a locker and runs the contents of args in an isolated environment
func (locker *Locker) Run() {
	fmt.Printf("Command: %v\n", strings.Join(locker.Config.Command, " "))

	if err := locker.Process.Run(); err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
	locker.PID = locker.Process.Process.Pid
	fmt.Printf("Container %s ran as PID: %i\n", locker.ID, locker.PID)

	locker.Filesystem.RemoveFilesystem()
}

// Fork should create an isolated process for us forked from our Locker process
func Fork(args []string) {

	cmd := exec.Command(args[0])
	if len(args) > 1 {
		cmd = exec.Command(args[0], strings.Join(args[1:], " "))
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = []string{"PATH=/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin"}
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	// emulate UTS namespace
	syscall.Sethostname([]byte("locker"))

	// implement filesystem namespace
	// syscall.Mount(Filesystem.Path, "/", "", syscall.MS_BIND, "")
	// os.MkdirAll(Filesystem.Path, 0700)
	os.Chdir("/")

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
}
