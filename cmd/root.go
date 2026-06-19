package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var namespace string

var rootCmd = &cobra.Command{
	Use:   "scout",
	Short: "Scout your Kubernetes cluster",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&namespace,
		"namespace",
		"n",
		"default",
		"Kubernetes namespace",
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil { //this will run the command entered by th user
		fmt.Println(err)

		os.Exit(1)
	}
}