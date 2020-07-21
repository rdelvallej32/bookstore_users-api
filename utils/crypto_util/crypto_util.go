package crypto_utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(input string) string {
	hash := md5.New()

	defer hash.Reset()

	_, err := hash.Write([]byte(input))

	if err != nil {
		return "error"
	}

	return hex.EncodeToString(hash.Sum(nil))
}
