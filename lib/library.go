package lib

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/rs/xid"
)

type PaginationDTO struct {
	BeforeCursor string `json:"before_cursor"`
	AfterCursor  string `json:"after_cursor"`
}

func GenerateID() string {
	guid := xid.New()
	return guid.String()
}

func EncodeHMACSHA256(text string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(text))
	sha256 := hex.EncodeToString(h.Sum(nil))
	return sha256
}

func VerifySignature(text string, signature string, salt string) bool {
	mac := hmac.New(sha256.New, []byte(salt))
	mac.Write([]byte(text))
	ExpectedMAC := mac.Sum(nil)
	SignatureHMAC, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}
	return hmac.Equal(ExpectedMAC, SignatureHMAC)
}

func BytesToString(data []byte) string {
	return string(data[:])
}
