package utility

import (
	"crypto/rand"
	"regexp"
)

func GenerateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6
	b := make([]byte, length)
	rand.Read(b)
	for i, v := range b {
		b[i] = charset[v%byte(len(charset))]
	}
	return string(b)
}

func IsOwnDomain(url string, domain string) bool {
	r, err := regexp.Compile("https?://" + domain + "/")
	if err != nil {
		return false
	}

	return r.MatchString(url)
}
