package jwt

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"

	"github.com/ldtrieu/cerberus/package/errors"

	"github.com/golang-jwt/jwt"
)

var (
	_ Verifier = (*verifier)(nil)
)

type Verifier interface {
	Verify(string, jwt.Claims) error
}

// verifier ...
type verifier struct {
	pubKey64 string

	publicKey *rsa.PublicKey
	secretKey string
}

// NewVerifier ...
func NewVerifier(opts ...Option) (Verifier, error) {
	v := &verifier{}
	for _, opt := range opts {
		opt(v)
	}

	if v.pubKey64 != "" {
		decoded, err := base64.StdEncoding.DecodeString(v.pubKey64)
		if err != nil {
			return nil, err
		}

		publicKey, err := jwt.ParseRSAPublicKeyFromPEM(decoded)
		if err != nil {
			return nil, err
		}
		v.publicKey = publicKey
	}

	if v.secretKey == "" && v.publicKey == nil {
		return nil, errors.New("publicKey or secretKey required")
	}

	return v, nil
}

func (v *verifier) verifyPubKeyFunc() jwt.Keyfunc {
	fn := func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}
		return v.publicKey, nil
	}
	return fn
}

func (v *verifier) verifySecretKeyFunc() jwt.Keyfunc {
	fn := func(jwtToken *jwt.Token) (interface{}, error) {
		key := []byte(v.secretKey)
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}
		return key, nil
	}
	return fn
}

// Verify ...
func (v *verifier) Verify(signed string, claims jwt.Claims) error {

	if v.secretKey != "" {
		_, err := jwt.ParseWithClaims(signed, claims, v.verifySecretKeyFunc())
		if err != nil {
			return err
		}

		return nil
	}

	_, err := jwt.ParseWithClaims(signed, claims, v.verifyPubKeyFunc())
	if err != nil {
		return err
	}

	return nil
}
