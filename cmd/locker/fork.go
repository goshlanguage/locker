package main // import "github.com/ryanhartje/locker/cmd/locker"

import (
	"github.com/ryanhartje/locker/pkg/locker"
	"github.com/spf13/cobra"
)

func newForkCmd(args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fork",
		Short: "fork a process in the runc runtime.",
		Long:  "fork a process in the runc runtime.",
		Run: func(cmd *cobra.Command, args []string) {
			locker.Fork(args)
		},
	}
	return cmd
}
