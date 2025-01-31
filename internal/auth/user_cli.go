package auth

import (
	"fmt"

	otfapi "github.com/leg100/otf/internal/api"
	"github.com/spf13/cobra"
)

type UserCLI struct {
	UserService
}

func NewUserCommand(api *otfapi.Client) *cobra.Command {
	cli := &UserCLI{}
	cmd := &cobra.Command{
		Use:   "users",
		Short: "User account management",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Parent().PersistentPreRunE(cmd.Parent(), args); err != nil {
				return err
			}
			cli.UserService = &Client{Client: api}
			return nil
		},
	}

	cmd.AddCommand(cli.userNewCommand())
	cmd.AddCommand(cli.userDeleteCommand())

	return cmd
}

func (a *UserCLI) userNewCommand() *cobra.Command {
	return &cobra.Command{
		Use:           "new [username]",
		Short:         "Create a new user account",
		Args:          cobra.ExactArgs(1),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			user, err := a.CreateUser(cmd.Context(), args[0])
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Successfully created user %s\n", user.Username)
			return nil
		},
	}
}

func (a *UserCLI) userDeleteCommand() *cobra.Command {
	return &cobra.Command{
		Use:           "delete [username]",
		Short:         "Delete a user account",
		Args:          cobra.ExactArgs(1),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := a.DeleteUser(cmd.Context(), args[0]); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Successfully deleted user %s\n", args[0])
			return nil
		},
	}
}
