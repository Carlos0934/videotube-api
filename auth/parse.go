package auth

import (
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func GetECPrivateKey(privateKeyDir, publicKeyDir string) *ecdsa.PrivateKey {

	privateKeyData, err := ioutil.ReadFile(privateKeyDir)
	checkErr(err)
	publicKeyData, err := ioutil.ReadFile(publicKeyDir)
	checkErr(err)

	privateKey, err := jwt.ParseECPrivateKeyFromPEM(privateKeyData)
	checkErr(err)
	publicKey, err := jwt.ParseECPublicKeyFromPEM(publicKeyData)
	checkErr(err)

	privateKey.PublicKey = *publicKey

	return privateKey
}
