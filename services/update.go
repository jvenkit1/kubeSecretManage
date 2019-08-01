package services

import (
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"

)

func UpdateSecret(ns, name, key, data string) error {
	// ensure secret exists
	err := ListSecret(ns, name)
	if err != nil {
		logrus.WithError(err).Errorf("Secret %s does not exist in the namespace %s", name, ns)
		return err
	}

	// update
	// TODO: Add functionality to append
	keyList := make(map[string][]byte)
	keyList[key] = []byte(data)
	s := &v1.Secret{
		Data: keyList,
	}
	s.SetName(name)
	_, err = clientset.CoreV1().Secrets(ns).Update(s)
	if err != nil {
		logrus.WithError(err).Error("Error creating secret")
		return err
	}
	return nil
}