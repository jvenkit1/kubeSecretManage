package services

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var clientset *kubernetes.Clientset

func init() {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	// if you want to change the loading rules (which files in which order), you can do so here

	configOverrides := &clientcmd.ConfigOverrides{}
	// if you want to change override values or bind them to flags, there are methods to help you

	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	config, err := kubeConfig.ClientConfig()
	if err != nil {
		logrus.WithError(err).Error("Encountered error generating new kubeconfig")
	}
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		logrus.WithError(err).Error("Encountered error generating clientset")
	}
}

func CreateSecret(ns, key, data, name string) error {
	keyList := make(map[string][]byte)
	keyList[key] = []byte(data)
	s := &v1.Secret{
		Data: keyList,
	}
	s.SetName(name)
	_, err := clientset.CoreV1().Secrets(ns).Create(s)
	if err != nil {
		logrus.WithError(err).Error("Error creating secret")
		return err
	}
	return nil
}

func ListSecret(ns, name string) error {
	secret, err := clientset.CoreV1().Secrets(ns).Get(name, metav1.GetOptions{})
	if err != nil {
		logrus.WithError(err).Error("Error retrieving secrets")
		return err
	}
	logrus.WithFields(logrus.Fields{
		"data": secret.Data,
	}).Info("Printing data")
	return nil
}
