package main

import (
	"bufio"
	"fmt"
	"os"
	"sonik/function"
	"strconv"
	"strings"
)

func main() {
	// Get the message to be encrypted from the user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the message to be encrypted: ")
	messageStr, _ := reader.ReadString('\n')
	messageStr = strings.TrimSpace(messageStr)
	message, _ := strconv.Atoi(messageStr)

	// Get the public key from the user
	fmt.Print("Enter the public key (e, n), separated by a space: ")
	keyStr, _ := reader.ReadString('\n')
	keyStr = strings.TrimSpace(keyStr)
	key := strings.Split(keyStr, " ")
	e, _ := strconv.Atoi(key[0])
	n, _ := strconv.Atoi(key[1])

	// Calculate the modular multiplicative inverse of e modulo phi(n)
	p := 61
	q := 53
	phi := (p - 1) * (q - 1)
	d := function.ModInverse(e, phi)

	// Encrypt the message using the public key
	ciphertext := function.Encrypt(message, e, n)

	// Decrypt the ciphertext using the private key
	plaintext := function.Decrypt(ciphertext, d, n)

	// Print the results
	fmt.Println("Public key (e, n):", e, n)
	fmt.Println("Private key (d, n):", d, n)
	fmt.Println("Original message:", message)
	fmt.Println("Encrypted message:", ciphertext)
	fmt.Println("Decrypted message:", plaintext)
}
