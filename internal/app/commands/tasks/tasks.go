package tasks

import (
	"github.com/spf13/cobra"
)

var (
	// Cmd is root command of tasks
	Cmd = &cobra.Command{
		Use:   "tasks",
		Short: "Tasks command",
	}
)

func init() {
	Cmd.AddCommand(listCmd)
	Cmd.AddCommand(deleteCmd)
	Cmd.AddCommand(addCmd)
	Cmd.AddCommand(completeCmd)
	Cmd.AddCommand(incompleteCmd)
}
