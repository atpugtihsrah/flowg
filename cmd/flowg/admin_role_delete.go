package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"link-society.com/flowg/internal/data/auth"
)

type adminRoleDeleteOpts struct {
	authDir string
	name    string
}

func NewAdminRoleDeleteCommand() *cobra.Command {
	opts := &adminRoleDeleteOpts{}

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an existing role",
		Run: func(cmd *cobra.Command, args []string) {
			authDb := auth.NewDatabase(
				auth.DefaultDatabaseOpts().WithDir(opts.authDir),
			)
			err := authDb.Open()
			if err != nil {
				fmt.Fprintln(os.Stderr, "ERROR: Failed to open auth database:", err)
				exitCode = 1
				return
			}
			defer func() {
				err := authDb.Close()
				if err != nil {
					fmt.Fprintln(os.Stderr, "ERROR: Failed to close auth database:", err)
					exitCode = 1
				}
			}()

			roleSys := auth.NewRoleSystem(authDb)
			err = roleSys.DeleteRole(opts.name)
			if err != nil {
				fmt.Fprintln(os.Stderr, "ERROR: Failed to delete role:", err)
				exitCode = 1
				return
			}

			fmt.Println("Role deleted")
		},
	}

	cmd.Flags().StringVar(
		&opts.authDir,
		"auth-dir",
		defaultAuthDir,
		"Path to the log database directory",
	)
	cmd.MarkFlagDirname("auth-dir")

	cmd.Flags().StringVar(
		&opts.name,
		"name",
		"",
		"Name of the role",
	)
	cmd.MarkFlagRequired("name")

	return cmd
}
