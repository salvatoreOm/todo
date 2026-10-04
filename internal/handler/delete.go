package handler

import (
	"fmt"

	"github.com/spf13/cobra"

	"todo/internal/service"
)

func newDeleteCmd(svc service.TodoService) *cobra.Command {
	return &cobra.Command{
		Use:     "delete <id>",
		Aliases: []string{"rm"},
		Short:   "Delete a todo",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}

			if err := svc.Delete(cmd.Context(), id); err != nil {
				return err
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Deleted todo #%d\n", id)
			return nil
		},
	}
}
