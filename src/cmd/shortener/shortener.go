package shortener

import (
	"crypto/md5"
	"encoding/hex"
)

func Shorten(url string) (string, error) {
	// naive hash
	hash := md5.Sum([]byte(url))
	short := hex.EncodeToString(hash[:])[:7]
	return short, nil
}
