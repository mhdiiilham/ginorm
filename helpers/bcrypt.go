package helpers

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword ...
func HashPassword(p []byte) string {
	hash, err := bcrypt.GenerateFromPassword(p, bcrypt.MinCost)

	if err != nil {
		log.Fatal("There's problem in hasing password. Error: ", err)
	}

	return string(hash)
}


// ComparePassword ...
func ComparePassword(hashed string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), plain)

	if err != nil {
		return false
	}

	return true
}