package healthcheck

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func GetJobStatus(namespace string, jobName string) {

	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Error getting in-cluster config: %s", err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes clientset: %s", err.Error())
	}

	job, err := clientset.BatchV1().Jobs(namespace).Get(context.TODO(), jobName, metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Error getting job: %s", err.Error())
	}

	fmt.Printf("Job Name: %s\n", job.Name)
	fmt.Printf("Namespace: %s\n", job.Namespace)
	fmt.Printf("Labels: %v\n", job.Labels)
	fmt.Printf("Completions: %d\n", *job.Spec.Completions)
	fmt.Printf("Active: %d\n", job.Status.Active)
	fmt.Printf("Succeeded: %d\n", job.Status.Succeeded)

}
