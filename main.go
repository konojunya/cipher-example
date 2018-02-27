package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func main() {
	// 暗号化したい文字列
	plainText := []byte("My name is konojunya")

	// aesの暗号文字列
	keyText := "astaxie12798akljzmsknm.ahkjkljl;"

	// 暗号化アルゴリズムaesを作成
	c, err := aes.NewCipher([]byte(keyText))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(keyText), err)
		os.Exit(1)
	}

	fmt.Printf("Plain Text: %s\n", plainText)

	// 暗号化文字列
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	fmt.Printf("Encryption string: %x\n", cipherText)

	// 復号文字列
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plainTextCopy := make([]byte, len(plainText))
	cfbdec.XORKeyStream(plainTextCopy, cipherText)
	fmt.Printf("Decryption string: %s\n", plainTextCopy)

}
