package handler

import (
	"fmt"

	"github.com/spf13/cobra"

	"todo/internal/service"
)

func newAddCmd(svc service.TodoService) *cobra.Command {
	return &cobra.Command{
		Use:     "add <title>",
		Short:   "Add a new todo",
		Example: `  todo add "buy milk"`,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			todo, err := svc.Add(cmd.Context(), args[0])
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Added todo #%d: %s\n", todo.ID, todo.Title)
			return nil
		},
	}
}
