package handler

import (
	"fmt"

	"github.com/spf13/cobra"

	"todo/internal/service"
)

func newListCmd(svc service.TodoService) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all todos",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			todos, err := svc.List(cmd.Context())
			if err != nil {
				return err
			}

			out := cmd.OutOrStdout()
			if len(todos) == 0 {
				fmt.Fprintln(out, "No todos yet.")
				return nil
			}

			for _, t := range todos {
				mark := " "
				if t.Done {
					mark = "x"
				}
				fmt.Fprintf(out, "[%s] %-3d %s\n", mark, t.ID, t.Title)
			}
			return nil
		},
	}
}
