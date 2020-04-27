package auth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"database/sql"
	"fmt"

	"github.com/carlos0934/videotube/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserClaims struct {
	jwt.StandardClaims
	models.User
}

type UserAuth struct {
	key     *ecdsa.PrivateKey
	method  *jwt.SigningMethodECDSA
	Storage *models.UserStorage
}

func NewUserAuth(conn *sql.DB) *UserAuth {
	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	return &UserAuth{
		key:     key,
		method:  jwt.SigningMethodES256,
		Storage: models.NewUserStorage(conn),
	}
}

func (auth *UserAuth) VerifyPassword(user models.User) bool {
	filter := map[string]interface{}{"id": user.ID, "username": user.Username, "email": user.Email}
	dbUser := models.User{}
	err := auth.Storage.FindOne(filter, dbUser)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	return err == nil
}
func (auth *UserAuth) GenerateToken(user models.User) string {
	claims := &UserClaims{
		User: user,
	}
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

func (auth *UserAuth) VefifyUser(payload string, claims *UserClaims) bool {
	if auth.VerifyToken(payload, claims) {
		query := map[string]interface{}{"id": claims.ID, "password": claims.Password, "username": claims.Username}
		user := &models.User{}
		err := auth.Storage.FindOne(query, user)

		if err != nil {
			return false
		}

		if user.Username != "" {
			return true
		}
	}

	return false
}
