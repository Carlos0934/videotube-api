package auth

import (
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/dgrijalva/jwt-go"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func logErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
func GetECPrivateKey(privateKeyDir, publicKeyDir string) *ecdsa.PrivateKey {

	privateKeyData, err := ioutil.ReadFile(privateKeyDir)
	logErr(err)
	publicKeyData, err := ioutil.ReadFile(publicKeyDir)
	logErr(err)

	privateKey, err := jwt.ParseECPrivateKeyFromPEM(privateKeyData)
	logErr(err)
	publicKey, err := jwt.ParseECPublicKeyFromPEM(publicKeyData)
	logErr(err)

	privateKey.PublicKey = *publicKey

	return privateKey
}
