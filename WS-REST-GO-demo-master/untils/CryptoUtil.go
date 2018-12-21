package untils

import (
	"github.com/dgrijalva/jwt-go"
)

func SignByJWT(pem string, data string) (string, error) {
	key, err := jwt.ParseECPrivateKeyFromPEM([]byte(pem))
	if nil != err {
		return "", err
	}

	alg := jwt.GetSigningMethod("ES256")
	return alg.Sign(data, key)
}
