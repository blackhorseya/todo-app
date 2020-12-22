package commands

import (
	"github.com/blackhorseya/todo-app/internal/app/commands/tasks"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "todo",
		Long: `

████████╗ ██████╗ ██████╗  ██████╗     ██╗     ██╗███████╗████████╗
╚══██╔══╝██╔═══██╗██╔══██╗██╔═══██╗    ██║     ██║██╔════╝╚══██╔══╝
   ██║   ██║   ██║██║  ██║██║   ██║    ██║     ██║███████╗   ██║   
   ██║   ██║   ██║██║  ██║██║   ██║    ██║     ██║╚════██║   ██║   
   ██║   ╚██████╔╝██████╔╝╚██████╔╝    ███████╗██║███████║   ██║   
   ╚═╝    ╚═════╝ ╚═════╝  ╚═════╝     ╚══════╝╚═╝╚══════╝   ╚═╝
`,
	}
)

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(tasks.Cmd)
}
