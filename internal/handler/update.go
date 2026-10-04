package handler

import (
	"fmt"

	"github.com/spf13/cobra"

	"todo/internal/service"
)

func newUpdateCmd(svc service.TodoService) *cobra.Command {
	return &cobra.Command{
		Use:     "update <id> <title>",
		Short:   "Change the title of a todo",
		Example: `  todo update 1 "buy oat milk"`,
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}

			todo, err := svc.UpdateTitle(cmd.Context(), id, args[1])
			if err != nil {
				return err
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Updated todo #%d: %s\n", todo.ID, todo.Title)
			return nil
		},
	}
}
