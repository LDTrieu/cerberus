package jwt

type Option func(v *verifier)

func WithPublicKeyBase64(keyBase64 string) Option {
	return func(v *verifier) {
		v.pubKey64 = keyBase64
	}
}

func WithSecretKey(key string) Option {
	return func(v *verifier) {
		v.secretKey = key
	}
}
