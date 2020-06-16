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

func verifyToken(ht string) (*jwt.Token, error) {
	token, err := jwt.Parse(ht, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", ok)
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil{
		return nil, err
	}

	return token, nil
}
