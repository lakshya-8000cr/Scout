package cmd

import (
	"fmt"
	"strings"

	"scout/internal/k8s"

	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs <pod-name>",  // command name 
	Short: "Show logs for a pod",

	Args: cobra.ExactArgs(1), // passing the api

Run: func(cmd *cobra.Command, args []string) {

	podName := args[0]   // getting the api

	clientset, err := k8s.NewClient() // hitting the k8s api

	if err != nil {
		fmt.Println(err)
		return
	}

	lines := int64(50) 

	logs, err := k8s.GetPodLogs(
		clientset,
		"default", // namespace
		podName,
		lines,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()

	printLabel(
		"SCOUT",
		fmt.Sprintf("Logs : %s", podName), //
	)

	fmt.Println()

for _, line := range strings.Split(logs, "\n") {  // logs printing thru loop line by line 

	if line == "" {
		continue
	}

	printLogLine(line)
}

	fmt.Println()
},
}

func init() {

	rootCmd.AddCommand(logsCmd)  //calling

}

func detectLogLevel(line string) string {

	line = strings.ToLower(line)

	switch {

	case strings.Contains(line, "error"),
		strings.Contains(line, "failed"),
		strings.Contains(line, "fatal"),
		strings.Contains(line, "panic"):

		return "ERROR"

	case strings.Contains(line, "warn"),
		strings.Contains(line, "warning"):

		return "WARN"

	case strings.Contains(line, "connected"),
		strings.Contains(line, "ready"),
		strings.Contains(line, "running"),
		strings.Contains(line, "deployed"),
		strings.Contains(line, "installing"):

		return "INFO"

	default:

		return "LOG"
	}
}

func printLogLine(line string) {

	level := detectLogLevel(line)

	printLabel(
		level,
		line,
	)
}