package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	namespace string = "default"
	sleepTime int    = 5
)

func init() {
	if len(os.Getenv("SLEEP_TIME")) != 0 {
		sleepTime, _ = strconv.Atoi(os.Getenv("SLEEP_TIME"))
	}
	if len(os.Getenv("NAMESPACE")) != 0 {
		namespace = os.Getenv("NAMESPACE")
	}
}

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	for {
		pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Existen %d pods en el namespace %s\n", len(pods.Items), namespace)

		time.Sleep(time.Duration(sleepTime) * time.Second)
	}
}
