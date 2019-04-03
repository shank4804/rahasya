package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func encryptFile(plainTextfilePath string, passphrase string) error {
	plainData, err := ioutil.ReadFile(plainTextfilePath)
	if err != nil {
		return err
	}
	f, err := os.Create("encrypted_file.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	f.Write(encrypt(plainData, passphrase))
	return nil
}

func decryptFile(encryptedFilePath string, passphrase string) error {
	encryptedData, err := ioutil.ReadFile(encryptedFilePath)
	if err != nil {
		return err
	}
	f, err := os.Create("decrypted_file.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	f.Write(decrypt(encryptedData, passphrase))
	return nil
}

func main() {
	typePtr := flag.String("t", "", "operation to encrypt / decrypt (enc | dec)")
	passwordPtr := flag.String("p", "", "password to encrypt / decrypt")
	filePath := flag.String("f", "", "filepath to do the type operation")

	flag.Parse()

	if *typePtr == "" || *passwordPtr == "" || *filePath == "" {
		fmt.Printf("Not enough parameters\n")
	} else if *typePtr == "enc" {
		err := encryptFile(*filePath, *passwordPtr)
		if err != nil {
			panic(err)
		}
	} else if *typePtr == "dec" {
		err := decryptFile(*filePath, *passwordPtr)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("Unsupported operation: %s\n", *typePtr)
	}
}
