package encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

const crypt_key = "xiaomiNB8hellomi"

func Encrypt(text string) string {
	c, err := aes.NewCipher([]byte(crypt_key))
	if err != nil {
		fmt.Println(err)
	}
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	crypt_text := make([]byte, len(text))
	cfb.XORKeyStream(crypt_text, []byte(text))
	return base64.StdEncoding.EncodeToString(crypt_text)
}

func Decrypt(crypt_text string) string {
	crypt_text_byte, _ := base64.StdEncoding.DecodeString(crypt_text)
	c, err := aes.NewCipher([]byte(crypt_key))
	if err != nil {
		fmt.Println(err)
	}
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	text := make([]byte, 100)
	cfbdec.XORKeyStream(text, crypt_text_byte)
	var text2 []byte
	for index, v := range text {
		if v != 0 {
			text2 = append(text2, text[index])
			continue
		}
	}
	return string(text2)
}
