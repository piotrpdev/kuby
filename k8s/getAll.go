package k8s

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// var c = GetClientset() // This causes errors when testing

func GetAllPods() (*v1.PodList, error) {
	return GetClientset().CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
}

func GetAllServices() (*v1.ServiceList, error) {
	return GetClientset().CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
}

func GetAllEndpoints() (*v1.EndpointsList, error) {
	return GetClientset().CoreV1().Endpoints("").List(context.TODO(), metav1.ListOptions{})
}
