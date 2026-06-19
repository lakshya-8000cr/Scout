package cmd

import (                 // top architecture is lil diff unlike of pods status which are directly available metrics are not available directly
	"context"            // it have to use metric api which we create thru the config( you can assume this is address of where metric is stored ) 
	"fmt"                // the  it calls MetricsV1beta1() go and take some cpu/ram , but remember Newconfig doesnot fetch its just adress

	"scout/internal/k8s"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metricsclient "k8s.io/metrics/pkg/client/clientset/versioned"
)

var topCmd = &cobra.Command{
	Use:   "top",
	Short: "Show pod CPU and memory usage",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := k8s.NewConfig()
		if err != nil {
			fmt.Println(err)
			return
		}

		metricsClient, err := metricsclient.NewForConfig(config)
		if err != nil {
			fmt.Println(err)
			return
		}

		podMetrics, err := metricsClient.MetricsV1beta1().
			PodMetricses("default").
			List(context.Background(), metav1.ListOptions{})

		if err != nil {
			fmt.Println("metrics not available. Enable metrics-server first.")
			fmt.Println(err)
			return
		}

		fmt.Println()
		printLabel("SCOUT", "Resource Utilization (CPU/RAM)")
		fmt.Println()

		for _, pod := range podMetrics.Items {
			var cpuMilli int64
			var memBytes int64

			for _, container := range pod.Containers {
				cpuMilli += container.Usage.Cpu().MilliValue()
				memBytes += container.Usage.Memory().Value()
			}

			memMB := memBytes / (1024 * 1024)

			fmt.Printf("  │  ├── %-45s CPU: %-8dm MEM: %dMB\n", pod.Name, cpuMilli, memMB)
		}

		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(topCmd)
}