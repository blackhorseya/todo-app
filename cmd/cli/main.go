package main

import (
	"flag"

	"github.com/blackhorseya/todo-app/internal/app/commands"
	"github.com/blackhorseya/todo-app/internal/pkg/config"
	"github.com/blackhorseya/todo-app/internal/pkg/utils/exit"
	"github.com/spf13/cobra"
)

var cfgPath = flag.String("c", "configs/app.yaml", "set config file path")

func init() {
	flag.Parse()

	cobra.OnInitialize(initConfig)
}

func main() {
	err := commands.Execute()
	if err != nil {
		exit.Er(err)
	}
}

func initConfig() {
	_, err := config.NewConfig(*cfgPath)
	if err != nil {
		exit.Er(err)
	}
}
