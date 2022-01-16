package xcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

func Encrypt(key, msg string) (ciphertext string, err error) {

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

	nonce := [12]byte{}
	_, err = rand.Read(nonce[:])
	if err != nil {
		return
	}

	bytes := aead.Seal(nonce[:], nonce[:], []byte(msg), nil)
	ciphertext = base64.RawURLEncoding.EncodeToString(bytes)

	return
}
