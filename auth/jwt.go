package auth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"

	"github.com/carlos0934/videotube/models"
	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	jwt.StandardClaims
	models.User
}

type UserAuth struct {
	key    *ecdsa.PrivateKey
	method *jwt.SigningMethodECDSA
}

func NewUserAuth() *UserAuth {
	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	return &UserAuth{
		key:    key,
		method: jwt.SigningMethodES256,
	}
}

func (auth *UserAuth) GenerateToken(claims *UserClaims) string {

	token := jwt.NewWithClaims(auth.method, claims)
	stringToken, err := token.SignedString(auth.key)

	if err != nil {

		fmt.Println(err)
	}
	return stringToken
}

func (auth UserAuth) VerifyToken(payload string, claims *UserClaims) bool {
	_, err := jwt.ParseWithClaims(payload, claims, auth.ParseToken)

	return err == nil
}

func (auth *UserAuth) ParseToken(token *jwt.Token) (interface{}, error) {

	return &auth.key.PublicKey, nil
}
