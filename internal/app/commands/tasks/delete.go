package tasks

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "delete a task",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("required id of task")
			}

			_, err := uuid.Parse(args[0])
			if err != nil {
				return err
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			url := fmt.Sprintf("%v/api/v1/tasks/%s", viper.Get("api.endpoint"), args[0])
			req, err := http.NewRequest(http.MethodDelete, url, nil)
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
