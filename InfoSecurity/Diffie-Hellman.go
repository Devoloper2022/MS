package infosecurity

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)


func DHAlgo(){

reader := bufio.NewReader(os.Stdin)

	// Read in the prime number p
	fmt.Print("Enter a prime number p: ")
	pStr, _ := reader.ReadString('\n')
	pStr = strings.TrimSpace(pStr)
	p, _ := strconv.ParseInt(pStr, 10, 64)

	// Read in the primitive root g
	fmt.Print("Enter a primitive root g: ")
	gStr, _ := reader.ReadString('\n')
	gStr = strings.TrimSpace(gStr)
	g, _ := strconv.ParseInt(gStr, 10, 64)

	// Alice chooses a secret number
	fmt.Print("Enter Alice's secret number: ")
	aStr, _ := reader.ReadString('\n')
	aStr = strings.TrimSpace(aStr)
	a, _ := strconv.ParseInt(aStr, 10, 64)

	// Bob chooses a secret number
	fmt.Print("Enter Bob's secret number: ")
	bStr, _ := reader.ReadString('\n')
	bStr = strings.TrimSpace(bStr)
	b, _ := strconv.ParseInt(bStr, 10, 64)

	// Alice computes A = g^a mod p and sends A to Bob
	A := new(big.Int).Exp(big.NewInt(g), big.NewInt(a), big.NewInt(p))

	// Bob computes B = g^b mod p and sends B to Alice
	B := new(big.Int).Exp(big.NewInt(g), big.NewInt(b), big.NewInt(p))

	// Alice computes the shared secret key K = B^a mod p
	Ka := new(big.Int).Exp(B, big.NewInt(a), big.NewInt(p))

	// Bob computes the shared secret key K = A^b mod p
	Kb := new(big.Int).Exp(A, big.NewInt(b), big.NewInt(p))

	fmt.Printf("Alice's secret number: %v\n", a)
	fmt.Printf("Bob's secret number: %v\n", b)

	// Read in the message to encrypt
	fmt.Print("Enter the message to encrypt: ")
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)

	// Encrypt the message using the shared secret key
	ciphertext := encrypt(message, Ka)

	fmt.Printf("Encrypted message: %v\n", ciphertext)

	// Decrypt the message using the shared secret key
	plaintext := decrypt(ciphertext, Kb)

	fmt.Printf("Decrypted message: %v\n", plaintext)
}

// Encrypts a message using a shared secret key
func encrypt(message string, key *big.Int) string {
	// Convert the message to a byte array
	messageBytes := []byte(message)

	// XOR each byte of the message with the corresponding byte of the key
	for i := 0; i < len(messageBytes); i++ {
		messageBytes[i] ^= byte(key.Uint64())
	}

	// Convert the encrypted byte array to a string
	return string(messageBytes)
}

// Decrypts a message using a shared secret key
func decrypt(ciphertext string, key *big.Int) string {
	// Convert the ciphertext to a byte array
	ciphertextBytes := []byte(ciphertext)

	// XOR each byte of the ciphertext with the corresponding byte of the key
	for i := 0; i < len(ciphertextBytes); i++ {
		ciphertextBytes[i] ^= byte(key.Uint64())
	}

	// Convert the decrypted byte array to a string
	return string(ciphertextBytes)
}