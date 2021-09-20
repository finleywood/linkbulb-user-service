package services

import (
	"os"
	"time"

	"codedolphin.io/users-service/models"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(user models.User) (string, int64) {
	exp := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     exp,
	})
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return tokenString, exp
}

func ValidateJWT(tokenString string) (bool, uint64) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return false, 0
	}
	claims := token.Claims.(jwt.MapClaims)
	uid := uint64(claims["user_id"].(float64))
	return true, uid
}
