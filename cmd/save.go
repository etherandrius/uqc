package cmd

import (
	"io/ioutil"
	"net/url"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

const (
	linuxSecurityMode = 0777
)

func init() {
	rootCmd.AddCommand(saveCmd)
}

var saveCmd = &cobra.Command{
	Use:   "save",
	Args:  cobra.ExactArgs(1),
	Short: "save url query",
	Long:  "extracts query arguments from url and saves them into a file",
	RunE:  saveQuery,
	//PersistentPostRunE: focusBrowser,
}

func saveQuery(_ *cobra.Command, args []string) error {
	url, err := url.Parse(args[0])
	if err != nil {
		return err
	}

	dir, err := homedir.Dir()
	if err != nil {
		return err
	}
	filedir := dir + "/.cache/uqc/"

	_, err = os.Stat(filedir)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(filedir, linuxSecurityMode)
		} else {
			return err
		}
	}

	err = ioutil.WriteFile(filedir+url.Host, []byte(args[0]), linuxSecurityMode)
	if err != nil {
		return err
	}
	return nil
}
