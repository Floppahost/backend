package lib

import (
	"math/rand"
	"time"
)

func Random(length int) string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range chars {
		result[i] = chars[seededRand.Intn(len(chars))]
	}
	return string(result)
}

func Token() string {
	return (Random(15))
}
