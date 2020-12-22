package tasks

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/blackhorseya/todo-app/internal/pkg/utils/exit"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "Print all tasks",
		Run: func(cmd *cobra.Command, args []string) {
			url := fmt.Sprintf("%v/api/v1/tasks", viper.Get("api.endpoint"))
			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				exit.Er(err)
			}

			c := new(http.Client)
			resp, err := c.Do(req)
			if err != nil {
				exit.Er(err)
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				exit.Er(err)
			}

			fmt.Println(string(body))
		},
	}
)
