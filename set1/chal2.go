package set1

import (
	"encoding/hex"
	"github.com/ImHereForTheCookies/crypto_challenges/util"
)

func Chal2(input, key string) string {
	raw_input, _ := hex.DecodeString(input)
	raw_key, _ := hex.DecodeString(key)
	return hex.StdEncoding.EncodeToString(util.xor_bytes(raw_key, raw_input))
}