package api

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func Encode(text []byte) string {
	return base64.StdEncoding.EncodeToString(text)
}

func Decode(text string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(text)
}

func Encrypt(text, secret string) (string, error) {
	textToEncrypt := []byte(text)
	keyForEncryption := []byte(secret) // 32 bytes
	c, err := aes.NewCipher(keyForEncryption)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	encText := gcm.Seal(nonce, nonce, textToEncrypt, nil)
	return Encode(encText), nil
}

func Decrypt(secret, key string) ([]byte, error) {
	encText, err := Decode(secret)
	if err != nil {
		return nil, err
	}
	keyForDecryption := []byte(key) // 32 bytes
	c, err := aes.NewCipher(keyForDecryption)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(encText) < nonceSize {
		return nil, err
	}
	nonce, encText := encText[:nonceSize], encText[nonceSize:]
	return gcm.Open(nil, nonce, encText, nil)
}
