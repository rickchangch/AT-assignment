package main

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const TTL = 5 * time.Second

var (
	expiredToken string = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsic29tZW9uZSJdLCJleHAiOjE2NzY5OTkxNTYsImZvbyI6ImJhciIsImlzcyI6InRlc3QiLCJuYmYiOjE2NzY5OTkxNTEsInN1YmplY3QiOiJ0ZXN0In0.nDtJpi8SFxWwMXl_AYOxpbuvUwh9ykkU5DuvVfc8T75JF62uykkQ4ndpejtnM0ixvRvbKZId3U-BtgDEwi_A2ha1TDk-V14teURwJCBWZXfMndyHJWB6EGKFz9ZVMxGeB9kmDabh7XfqsqpO2dgT9XlXe4WUJ1QGosn3qHjRxSoMQfl1xAKcpeIWl5ryQOTbIGvqxyKYlf2mwnc1HMJ3ywcHiKOR763pf3KltG74pDRmAHWbRxddYJOT-TbnhL6emMiuFevi-EkC7I0ZEI35vAIAFbqDMXhyZ5Mg6IheS5J6M6VVMH0O12C78MGdrpjS9af-laeoxR9BXhXnGELADA"
	privateKey   *rsa.PrivateKey
	publicKey    *rsa.PublicKey
)

func init() {
	// Parse private key
	privKeyByte, err := ioutil.ReadFile("rsa-private-key.pem")
	if err != nil {
		log.Println(err.Error())
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privKeyByte)
	if err != nil {
		log.Println(err.Error())
	}

	// Parse public key
	pubByte, err := ioutil.ReadFile("rsa-public-key.pem")
	if err != nil {
		log.Println(err.Error())
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(pubByte)
	if err != nil {
		log.Println(err.Error())
	}
}

// Generate the development key pair for encrypting/validating JWT.
func main() {

	token, err := generateJWT()
	if err != nil {
		fmt.Println("Error occurred when generated JWT.", err.Error())
	}

	res, _ := validateJWT(token)
	log.Printf("The result of token validation: [%v]", res)

	res, _ = validateJWT(expiredToken)
	log.Printf("The result of expired token validation: [%v]", res)
}

// Generate JWT token
func generateJWT() (string, error) {

	// Create a new token object
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"aud":     []string{"someone"},
		"iss":     "test",
		"subject": "test",
		"exp":     jwt.NewNumericDate(time.Now().Add(TTL)),
		"nbf":     jwt.NewNumericDate(time.Now()),
		"foo":     "bar",
	})

	// Sign and get the complete encoded token as a string using the private key
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Validate JWT
func validateJWT(tokenString string) (bool, error) {

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
