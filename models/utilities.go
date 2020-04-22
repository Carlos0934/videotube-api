package models

import (
	"golang.org/x/crypto/bcrypt"
)

type Hasher struct {
}

func (Hasher) Hash(data string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	CheckError(err)
	return string(hash[:])
}

func (Hasher) Verify(hash string, compare string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(compare))

	return err == nil

}
