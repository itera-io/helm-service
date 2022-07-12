package models

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func GetKubeconfigPath(headers *http.Header) (string, error) {
	log.Println("get kube config")
	if headers == nil || headers.Get("Kubeconfig") == "" {
		return "", errors.New("please specify kubeconfig on the headers")
	}
	log.Println(headers.Get("Kubeconfig"))
	kubeconfig, err := base64.StdEncoding.DecodeString(headers.Get("Kubeconfig"))
	if err != nil {
		return "", errors.New("please specify base64 encoded kubeconfig")
	}
	kubeconfigPath := StoreKubeConfig(kubeconfig)
	return kubeconfigPath, nil
}

func StoreKubeConfig(kubeconfig []byte) string {
	log.Println("store kube config")
	filepath := uuid.New().String() + ".yaml"
	err := ioutil.WriteFile(filepath, kubeconfig, 0)

	if err != nil {
		log.Fatal(err)
	}
	return filepath
}

func RemoveTempFile(kubeconfigPath string) {
	// Removing file from the directory
	// Using Remove() function
	e := os.Remove(kubeconfigPath)
	if e != nil {
		log.Fatal(e)
	}
}
