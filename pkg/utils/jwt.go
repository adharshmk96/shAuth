package utils

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"os"
	"time"
)

var (
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
)

func loadPrivateKey() (*ecdsa.PrivateKey, error) {
	if privateKey != nil {
		return privateKey, nil
	}

	privateKeyPath := viper.GetString("auth.jwt.private_key_path")
	content, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(content)
	if block == nil {
		return nil, err
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func loadPublicKey() (*ecdsa.PublicKey, error) {
	if publicKey != nil {
		return publicKey, nil
	}

	publicKeyPath := viper.GetString("auth.jwt.public_key_path")
	content, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(content)
	if block == nil {
		return nil, err
	}

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey, ok := publicKeyInterface.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}

	return publicKey, nil
}

func SetKeys(priv *ecdsa.PrivateKey, pub *ecdsa.PublicKey) {
	privateKey = priv
	publicKey = pub
}

func EncodeJWT(claims jwt.Claims) (string, error) {
	privateKey, err := loadPrivateKey()
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES512, claims)
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DecodeJWT(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return loadPublicKey()
	})
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func NewClaims() *jwt.MapClaims {

	expSeconds := viper.GetInt("auth.jwt.expirySeconds")
	now := time.Now()
	expiry := time.Now().Add(time.Second * time.Duration(expSeconds)).Unix()

	return &jwt.MapClaims{
		"sub": "authentication",
		"iss": "ServiceHub",
		"aud": []string{
			"ServiceHub",
		},
		"iat": now,
		"exp": expiry,
	}
}
