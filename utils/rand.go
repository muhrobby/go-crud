package utils

import (
	"math/rand"
	"strconv"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_1234567890")

func Init() string {
	timeUnix := rand.NewSource(time.Now().UnixMicro())
	timeC := strconv.FormatInt(timeUnix.Int63(), 10)

	return timeC
}

func RandomString(n int) string {

	b := make([]rune, n)

	for i := range b {

		b[i] = letters[rand.Intn(len(letters))]

	}
	return string(b)

}
