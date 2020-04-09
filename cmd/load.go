package cmd

import (
	"fmt"
	"io/ioutil"
	"net/url"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(gitRemoteCmd)
}

var gitRemoteCmd = &cobra.Command{
	Use:   "load",
	Args:  cobra.ExactArgs(1),
	Short: "load perviously saved query arguments",
	Long:  "finds and loads previously saved query arguments",
	RunE:  loadQuery,
	// PersistentPostRunE: focusBrowser,
}

func loadQuery(_ *cobra.Command, args []string) error {
	inputURL, err := url.Parse(args[0])
	if err != nil {
		return err
	}

	dir, err := homedir.Dir()
	if err != nil {
		return err
	}
	filedir := dir + "/.cache/uqc/"

	fileBytes, err := ioutil.ReadFile(filedir + inputURL.Host)
	if err != nil {
		return err
	}

	savedURL, err := inputURL.Parse(string(fileBytes))
	if err != nil {
		return err
	}

	queryValues := inputURL.Query()
	for key, valArr := range savedURL.Query() {
		queryValues.Del(key)
		for _, val := range valArr {
			queryValues.Add(key, val)
		}
	}

	outputURL := url.URL{
		Scheme:     inputURL.Scheme,
		Opaque:     inputURL.Opaque,
		User:       inputURL.User,
		Host:       inputURL.Host,
		Path:       inputURL.Path,
		RawPath:    inputURL.RawPath,
		ForceQuery: inputURL.ForceQuery,
		RawQuery:	queryValues.Encode(),
		Fragment:   inputURL.Fragment,
	}

	fmt.Println(outputURL.String())
	return nil
}
