package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/ldtrieu/cerberus/config"

	"github.com/golang-jwt/jwt"
)

type jwtClaims struct {
	jwt.StandardClaims
	UserID int64 `json:"sub"` // for map with claims inside
	Email  string
}

func GenerateToken(user_id int64, email string, config *config.JwtKeyConfig) (signedToken string, err error) {
	token := ""
	now := time.Now().UTC()
	claims := &jwtClaims{
		UserID: user_id,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(time.Hour * time.Duration(config.TTL)).Unix(),
			// Issuer:    w.Issuer,
		},
	}

	if config.SecretKey != "" {
		key := []byte(config.SecretKey)
		token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
		if err != nil {
			return "", fmt.Errorf("create: sign token: %w", err)
		}
	} else {
		if len(config.AuthPrivateKey1) == 0 {
			return "", fmt.Errorf("read: auth private key : not config")
		}
		private_key, err := os.ReadFile(config.AuthPrivateKey1)
		if err != nil {
			return "", fmt.Errorf("read: auth private key: %w", err)
		}
		key, err := jwt.ParseRSAPrivateKeyFromPEM(private_key)
		if err != nil {
			return "", fmt.Errorf("create: parse key: %w", err)
		}
		token, err = jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)

		if err != nil {
			return "", fmt.Errorf("create: sign token: %w", err)
		}
	}

	return token, nil
}

func ValidateToken(tokenString string, config *config.JwtKeyConfig) (*jwtClaims, error) {
	var token *jwt.Token
	var err error
	if config.SecretKey != "" {
		key := []byte(config.SecretKey)
		token, err = jwt.ParseWithClaims(
			tokenString,
			&jwtClaims{},
			func(jwtToken *jwt.Token) (interface{}, error) {
				if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
				}
				return key, nil
			},
		)
		if err != nil {
			return nil, fmt.Errorf("failed to parse token: %w", err)
		}
	} else {
		if len(config.AuthPublicKey1) == 0 {
			return nil, fmt.Errorf("read: auth public key 1: not config")
		}
		public_key, err := os.ReadFile(config.AuthPublicKey1)
		if err != nil {
			return nil, fmt.Errorf("read: auth public key 1: %w", err)
		}
		key, err := jwt.ParseRSAPublicKeyFromPEM(public_key)
		if err != nil {
			// try again with PublicKey2 if configured
			if len(config.AuthPublicKey2) > 0 {
				public_key, err = os.ReadFile(config.AuthPublicKey2)
				if err != nil {
					return nil, fmt.Errorf("read: auth public key 2: %w", err)
				}
				key, err = jwt.ParseRSAPublicKeyFromPEM(public_key)
				if err != nil {
					return nil, fmt.Errorf("failed to parse auth public key 2: %w", err)
				}
			} else {
				return nil, fmt.Errorf("failed to parse auth public key: %w", err)
			}
		}

		token, err = jwt.ParseWithClaims(
			tokenString,
			&jwtClaims{},
			func(jwtToken *jwt.Token) (interface{}, error) {
				if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
				}
				return key, nil
			},
		)
		if err != nil {
			return nil, fmt.Errorf("failed to parse token: %w", err)
		}
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return nil, fmt.Errorf("JWT is expired")
	}

	return claims, nil
}
