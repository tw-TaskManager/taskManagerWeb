package tokenValidator

import (
	"github.com/dgrijalva/jwt-go"
	"encoding/base64"
)

func validate(originalToken string) (string, bool) {
	data, err := jwt.Parse(originalToken, func(token *jwt.Token) (interface{}, error) {
		key := "TM2016"
		decoded, _ := base64.URLEncoding.DecodeString(key)
		return decoded, nil
	})
	if err != nil {
		return "", false
	}
	userId := data.Claims["Id"]
	return userId.(string), true
}
