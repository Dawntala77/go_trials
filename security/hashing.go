package security

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

func Hash(key string, password string) string {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}
	msg := make([]byte, len(password))
	c.Encrypt(msg, []byte(password))
	return hex.EncodeToString(msg)
}

func Digest(key string, hashed_password string) string {
	text, _ := hex.DecodeString(hashed_password)
	c, err := aes.NewCipher([]byte(key))

	if err != nil {
		fmt.Println(err)
	}
	msg := make([]byte, len(text))
	c.Decrypt(msg, []byte(text))
	msgbyte := string(msg[:])
	return msgbyte
}

// func main() {

// 	plainText := "This is a secret"
// 	key := "this_must_be_of_32_byte_length!!"

// 	emsg := hash(key, plainText)
// 	dmesg := digest(key, emsg)

// 	fmt.Println("Encrypted Message: ", emsg)
// 	fmt.Println("Decrypted Message: ", dmesg)

// }
