package auth

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"os"
)

var SecretKey = os.Getenv("SECRET_KEY")

// hmacSha1 calculates HMAC-SHA1 digest
func hmacSha1(key []byte, data []byte) string {
	mac := hmac.New(sha1.New, key)
	mac.Write(data)
	return hex.EncodeToString(mac.Sum(nil))
}

// checkAuth checks if the request is authorized
func checkAuth(r *http.Request) bool {
	userID := r.Header.Get("X-UserId")
	if userID == "" {
		return false
	}
	digest := r.Header.Get("X-Digest")
	if digest == "" {
		return false
	}
	requestBody, _ := ioutil.ReadAll(r.Body)
	expectedDigest := hmacSha1([]byte(SecretKey), requestBody)
	return digest == expectedDigest
}
