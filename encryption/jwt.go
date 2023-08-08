package encryption

import (
	"max-inventory/internal/models"

	"github.com/golang-jwt/jwt/v4"
)

func SignedLoginToken(u *models.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"name":  u.Name,
	})

	jwtString, err := token.SignedString([]byte(key))
	if err!= nil {
        return "", err
    }

	return jwtString, nil
}

func ParseLoginJWT(value string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(value, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err	
	}

	return token.Claims.(jwt.MapClaims), nil
}