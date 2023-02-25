package service

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Auth struct{}

var AuthService Auth

const TTL = 2 * time.Hour

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func init() {
	// Parse private key
	privKeyByte, err := ioutil.ReadFile("base/secret/rsa-private-key.pem")
	if err != nil {
		log.Println(err.Error())
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privKeyByte)
	if err != nil {
		log.Println(err.Error())
	}

	// Parse public key
	pubByte, err := ioutil.ReadFile("base/secret/rsa-public-key.pem")
	if err != nil {
		log.Println(err.Error())
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(pubByte)
	if err != nil {
		log.Println(err.Error())
	}
}

// Generate JWT token
func (a *Auth) GenerateJWT(acct string) (string, error) {

	// Create a new token object
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"aud":     []string{"someone"},
		"iss":     "test",
		"subject": "test",
		"exp":     jwt.NewNumericDate(time.Now().Add(TTL)),
		"nbf":     jwt.NewNumericDate(time.Now()),
		"acct":    acct,
	})

	// Sign and get the complete encoded token as a string using the private key
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Validate JWT
func (a *Auth) ValidateJWT(tokenString string) (bool, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// Validate alg
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return publicKey, nil
	})
	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
