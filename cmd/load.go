package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(gitRemoteCmd)
}

var gitRemoteCmd = &cobra.Command{
	Use:   "load",
	Args:  cobra.NoArgs,
	Short: "load perviously saved query arguments",
	Long:  "finds and loads previously saved query arguments",
	RunE:  loadQuery,
	// PersistentPostRunE: focusBrowser,
}

func loadQuery(_ *cobra.Command, _ []string) error {
	// s, err := git.GetRoot()
	// if err != nil {
	// 	return err
	// }
	// ok, err := tabNavigaotr.Focus(s)
	// if err != nil {
	// 	return err
	// }
	// if !ok {
	// 	tabNavigaotr.Open(s)
	// }
	return nil
}
