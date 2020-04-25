package auth

import (
	"fmt"

	"github.com/carlos0934/videotube/models"
	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	jwt.StandardClaims
	models.User
}

type UserAuth struct {
	key    string
	method jwt.SigningMethod
}

func (auth UserAuth) GenerateToken(claim interface{}) string {
	token, err := auth.method.Sign(auth.key, claim)

	if err != nil {
		fmt.Println(err)
	}

	return token
}

func (auth UserAuth) VerifyToken(payload string) bool {
	token, err := jwt.Parse(payload, auth.ParseToken)
	if err != nil {
		fmt.Println(err)
	}

	err = token.Method.Verify(payload, token.Signature, auth.key)

	return err == nil
}

func (auth UserAuth) ParseToken(token *jwt.Token) (interface{}, error) {

	if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	return []byte(auth.key), nil
}
