package main // import "github.com/ryanhartje/locker/cmd/locker"

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newVersionCmd(args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "version",
		Long:  "Get locker version",
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("0.0.5")
		},
	}

	return cmd
}
