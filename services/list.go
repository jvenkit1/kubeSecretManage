package services

import (
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListSecret(ns, name string) error {

	if name != "" {
		secret, err := clientset.CoreV1().Secrets(ns).Get(name, metav1.GetOptions{})
		if err != nil {
			logrus.WithError(err).Error("Error retrieving secrets")
			return err
		}

		for key, value := range secret.Data {
			logrus.WithFields(
				logrus.Fields{
					"Key": key,
					"Value": string(value),
				}).Info("Printing kvp")
		}
	}else {

		secretList, err := clientset.CoreV1().Secrets(ns).List(metav1.ListOptions{})
		if err != nil {
			logrus.WithError(err).Error("Error retrieving secret list")
			return err
		}

		for _, element := range secretList.Items {
			logrus.WithFields(logrus.Fields{
				"SecretName": element.Name,
			}).Info("Printing List of secrets")
		}

	}
	return nil
}
