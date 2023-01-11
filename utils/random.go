package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	var sb strings.Builder
	sb.Grow(10)
	l := len(alphabet)

	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		c := alphabet[rand.Intn(l)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomInteger(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min+1)
}

func RandomEmail() string {
	s := RandomString(10)
	domain := "@email.com"
	return s + domain
}

func RandomUser(n int) string {
	rand.Seed(time.Now().UnixNano())
	return RandomString(n)
}

func RandomPassword(n int) string {
	rand.Seed(time.Now().UnixNano())
	return RandomString(n)
}
