package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(saveCmd)
}

var saveCmd = &cobra.Command{
	Use:   "save",
	Args:  cobra.NoArgs,
	Short: "save url query",
	Long:  "extracts query arguments from url and saves them into a file",
	RunE:  saveQuery,
	//PersistentPostRunE: focusBrowser,
}

func saveQuery(_ *cobra.Command, _ []string) error {
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
