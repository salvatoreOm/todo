package handler

import (
	"github.com/spf13/cobra"

	"todo/internal/service"
)

// NewRootCmd builds the root "todo" command and attaches all subcommands.
func NewRootCmd(svc service.TodoService) *cobra.Command {
	root := &cobra.Command{
		Use:           "todo",
		Short:         "A simple todo app for the command line",
		SilenceUsage:  true, // don't print usage on every error
		SilenceErrors: true, // main prints the error once
	}

	root.AddCommand(
		newAddCmd(svc),
		newListCmd(svc),
		newToggleCmd(svc),
		newUpdateCmd(svc),
		newDeleteCmd(svc),
	)

	return root
}
