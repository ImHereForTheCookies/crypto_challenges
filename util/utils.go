package util

import (
	"encoding/hex"
)


func xor_bytes(arr1, arr2 []byte) []byte {
	result_arr := arr1
	for i, b in range arr2 {
		result_arr[i] ^= b
	}
}

func from_hex()