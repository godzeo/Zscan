package util

import (
	"math/rand"
	"strings"
)

// TrimProtocol removes the HTTP scheme from an URI
func TrimProtocol(targetURL string) string {
	URL := strings.TrimSpace(targetURL)
	if strings.HasPrefix(strings.ToLower(URL), "http://") || strings.HasPrefix(strings.ToLower(URL), "https://") {
		URL = URL[strings.Index(URL, "//")+2:]
	}
	URL = strings.TrimRight(URL, "/")
	return URL
}
func RandomStr(randSource *rand.Rand, letterBytes string, n int) string {
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
		//letterBytes   = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	)
	randBytes := make([]byte, n)
	for i, cache, remain := n-1, randSource.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSource.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			randBytes[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(randBytes)
}
func StrInSlice(i string, array []string) bool {
	ret := false
	for _, a := range array {
		if i == a {
			ret = true
			break
		}
	}
	return ret
}
