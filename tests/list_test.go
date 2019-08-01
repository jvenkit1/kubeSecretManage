package tests

import (
	"github.com/stretchr/testify/assert"
	"kubeSecretManage/services"
	"testing"
)

func TestGetOfListSecrets(t *testing.T) {
	ns := "default"
	err := services.ListSecret(ns, "")
	if !assert.Nil(t, err) {
		t.Errorf("Error retrieving list of secrets : %s", err)
	}
}

func TestGetSecretDetails(t *testing.T) {
	ns := "default"
	secretName := "testsecret"
	err := services.ListSecret(ns, secretName)
	if !assert.Nil(t, err) {
		t.Errorf("Error retrieving list of secrets : %s", err)
	}
}


func TestGetSecretDetailsOfNonExistentSecret(t *testing.T) {
	ns:="default"
	secretName := "invalidName"
	err := services.ListSecret(ns, secretName)
	if !assert.NotNil(t, err) {
		t.Errorf("Expected error")
	}
}