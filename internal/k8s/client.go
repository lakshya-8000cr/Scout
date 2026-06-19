package k8s

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"context"
	"io"

	corev1 "k8s.io/api/core/v1"
)

func NewClient() (*kubernetes.Clientset, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	kubeconfig := filepath.Join(home, ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}

func GetPodLogs(
	clientset *kubernetes.Clientset,
	namespace string,
	podName string,
	lines int64,
) (string, error) {

	req := clientset.CoreV1().
		Pods(namespace).
		GetLogs(
			podName,
			&corev1.PodLogOptions{
				TailLines: &lines,
			},
		)

	stream, err := req.Stream(context.Background())

	if err != nil {
		return "", err
	}

	defer stream.Close()

	data, err := io.ReadAll(stream)

	if err != nil {
		return "", err
	}

	return string(data), nil
}