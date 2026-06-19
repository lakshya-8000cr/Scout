package cmd

import (
	"context"
	"fmt"
"scout/internal/k8s"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect pod <pod-name>",
	Short: "Show simplified pod diagnosis",
	Args:  cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		resourceType := args[0]
		podName := args[1]

		if resourceType != "pod" {
			fmt.Println("Only pod inspection is supported for now")
			return
		}

		clientset, err := k8s.NewClient()
		if err != nil {
			fmt.Println(err)
			return
		}

		pod, err := clientset.CoreV1().
			Pods(namespace).
			Get(context.Background(), podName, metav1.GetOptions{})

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println()
		printLabel("SCOUT", fmt.Sprintf("Pod diagnosis : %s", podName))
		fmt.Println()

		printInspectLine("INFO", "STATUS", string(pod.Status.Phase))
		printInspectLine("INFO", "NODE", pod.Spec.NodeName)
		isHealthy := true

		// ── Step 1: Handle Scheduler Level / Pending Errors ──
		if pod.Status.Phase == "Pending" {
			for _, cond := range pod.Status.Conditions {
				if cond.Type == "PodScheduled" && cond.Status == "False" {
					isHealthy = false
					printInspectLine("ERROR", "SCHEDULER", cond.Reason)
					if cond.Message != "" {
						color.New(color.Faint).Print("  │    ├── ")
						color.New(color.FgRed).Println(cond.Message)
					}
					
					// Print specific hint for scheduling issues
					color.New(color.Faint).Print("  │    └── ")
					color.New(color.FgCyan, color.Bold).Print("HINT: ")
					color.New(color.FgHiWhite).Println("Cluster resource crunch. Scale your node capacity or delete old pods.")
				}
			}
		}

		// ── Step 2: Handle Container Level Errors (Only if statuses exist) ──
		if len(pod.Status.ContainerStatuses) > 0 {
			for _, cs := range pod.Status.ContainerStatuses {
				printInspectLine("INFO", "CONTAINER", cs.Name)
				printInspectLine("INFO", "RESTARTS", fmt.Sprintf("%d", cs.RestartCount))

				if cs.State.Waiting != nil {
					isHealthy = false
					reason := cs.State.Waiting.Reason
					message := cs.State.Waiting.Message

					printInspectLine("WARN", "WAITING", reason)

					if message != "" {
						color.New(color.Faint).Print("  │    ├── ")
						color.New(color.FgRed).Println(message)
					}

					printHint(reason)
				}

				if cs.State.Terminated != nil {
					isHealthy = false
					printInspectLine("ERROR", "TERMINATED", cs.State.Terminated.Reason)
					if cs.State.Terminated.Message != "" {
						color.New(color.Faint).Print("  │    ├── ")
						color.New(color.FgRed).Println(cs.State.Terminated.Message)
					}
				}
			}
		} else if pod.Status.Phase != "Pending" {
			// If no containers exist and it's not pending, something is structurally off
			isHealthy = false
			printInspectLine("ERROR", "CONTAINERS", "No containers initialized or metadata missing")
		}

		if isHealthy {
			fmt.Println("  │")
			color.New(color.FgGreen, color.Bold).Println("  ✓ READY    POD HEALTHY - No issues detected")
		}

		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}

func printHint(reason string) {
	color.New(color.Faint).Print("  │    └── ")
	color.New(color.FgCyan, color.Bold).Print("HINT: ")

	switch reason {
	case "ImagePullBackOff", "ErrImagePull":
		color.New(color.FgHiWhite).Println("Check image name, tag, registry access, or imagePullSecrets")
	case "CrashLoopBackOff":
		color.New(color.FgHiWhite).Println("Check application logs using: scout logs <pod-name>")
	case "CreateContainerConfigError":
		color.New(color.FgHiWhite).Println("Check ConfigMap, Secret, env vars, or volume mounts")
	case "RunContainerError":
		color.New(color.FgHiWhite).Println("Check container command, args, mounts, and permissions")
	default:
		color.New(color.FgHiWhite).Println("Check pod events using: scout events")
	}
}

func printInspectLine(label, key, value string) {
	color.New(color.Faint).Print("  │  ")

	switch label {
	case "INFO":
		color.New(color.FgGreen).Printf("├── %-12s", key)
	case "WARN":
		color.New(color.FgYellow, color.Bold).Printf("├── %-12s", key)
	case "ERROR":
		color.New(color.FgRed, color.Bold).Printf("├── %-12s", key)
	}

	color.New(color.FgWhite).Println(value)
}