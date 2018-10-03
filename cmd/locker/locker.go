package main // import "github.com/ryanhartje/locker/cmd/locker"

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var usage string = `
locker <command> <flags> <image> <args>

locker is a simple lightweight container runtime. This is mostly for example, please don't use this in production.

locker provides the following namespacing:
- Unix TimeSharing system [https://people.eecs.berkeley.edu/~brewer/cs262/unix.pdf]
-- Filesystem 
--- Process IDs
--- Users
- IPC
- Networking
`

func newRootCmd(args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "locker",
		Short: "locker is a lightweight container runtime.",
		Long:  "locker is a lightweight container runtime. Do not use this in production.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(usage)
		},
	}

	return cmd
}

func main() {
	rootCmd := newRootCmd(os.Args)
	rootCmd.AddCommand(newRunCmd(os.Args))
	rootCmd.AddCommand(newVersionCmd(os.Args))
	rootCmd.AddCommand(newForkCmd(os.Args))
	rootCmd.Execute()
}
