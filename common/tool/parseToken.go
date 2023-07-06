package tool

import "github.com/golang-jwt/jwt/v4"

//解析token
func ParseToken(token string, secret string) (*jwt.Token, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return claim, nil
}
