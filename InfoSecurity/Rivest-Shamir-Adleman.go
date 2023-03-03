package infosecurity

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func RSAAlgo(){
	
	// Generate a key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Get the public key
	publicKey := privateKey.PublicKey

	// Encrypt the message using the public key
	message := []byte("Hello, world!")
	label := []byte("")
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, &publicKey, message, label)
	if err != nil {
		panic(err)
	}

	// Decrypt the message using the private key
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, ciphertext, label)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Original message: %s\n", message)
	fmt.Printf("Decrypted message: %s\n", plaintext)
}

