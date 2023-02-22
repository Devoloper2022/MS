package function

import "math/big"

// Function to find the greatest common divisor of two numbers
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to find the modular multiplicative inverse of a modulo m
func ModInverse(a, m int) int {
	gcd, s, _ := ExtendedEuclidean(a, m)
	if gcd != 1 {
		panic("a and m are not coprime")
	} else {
		return (s%m + m) % m
	}
}

// Function to perform the extended Euclidean algorithm
func ExtendedEuclidean(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	} else {
		gcd, x, y := ExtendedEuclidean(b, a%b)
		return gcd, y, x - (a/b)*y
	}
}

// Function to encrypt a message using the public key (e, n)
func Encrypt(message, e, n int) int {
	ciphertext := new(big.Int).Exp(big.NewInt(int64(message)), big.NewInt(int64(e)), big.NewInt(int64(n)))
	return int(ciphertext.Int64())
}

// Function to decrypt a ciphertext using the private key (d, n)
func Decrypt(ciphertext, d, n int) int {
	plaintext := new(big.Int).Exp(big.NewInt(int64(ciphertext)), big.NewInt(int64(d)), big.NewInt(int64(n)))
	return int(plaintext.Int64())
}
