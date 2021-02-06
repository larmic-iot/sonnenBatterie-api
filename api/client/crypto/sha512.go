package crypto

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
)

func Encrypt(data, salt string) string {
	sha512Bytes := sha512.Sum512([]byte(data))
	sha512BytesInHex := []byte(hex.EncodeToString(sha512Bytes[:]))
	key := pbkdf2.Key(sha512BytesInHex, []byte(salt), 7500, 64, sha512.New)
	return fmt.Sprintf("%x", key)
}
