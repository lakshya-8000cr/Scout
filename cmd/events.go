package cmd

import (
	"context"
	"fmt"

	"scout/internal/k8s"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var eventsCmd = &cobra.Command{ // same thing is repeating in every function 1> configure the kube/config file in client.go 
	Use:   "events",                                                     //  2> getting the k8s api hitting byt auth , rbac
	Short: "Show Kubernetes events in a pretty format",                  //  3> corev1 = it containe all the resources like pods deploy etc
	Run: func(cmd *cobra.Command, args []string) {                       //  4> looping and printing the events by .header ex .reason
		clientset, err := k8s.NewClient()
		if err != nil {
			fmt.Println(err)
			return
		}

		events, err := clientset.CoreV1().
			Events("default").
			List(context.Background(), metav1.ListOptions{})

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println()
		printLabel("SCOUT", "Cluster events")
		fmt.Println()

		for _, event := range events.Items {
			label := "INFO"
			if event.Type == "Warning" {
				label = "WARN"
			}

			printEventLine(label, event.Reason, event.InvolvedObject.Kind+"/"+event.InvolvedObject.Name, event.Message)
		}

		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(eventsCmd)
}

func printEventLine(label, reason, object, message string) {
	switch label {
	case "INFO":
		color.New(color.FgGreen, color.Bold).Printf("  ✓ %-8s", label)
	case "WARN":
		color.New(color.FgYellow, color.Bold).Printf("  ⚠ %-8s", label)
	default:
		color.New(color.FgWhite, color.Bold).Printf("  • %-8s", label)
	}

	color.New(color.FgHiCyan).Printf(" %-20s", reason)
	color.New(color.FgWhite).Printf(" %-30s", object)
	color.New(color.Faint).Printf(" %s\n", message)
}