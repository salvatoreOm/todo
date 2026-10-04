package handler

import (
	"fmt"

	"github.com/spf13/cobra"

	"todo/internal/service"
)

func newToggleCmd(svc service.TodoService) *cobra.Command {
	return &cobra.Command{
		Use:   "toggle <id>",
		Short: "Mark a todo as done, or as not done if it already is",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}

			todo, err := svc.Toggle(cmd.Context(), id)
			if err != nil {
				return err
			}

			state := "not done"
			if todo.Done {
				state = "done"
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Todo #%d marked as %s\n", todo.ID, state)
			return nil
		},
	}
}
