package addrconv

import (
	"github.com/RaghavSood/blockutils"
	"math/big"
)

var bigRadix = big.NewInt(58)
var bigZero = big.NewInt(0)

// The Base58 checksum is the first four bytes of sha256(sha256(data))
func checksum(input []byte) (chksum [4]byte) {
	doubleSha := blockutils.DoubleSha256(input)
	copy(chksum[:], doubleSha[:4])
	return
}

func encode(b []byte) string {
	x := new(big.Int)
	x.SetBytes(b)

	answer := make([]byte, 0, len(b)*136/100)
	for x.Cmp(bigZero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, bigRadix, mod)
		answer = append(answer, alphabet[mod.Int64()])
	}

	// leading zero bytes
	for _, i := range b {
		if i != 0 {
			break
		}
		answer = append(answer, alphabetAtIndexZero)
	}

	// reverse
	alen := len(answer)
	for i := 0; i < alen/2; i++ {
		answer[i], answer[alen-1-i] = answer[alen-1-i], answer[i]
	}

	return string(answer)
}

// CheckEncode calculates the checksum and encodes the input with the
// provided prefix from a network
func CheckEncode(input []byte, version byte) string {
	b := make([]byte, 0, 1+len(input)+4)
	b = append(b, version)
	b = append(b, input[:]...)
	chksum := checksum(b)
	b = append(b, chksum[:]...)
	return encode(b)
}
