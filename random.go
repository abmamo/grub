package main

import (
	"math/rand"
	"time"
)

// seed
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// alphabet
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// string generator (with charset)
func stringWithCharset(length int, charset string) string {
	// make array
	b := make([]byte, length)
	// iteratively fill array
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	// return string
	return string(b)
}

// RandomString : function that takes a int length and generates string of size length
func RandomString(length int) string {
	return stringWithCharset(length, charset)
}
