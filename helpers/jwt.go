package helpers

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	m "github.com/mhdiiilham/ginorm/models"
	log "github.com/sirupsen/logrus"
)

// CreateJWTToken ...
func CreateJWTToken(id uint, email string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authenticated"] = true
	atClaims["user_id"] = id
	atClaims["user_email"] = email
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	return at.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// VerifyToken ...
func VerifyToken(ht string) (*jwt.Token, error) {
	token, err := jwt.Parse(ht, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", ok)
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

// TokenValid ...
func TokenValid(ht string) error {
	token, err := VerifyToken(ht)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// ExtractedJWT ...
func ExtractedJWT(ht string) (m.TokenMetaData, error) {
	tokenMetaData := m.TokenMetaData{}
	token, err := VerifyToken(ht)
	if err != nil {
		log.Info(err)
		return tokenMetaData, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userEmail, ok := claims["user_email"].(string)
		if !ok {
			return tokenMetaData, err
		}
		userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return tokenMetaData, err
		}
		tokenMetaData.ID = userID
		tokenMetaData.Email = userEmail
		return tokenMetaData, nil
	}
	return tokenMetaData, err
}
