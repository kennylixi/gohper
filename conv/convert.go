package conv

import (
	"encoding/hex"
	"strconv"
	"unsafe"

	"github.com/cosiner/gohper/errors"
)

const (
	HEX  = "0123456789ABCDEF"
	LHEX = "0123456789abcdef"
)

func Uint2Hex(u uint64) []byte {
	s := make([]byte, 16)

	var idx int
	for idx = 15; idx > -1 && u > 0; idx-- {
		s[idx] = HEX[u&0xF]
		u = u >> 4
	}

	return s[idx+1:]
}

func Uint2LowerHex(u uint64) []byte {
	s := make([]byte, 16)

	var idx int
	for idx = 15; idx > -1 && u > 0; idx-- {
		s[idx] = LHEX[u&0xF]
		u = u >> 4
	}

	return s[idx+1:]
}

// Hex2Uint convert a hexadecimal string to uint
// if string is invalid, return an error
func Hex2Uint(str string) (uint64, error) {
	if len(str) > 2 {
		if head := str[:2]; head == "0x" || head == "0X" {
			str = str[2:]
		}
	}

	var n uint64
	for _, c := range str {
		if c >= '0' && c <= '9' {
			c = c - '0'
		} else if c >= 'a' && c <= 'f' {
			c = c - 'a' + 10
		} else if c >= 'A' && c <= 'F' {
			c = c - 'A' + 10
		} else {
			return 0, errors.Newf("Invalid hexadecimal string %s", str)
		}
		n = n << 4
		n |= uint64(c)
	}

	return n, nil
}

// BytesToHex transfer binary to hex bytes
func Bytes2Hex(src []byte) []byte {
	dst := make([]byte, 2*len(src))
	hex.Encode(dst, src)

	return dst
}

// HexToBytes transfer hex bytes to binary
func Hex2Bytes(src []byte) []byte {
	dst := make([]byte, len(src)/2)
	hex.Decode(dst, src)

	return dst
}

// ReverseBits reverse all bits in number
func ReverseBits(num uint) uint {
	var n uint
	size := uint(unsafe.Sizeof(num))
	for s := size * 8; s > 0; s-- {
		n = n << 1
		n |= (num & 1)
		num = num >> 1
	}

	return n
}

// ReverseByte reverse all bits for a byte
func ReverseByte(num byte) byte {
	var n byte
	for s := 8; s > 0; s-- {
		n = n << 1
		n |= (num & 1)
		num = num >> 1
	}

	return n
}

func IntDef(val string, def int) int {
	if val == "" {
		return def
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		i = def
	}

	return i
}
