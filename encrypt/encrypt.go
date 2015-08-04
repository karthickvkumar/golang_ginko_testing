package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func EncryptAes(text string) string {
	plaintext := []byte(text)
	key_text := "12345678912345678912345678912345"
	// Create the aes encryption algorithm
	c, _ := aes.NewCipher([]byte(key_text))
	// Encrypted string
	e := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	e.XORKeyStream(ciphertext, plaintext)
	//fmt.Printf("%s=>%x\n", plaintext, ciphertext)
	// Decrypt strings
	d := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(plaintext))
	d.XORKeyStream(plaintextCopy, ciphertext)
	// fmt.Printf("%x=>%s\n", ciphertext, plaintextCopy)
	return string(plaintextCopy)
}
