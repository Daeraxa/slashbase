package utils

import (
	"encoding/hex"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"unsafe"
)

func ContainsString(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func ContainsInt(a []int, integer int) bool {
	for _, v := range a {
		if v == integer {
			return true
		}
	}
	return false
}

func UnixNanoToTime(nanoInt int64) time.Time {
	msInt := nanoInt / 1000000000
	remainder := nanoInt % 1000000000
	value := time.Unix(msInt, remainder*int64(time.Nanosecond))
	return value
}

func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// fast & unsafe pointer function
func RandString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

func GetRequestScheme(r *http.Request) string {
	if r.TLS != nil {
		return "https"
	}
	return "http"
}

func GetRequestCookieHost(r *http.Request) string {
	if strings.HasPrefix(r.Host, "localhost:") {
		return "localhost"
	}
	return GetRequestScheme(r) + "://" + r.Host
}
