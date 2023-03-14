package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

const SigningSecretKey = "secret"

func main() {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"user_id":    "foo",
		"exp":        time.Now().Add(time.Hour).Unix(),
	}).SignedString([]byte(SigningSecretKey))
	if err != nil {
		logrus.WithError(err).Error("Could not create a new token")
		return
	}
	logrus.WithField("token", token).Info("Create a new token")

	validToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SigningSecretKey), nil
	})
	if err != nil {
		logrus.WithError(err).Error("Invalid token")
		return
	}
	logrus.WithField("Claims", validToken.Claims).Info("Detected jwt token claims")
}
