package tests

import (
	"github.com/stretchr/testify/assert"
	"kubeSecretManage/services"
	"testing"
)



func TestCreateWithValidCredentials(t *testing.T) {
	ns := "default"
	key := "secretkey"
	data := "secretdata"
	secretName := "testsecret"
	err := services.CreateSecret(ns, key, data, secretName)
	if !assert.Nil(t, err) {
		t.Errorf("Error creating secrets : %s", err)
	}
}

func TestCreateSecretWithInvalidSecretName(t *testing.T ){
	ns := "default"
	key := "secretkey"
	data := "secretdata"
	secretName := "testSecret" // creating secrets with name in camel case is not permitted
	err := services.CreateSecret(ns, key, data, secretName)
	if !assert.NotNil(t, err) {
		t.Errorf("Error creating secrets : %s", err)
	}
}