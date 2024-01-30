package generate

import (
	"math/rand"
	"time"
)

func RandomString(length int) string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)

	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}

	return string(result)
}
