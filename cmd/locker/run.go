package main // import "github.com/ryanhartje/locker/cmd/locker"

import (
	"github.com/ryanhartje/locker/pkg/locker"
	"github.com/spf13/cobra"
)

type runCmd struct {
	name     string
	env      []string
	command  []string
	hostname string
}

func newRunCmd(args []string) *cobra.Command {

	run := &runCmd{
		name:     "salty_simon",
		hostname: "locker",
	}

	cmd := &cobra.Command{
		Use:   "run",
		Short: "runs a process in a container",
		Long:  "runs a process in a container",
		Run: func(cmd *cobra.Command, args []string) {

			if !(len(args) > 0) {
				panic("Didn't get enough arguments for run command. Please provide a command to run, eg: locker run echo \"hello world\"")
			}
			run.command = args
			run.run()

			return
		},
	}

	f := cmd.Flags()
	f.StringVar(&run.name, "name", "", "Manually name your locker")
	f.StringArrayVar(&run.env, "env", []string{}, "Set environment variables eg: foo=bar")
	f.StringVar(&run.hostname, "hostname", "", "Set the hostname for your locker")

	return cmd
}

func (r *runCmd) run() error {
	config := locker.LockerOpts{
		Name:     r.name,
		Env:      r.env,
		Hostname: r.hostname,
		Command:  r.command,
	}

	locker := config.Build()
	locker.Run()

	return nil
}
