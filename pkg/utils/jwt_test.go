package utils_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"
	"time"

	"github.com/adharshmk96/shAuth/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func setup() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		panic(err)
	}

	publicKey := &privateKey.PublicKey

	utils.SetKeys(privateKey, publicKey)
}

func TestEncodeJWT(t *testing.T) {
	setup()

	claims := &jwt.MapClaims{
		"email": "test@email.com",
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token, err := utils.EncodeJWT(claims)
	assert.NoError(t, err, "Error should be nil")
	assert.NotEmpty(t, token, "Token should not be empty")
}

func TestDecodeJWT(t *testing.T) {
	setup()

	claims := &jwt.MapClaims{
		"email": "test@email.com",
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token, err := utils.EncodeJWT(claims)
	assert.NoError(t, err, "Error should be nil")
	assert.NotEmpty(t, token, "Token should not be empty")

	decodedClaims, err := utils.DecodeJWT(token)
	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, "test@email.com", decodedClaims["email"], "Email should be equal")
}
