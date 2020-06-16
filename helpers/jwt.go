package helpers

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	// log "github.com/sirupsen/logrus"
)

// CreateJWTToken ...
func CreateJWTToken(id uint, email string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authenticated"] = true
	atClaims["user_id"] = id
	atClaims["user_email"] = email
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	return at.SignedString([]byte("HelloWorld123"))
}