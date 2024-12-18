package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"os"
)

// Função para encriptar uma string
func encrypt(key string, nonce string, str string) (string, error) {
	// key := []byte("blogPostGeekHunterblogPostGeekHu")
	// plaintext := []byte("Este é o texto plano a ser cifrado")
	keyb := []byte(key)
	plaintext := []byte(str)
	block, err := aes.NewCipher(keyb)
	if err != nil {
		return "", err
	}
	// nonce := []byte("blogPostGeek")
	nonceb := []byte(nonce)
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	ciphertext := aesgcm.Seal(nil, nonceb, plaintext, nil)
	fmt.Printf("Ciphertext: %x\n", ciphertext)
	return fmt.Sprintf("%x", ciphertext), nil
}

func decrypt(key string, nonce string, ciphertext string) (string, error) {
	// key := []byte("blogPostGeekHunterblogPostGeekHu")
	// ciphertext, _ := hex.DecodeString("839a362d3d61d1c2987c73d751058a50dfe39bcd06cd0cdbcd29f3b0a06368874aa2ba8fe7c9f386d77ed5a554e0f17307ec9b")
	// nonce := []byte("blogPostGeek")
	keyb := []byte(key)
	ciphert, _ := hex.DecodeString(ciphertext)
	nonceb := []byte(nonce)
	block, err := aes.NewCipher(keyb)
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	plaintext, err := aesgcm.Open(nil, nonceb, ciphert, nil)
	if err != nil {
		return "", err
	}
	fmt.Printf("Plaintext: %s\n", string(plaintext))
	return fmt.Sprintf("%s", plaintext), nil
}

func main() {

	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(host)

	c, err := encrypt("blogPostGeekHunterblogPostGeekHu", "blogPostGeek", "Este é o texto plano a ser cifrado")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(c)
	s, err := decrypt("blogPostGeekHunterblogPostGeekHu", "blogPostGeek", "839a362d3d61d1c2987c73d751058a50dfe39bcd06cd0cdbcd29f3b0a06368874aa2ba8fe7c9f386d77ed5a554e0f17307ec9b")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(s)

	fmt.Println("---------------------------------------------------")

	c, err = encrypt("Bueno                           ", "FullControl ", "FullControl sistemas legais")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(c)
	fmt.Println("---------------------------------------------------")
	// s = decrypt("Bueno                           ", "FullControl ", "bfe41a161d9ffb688ad5b4c81cc5e4abb02a002cab33e5473ffe57c891d429bf32d73caed1a53b4fe74a2a")
	s, err = decrypt("Bueno                           ", "FullControl ", "bfe41a161d9ffb688ad5b4c81cc5e4abb02a002cab33e5473ffe57c891d429bf32d73caed1a53b4fe74a2a")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(s)

	fmt.Println(fmt.Sprintf("%32s", "1234567890123456789012345678901234567890"))
	fmt.Println(fmt.Sprintf("%32s", "bueno"))

}
