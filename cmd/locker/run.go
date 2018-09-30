package main // import "github.com/ryanhartje/locker/cmd/locker"

import (
	"fmt"

	"github.com/ryanhartje/locker/pkg/locker"
	"github.com/spf13/cobra"
)

func newRunCmd(args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "runs a process in a container",
		Long:  "runs a process in a container",
		Run: func(cmd *cobra.Command, args []string) {

			if !(len(args) > 0) {
				panic("Didn't get enough arguments for run command. Please provide a command to run, eg: locker run echo \"hello world\"")
			}

			l := locker.NewLocker("", args)

			fmt.Printf("Exiting from: %s on PID: %d\n", l.ID, l.PID)

		},
	}

	return cmd
}
