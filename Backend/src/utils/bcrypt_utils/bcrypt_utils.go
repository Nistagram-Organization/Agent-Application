package bcrypt_utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func GetHash(s string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to hash value %s", s))
	}
	return string(hash)
}

func CompareHashAndValue(hash string, s string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(s)); err != nil {
		return false
	}
	return true
}
