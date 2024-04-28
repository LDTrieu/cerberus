package jwt

import "github.com/golang-jwt/jwt"

func IsExpiredJWTError(err error) bool {
	if err == nil {
		return false
	}

	switch v := err.(type) {
	case *jwt.ValidationError:
		return (v != nil) && (v.Errors&jwt.ValidationErrorExpired != 0)
	default:
		return false
	}
}
