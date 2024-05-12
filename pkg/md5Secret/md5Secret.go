package md5secret

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Secret(data, secret string) string {
	h := md5.New()
	h.Write([]byte(data))
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum(nil))
}
