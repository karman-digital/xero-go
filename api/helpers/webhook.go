package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"hash"

	xeroerrors "github.com/karman-digital/xero-go/app/errors"
)

func ValidateWebhookSignature(method, signature string, secret, body []byte) error {
	hash := encrypt(secret, body)
	fmt.Println("hash: ", hash)
	fmt.Println("signature: ", signature)
	if signature != base64.StdEncoding.EncodeToString(hash.Sum(nil)) {
		return xeroerrors.ErrInvalidSignature
	}
	return nil
}

func encrypt(secret []byte, input []byte) hash.Hash {
	hash := hmac.New(sha256.New, secret)
	hash.Write(input)
	return hash
}
