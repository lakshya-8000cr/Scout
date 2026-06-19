package cmd

import (
	"fmt"
	"context"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"scout/internal/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var podsCmd = &cobra.Command{
	Use:   "pods",
	Short: "Show pods in a pretty format",
	Run: func(cmd *cobra.Command, args []string) {

		clientset, err := k8s.NewClient() // this is the part to established connection with k8s api

		if err != nil {
			fmt.Println(err)
			return
		}

		pods, err := clientset.CoreV1(). // v1 refers to all the resources like deployments , services , pods etc 
			Pods(namespace).  // show pods in default namespce
			List(
				context.Background(),
				metav1.ListOptions{}, // list all
			)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println()

		printLabel("SCOUT", "Kubernetes pod status")

		fmt.Println()

		for _, pod := range pods.Items {  // for looop iteration over the pods we are fetching

			printPod(
				"READY",
				pod.Name,
				string(pod.Status.Phase),
				0,
			)
		}

		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(podsCmd) 
}

func printPod(label, name, status string, restarts int) {
	switch label {
	case "READY":
		color.New(color.FgGreen, color.Bold).Printf("  ✓ %-8s", label)
	case "WARN":
		color.New(color.FgYellow, color.Bold).Printf("  ⚠ %-8s", label)
	case "ERROR":
		color.New(color.FgRed, color.Bold).Printf("  ✗ %-8s", label)
	default:
		color.New(color.FgWhite, color.Bold).Printf("  • %-8s", label)
	}

	color.New(color.FgWhite).Printf(" %-38s", name)

	switch status {
	case "Running":
		color.New(color.FgGreen).Printf("%-14s", status)
	case "Pending":
		color.New(color.FgYellow).Printf("%-14s", status)
	default:
		color.New(color.FgRed).Printf("%-14s", status)
	}

	restartStr := fmt.Sprintf("restarts(%d)", restarts)
	if restarts > 0 {
		color.New(color.FgYellow).Printf(" %s\n", restartStr)
	} else {
		color.New(color.Faint).Printf(" %s\n", restartStr)
	}
}

func printLabel(label, msg string) {
	switch label {
	case "INFO":
		color.New(color.FgGreen, color.Bold).Printf("  ✓ %-8s", label)
		color.New(color.FgWhite).Println(" " + msg)
	case "WARN":
		color.New(color.FgYellow, color.Bold).Printf("  ⚠ %-8s", label)
		color.New(color.FgWhite).Println(" " + msg)
	case "ERROR":
		color.New(color.FgRed, color.Bold).Printf("  ✗ %-8s", label)
		color.New(color.FgWhite).Println(" " + msg)
	case "LOG":
		color.New(color.FgCyan, color.Bold).Printf("  • %-8s", label)
		color.New(color.FgWhite).Println(" " + msg)
	case "SCOUT":
		color.New(color.BgHiBlue, color.FgWhite, color.Bold).Printf(" %s ", label)
		color.New(color.Bold, color.FgWhite).Println("  " + msg)
		color.New(color.Faint).Println("  └───" + "─────────────────────────────────────────────────────────────")
	default:
		color.New(color.BgWhite, color.FgBlack, color.Bold).Printf(" %-7s ", label)
		fmt.Println(" " + msg)
	}
}