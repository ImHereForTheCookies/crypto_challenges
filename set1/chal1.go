package set1

import (
	"encoding/hex"
	b64 "encoding/base64"
)

func chal1(input string) string {
	raw_input, _ := hex.DecodeString(input)
	return b64.StdEncoding.EncodeToString(raw_input)
}