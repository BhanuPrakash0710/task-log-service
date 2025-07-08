package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
)

func GenerateUserId(name string) string {
	// Generate a random number
	randomNumber := rand.Intn(1000)

	// Create a unique identifier by combining the name and random number
	uniqueIdentifier := name + strconv.Itoa(randomNumber)

	// Hash the unique identifier using MD5
	hash := md5.Sum([]byte(uniqueIdentifier))

	// Convert the hash to a hexadecimal string
	userId := hex.EncodeToString(hash[5:11])

	return userId
}
