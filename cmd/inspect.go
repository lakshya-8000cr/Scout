package cmd
                                                      
import (
	"context"
	"fmt"

	"scout/internal/k8s"

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
			Pods("default").
			Get(context.Background(), podName, metav1.GetOptions{})

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println()
		printLabel("SCOUT", fmt.Sprintf("Pod diagnosis : %s", podName))
		fmt.Println()

		printLabel("INFO", fmt.Sprintf("STATUS   %s", pod.Status.Phase))
		printLabel("INFO", fmt.Sprintf("NODE     %s", pod.Spec.NodeName))
		isHealthy := true

		for _, cs := range pod.Status.ContainerStatuses {
			printLabel("INFO", fmt.Sprintf("CONTAINER %s", cs.Name))
			printLabel("INFO", fmt.Sprintf("RESTARTS  %d", cs.RestartCount))

			if cs.State.Waiting != nil {
				isHealthy = false
				reason := cs.State.Waiting.Reason
				message := cs.State.Waiting.Message

				printLabel("WARN", fmt.Sprintf("WAITING  %s", reason))

				if message != "" {
					printLabel("ERROR", message)
				}

				printHint(reason)
			}

			if cs.State.Terminated != nil {
				isHealthy = false
				printLabel("ERROR", fmt.Sprintf("TERMINATED %s", cs.State.Terminated.Reason))
				if cs.State.Terminated.Message != "" {
					printLabel("ERROR", cs.State.Terminated.Message)
				}
			}
		}

		if isHealthy {

	printLabel(
		"READY",
		"POD HEALTHY - No issues detected",
	)
}

		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}

func printHint(reason string) {
	switch reason {
	case "ImagePullBackOff", "ErrImagePull":
		printLabel("INFO", "HINT     Check image name, tag, registry access, or imagePullSecrets")
	case "CrashLoopBackOff":
		printLabel("INFO", "HINT     Check application logs using: scout logs <pod-name>")
	case "CreateContainerConfigError":
		printLabel("INFO", "HINT     Check ConfigMap, Secret, env vars, or volume mounts")
	case "RunContainerError":
		printLabel("INFO", "HINT     Container failed to start. Check command, args, mounts, and permissions")
	default:
		printLabel("INFO", "HINT     Check pod events using: scout events")
	}
}