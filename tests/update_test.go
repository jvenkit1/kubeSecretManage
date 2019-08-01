package tests

import (
	"github.com/stretchr/testify/assert"
	"kubeSecretManage/services"
	"testing"
)

func TestUpdateWithValidName(t *testing.T) {
	ns := "default"
	secretName := "testsecret"
	key := "secretKey"
	value := "secretValue"
	err := services.UpdateSecret(ns, secretName, key, value)
	if !assert.Nil(t, err) {
		t.Errorf("Error performing update operation : %s", err)
	}
}

func TestUpdateWithInvalidName(t *testing.T) {
	ns := "default"
	secretName := "invalidName"
	key := "secretKey"
	value := "secretValue"
	err := services.UpdateSecret(ns, secretName, key, value)
	if !assert.NotNil(t, err) {
		t.Errorf("Error performing update operation : %s", err)
	}
}
