package tasks

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/blackhorseya/todo-app/internal/app/entities"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "add a task",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("required title of task")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			url := fmt.Sprintf("%v/api/v1/tasks", viper.Get("api.endpoint"))
			newTask := &entities.Task{
				Title: args[0],
			}
			data, err := json.Marshal(newTask)
			if err != nil {
				return err
			}

			req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
			if err != nil {
				return err
			}

			c := &http.Client{}
			resp, err := c.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			fmt.Println(string(body))

			return nil
		},
	}
)
