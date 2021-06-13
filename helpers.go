package indodana

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"strconv"
	"time"
)

func Stringify(m interface{}) string {
	json, _ := json.Marshal(m)
	return string(json)
}

func Sha1(s string) string {
	h := sha1.New()
	io.WriteString(h, s)
	bs := h.Sum(nil)

	result := fmt.Sprintf("%x", bs)

	return result
}

func CreateSignatureHeaders(apiKey string, secretKey string) string {
	nonce := math.Floor(float64(time.Now().UTC().UnixNano()) / 1000)
	key := fmt.Sprintf("%s:%s", apiKey, strconv.Itoa(int(nonce)))
	byteKey := []byte(key)

	mac := hmac.New(sha256.New, []byte(secretKey))
	// Cannot return error
	if _, err := mac.Write(byteKey); err != nil {
		return ""
	}

	hash := hex.EncodeToString(mac.Sum(nil))
	signature := fmt.Sprintf("%s:%s", key, hash)
	return signature
}
