package xcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
)

func Decrypt(key, msg string) (plaintext string, err error) {

	keyhash := sha256.Sum256([]byte(key))
	var block cipher.Block
	block, err = aes.NewCipher(keyhash[:])
	if err != nil {
		return
	}

	var aead cipher.AEAD
	aead, err = cipher.NewGCM(block)
	if err != nil {
		return
	}

	var ciphertext []byte
	ciphertext, err = base64.RawURLEncoding.DecodeString(msg)
	if err != nil {
		return
	}

	var bytes []byte
	bytes, err = aead.Open(nil, ciphertext[:12], ciphertext[12:], nil)
	if err != nil {
		return
	}

	plaintext = string(bytes)

	return
}
