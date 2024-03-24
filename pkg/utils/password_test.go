package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateSecureSalt(t *testing.T) {
	result, err := GenerateSecureSalt(32)
	assert.NotEqual(t, "", result, "Expected salt to not be empty")
	assert.True(t, len(result) > 10, "Expected salt to be longer than 10 characters")
	assert.Nil(t, err, "Expected no error")
}

func TestHashPassword(t *testing.T) {
	result, err := HashPassword("password")
	assert.NotEqual(t, "", result, "Expected hash to not be empty")
	assert.NoError(t, err, "Expected no error")
}

func TestVerifyPassword(t *testing.T) {
	hash, _ := HashPassword("password")
	result, err := VerifyPassword(hash, "password")
	assert.True(t, result, "Expected password to be verified")
	assert.NoError(t, err, "Expected no error")
}
