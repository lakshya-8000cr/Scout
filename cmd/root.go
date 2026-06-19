package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scout",
	Short: "Scout your Kubernetes cluster",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil { //this will run the command entered by th user
		fmt.Println(err)

		os.Exit(1)
	}
}