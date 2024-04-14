package ranhash

import "math/rand"

// HashPool has alphabets and numbers
const HashPool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// NumberPool has only numbers
const NumberPool = "0123456789"

// AlphabetPool has only alphabets
const AlphabetPool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Generate a Random String from a pool, with given length
func Generate(length int, pool string) string {
	var otp string

	for i := 0; i < length; i++ {
		otp += string(pool[rand.Intn(len(pool))])
	}

	return otp
}

// RanHash Generates a random hash with a given length, both alphabets and numbers
// are included
func RanHash(length int) string {
	pool := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return Generate(length, pool)
}

// GenerateRandomString generates a random string of alphabets of given length
func GenerateRandomString(length int) string {
	pool := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	return Generate(length, pool)
}

// GenerateRandomNumber generates a random string of numbers of length len
func GenerateRandomNumber(length int) string {
	pool := "0123456789"

	return Generate(length, pool)
}
