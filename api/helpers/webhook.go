package helpers

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"hash"

	xeroerrors "github.com/karman-digital/xero-go/app/errors"
)

func ValidateWebhookSignature(method, signature string, secret, body []byte) error {
	checkSum, err := toCompactJSONString(body)
	if err != nil {
		return err
	}
	hash := encrypt(secret, checkSum)
	if signature != base64.StdEncoding.EncodeToString(hash.Sum(nil)) {
		return xeroerrors.ErrInvalidSignature
	}
	return nil
}

func encrypt(secret []byte, input string) hash.Hash {
	hash := hmac.New(sha256.New, secret)
	hash.Write([]byte(input))
	return hash
}

func toCompactJSONString(input []byte) (string, error) {
	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, input); err != nil {
		return "", err
	}
	return buffer.String(), nil
}
