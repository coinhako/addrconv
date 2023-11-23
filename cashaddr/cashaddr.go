package cashaddr

import (
	"errors"
	"fmt"

	"github.com/coinhako/addrconv/address"
	// "encoding/hex"
	// "encoding/binary"
)

// https://github.com/bitcoincashorg/spec/blob/master/cashaddr.md

// A cashaddr consists of three parts
// A prefix, "bitcoincash" for bch
// A separator, :
// A base32 encoded payload

// The payload further consists of 3 parts
// A Version Byte consisting of:
//     MSB is always 0
//     4 bits for address type (P2PKH/P2SH)
//     3 (LS)bits for the size of the hash (see spec link for size table)
// A hash (this is the actual address bit)
// A 40 bit checksum

const CHARSET string = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

// charset inversed for decoding
var CHARSET_REVERSED = [128]int8{
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 15, -1, 10, 17, 21, 20, 26, 30, 7,
	5, -1, -1, -1, -1, -1, -1, -1, 29, -1, 24, 13, 25, 9, 8, 23, -1, 18, 22,
	31, 27, 19, -1, 1, 0, 3, 16, 11, 28, 12, 14, 6, 4, 2, -1, -1, -1, -1,
	-1, -1, 29, -1, 24, 13, 25, 9, 8, 23, -1, 18, 22, 31, 27, 19, -1, 1, 0,
	3, 16, 11, 28, 12, 14, 6, 4, 2, -1, -1, -1, -1, -1,
}

/**
 * Concatenate two byte arrays.
 */
func Concat(x, y []byte) []byte {
	return append(x, y...)
}

/**
 * Convert to lower case.
 *
 * Assumes the input is a character.
 */
func makeLowerCase(c byte) byte {
	return c | 0x20
}

func packAddressData(addrType address.AddressType, addrHash []byte) ([]byte, error) {
	// Pack addr data with version byte.
	if addrType != address.P2PKH && addrType != address.P2SH {
		return []byte{}, errors.New("invalid addrtype")
	}

	var versionByte uint
	switch addrType {
	case address.P2PKH:
		versionByte = 0
	case address.P2SH:
		versionByte = 8
	}

	// hash can only be of 160 | 192 | 224 | 256 | 320 | 384 | 448 | 512 size
	size := len(addrHash)
	var encodedSize uint
	// Get the bit size, mostly for readability
	// against the spec. You could also just
	// compare to 20, 24 etc.
	switch size * 8 {
	case 160:
		encodedSize = 0
		break
	case 192:
		encodedSize = 1
		break
	case 224:
		encodedSize = 2
		break
	case 256:
		encodedSize = 3
		break
	case 320:
		encodedSize = 4
		break
	case 384:
		encodedSize = 5
		break
	case 448:
		encodedSize = 6
		break
	case 512:
		encodedSize = 7
		break
	default:
		return []byte{}, errors.New("invalid address size")
	}

	versionByte |= encodedSize // Just some OR
	var addrHashUint []byte
	for _, e := range addrHash {
		addrHashUint = append(addrHashUint, byte(e))
	}
	data := append([]byte{byte(versionByte)}, addrHashUint...)
	packedData, err := convertBits(data, 8, 5, true) // Convert 8-bit to 5-bit with 0-left padding
	if err != nil {
		return []byte{}, err
	}
	return packedData, nil
}

/**
 * Verify a checksum.
 */
func VerifyChecksum(prefix string, payload []byte) bool {
	return PolyMod(Concat(ExpandPrefix(prefix), payload)) == 0
}

/**
 * Create a checksum.
 */
func CreateChecksum(prefix string, payload []byte) []byte {
	// The checksum is computed over the address including
	// the prefix, so we compile the prefix and the payload
	// into a byte array
	// This prevents someone from sending to an invalid
	// address for their network even if it is displayed
	// without the prefix
	enc := Concat(ExpandPrefix(prefix), payload)

	// Append 8 zeroes.
	// This is the 8 bytes that'll hold the checksum
	enc = Concat(enc, []byte{0, 0, 0, 0, 0, 0, 0, 0})
	// Determine what to XOR into those 8 zeroes.
	mod := PolyMod(enc)    // Black magic
	ret := make([]byte, 8) // We need to convert it to 5-bit numbers
	for i := 0; i < 8; i++ {
		// Convert the 5-bit groups in mod to checksum values.
		ret[i] = byte((mod >> uint(5*(7-i))) & 0x1f)
	}
	return ret
}

/**
 * Encode a cashaddr string.
 */
func Encode(prefix string, payload []byte) string {
	checksum := CreateChecksum(prefix, payload)
	combined := Concat(payload, checksum)
	ret := prefix + ":"

	for _, c := range combined {
		ret += string(CHARSET[c])
	}

	return ret
}

func CheckEncodeCashAddress(input []byte, prefix string, t address.AddressType) string {
	k, err := packAddressData(t, input)
	if err != nil {
		fmt.Printf("%v", err)
		return ""
	}
	return Encode(prefix, k)
}

/**
 * Expand the address prefix for the checksum computation.
 */
func ExpandPrefix(prefix string) []byte {
	ret := make([]byte, len(prefix)+1)

	for i := 0; i < len(prefix); i++ {
		ret[i] = byte(prefix[i]) & 0x1f
	}

	ret[len(prefix)] = 0
	return ret
}

// CheckDecode decodes a string that was encoded with CheckEncode and verifies the checksum.
func CheckDecodeCashAddress(input string) (decodedAddress address.Address, err error) {
	prefix, data, err := DecodeCashAddress(input)
	if err != nil {
		return decodedAddress, err
	}
	data, err = convertBits(data, 5, 8, false)
	if err != nil {
		return decodedAddress, err
	}
	if len(data) != 21 {
		return decodedAddress, errors.New("Incorrect data length")
	}

	switch data[0] {
	case 0x00:
		decodedAddress.Type = address.P2PKH
	case 0x08:
		decodedAddress.Type = address.P2SH
	}

	decodedAddress.Hash = data[1:21]
	decodedAddress.CashAddrPrefix = prefix

	return decodedAddress, nil
}

/**
 * Decode a cashaddr string.
 */
func DecodeCashAddress(str string) (string, []byte, error) {
	// Go over the string and do some sanity checks.
	lower, upper := false, false
	prefixSize := 0
	for i := 0; i < len(str); i++ {
		c := byte(str[i])
		if c >= 'a' && c <= 'z' {
			lower = true
			continue
		}

		if c >= 'A' && c <= 'Z' {
			upper = true
			continue
		}

		if c >= '0' && c <= '9' {
			// We cannot have numbers in the prefix.
			if prefixSize == 0 {
				return "", []byte{}, errors.New("Addresses cannot have numbers in the prefix")
			}

			continue
		}

		if c == ':' {
			// The separator must not be the first character, and there must not
			// be 2 separators.
			if i == 0 || prefixSize != 0 {
				return "", []byte{}, errors.New("The separator must not be the first character")
			}

			prefixSize = i
			continue
		}

		// We have an unexpected character.
		return "", []byte{}, errors.New("Unexpected character")
	}

	// We must have a prefix and a data part and we can't have both uppercase
	// and lowercase.
	if prefixSize == 0 {
		return "", []byte{}, errors.New("Address must have a prefix")
	}

	if upper && lower {
		return "", []byte{}, errors.New("Addresses cannot use both upper and lower case characters")
	}

	// Get the prefix.
	var prefix string
	for i := 0; i < prefixSize; i++ {
		prefix += string(makeLowerCase(str[i]))
	}

	// Decode values.
	valuesSize := len(str) - 1 - prefixSize
	values := make([]byte, valuesSize)
	for i := 0; i < valuesSize; i++ {
		c := byte(str[i+prefixSize+1])
		// We have an invalid char in there.
		if c > 127 || CHARSET_REVERSED[c] == -1 {
			return "", []byte{}, errors.New("Invalid character")
		}

		values[i] = byte(CHARSET_REVERSED[c])
	}

	// Verify the checksum.
	if !VerifyChecksum(prefix, values) {
		return "", []byte{}, errors.New("Inavlid checksum")
	}

	return prefix, values[:len(values)-8], nil
}

// This is from the bech32 package, which is from github.com/sipa
func convertBits(data []byte, fromBits uint, tobits uint, pad bool) ([]byte, error) {
	// General power-of-2 base conversion.
	var uintArr []uint
	for _, i := range data {
		uintArr = append(uintArr, uint(i))
	}
	acc := uint(0)
	bits := uint(0)
	var ret []uint
	maxv := uint((1 << tobits) - 1)
	maxAcc := uint((1 << (fromBits + tobits - 1)) - 1)
	for _, value := range uintArr {
		acc = ((acc << fromBits) | value) & maxAcc
		bits += fromBits
		for bits >= tobits {
			bits -= tobits
			ret = append(ret, (acc>>bits)&maxv)
		}
	}
	if pad {
		if bits > 0 {
			ret = append(ret, (acc<<(tobits-bits))&maxv)
		}
	} else if bits >= fromBits || ((acc<<(tobits-bits))&maxv) != 0 {
		return []byte{}, errors.New("encoding padding error")
	}
	var dataArr []byte
	for _, i := range ret {
		dataArr = append(dataArr, byte(i))
	}
	return dataArr, nil
}

/**
 * This function will compute what 8 5-bit values to XOR into the last 8 input
 * values, in order to make the checksum 0. These 8 values are packed together
 * in a single 40-bit integer. The higher bits correspond to earlier values.
 */
func PolyMod(v []byte) uint64 {
	/**
	 * The input is interpreted as a list of coefficients of a polynomial over F
	 * = GF(32), with an implicit 1 in front. If the input is [v0,v1,v2,v3,v4],
	 * that polynomial is v(x) = 1*x^5 + v0*x^4 + v1*x^3 + v2*x^2 + v3*x + v4.
	 * The implicit 1 guarantees that [v0,v1,v2,...] has a distinct checksum
	 * from [0,v0,v1,v2,...].
	 *
	 * The output is a 40-bit integer whose 5-bit groups are the coefficients of
	 * the remainder of v(x) mod g(x), where g(x) is the cashaddr generator, x^8
	 * + {19}*x^7 + {3}*x^6 + {25}*x^5 + {11}*x^4 + {25}*x^3 + {3}*x^2 + {19}*x
	 * + {1}. g(x) is chosen in such a way that the resulting code is a BCH
	 * code, guaranteeing detection of up to 4 errors within a window of 1025
	 * characters. Among the various possible BCH codes, one was selected to in
	 * fact guarantee detection of up to 5 errors within a window of 160
	 * characters and 6 erros within a window of 126 characters. In addition,
	 * the code guarantee the detection of a burst of up to 8 errors.
	 *
	 * Note that the coefficients are elements of GF(32), here represented as
	 * decimal numbers between {}. In this finite field, addition is just XOR of
	 * the corresponding numbers. For example, {27} + {13} = {27 ^ 13} = {22}.
	 * Multiplication is more complicated, and requires treating the bits of
	 * values themselves as coefficients of a polynomial over a smaller field,
	 * GF(2), and multiplying those polynomials mod a^5 + a^3 + 1. For example,
	 * {5} * {26} = (a^2 + 1) * (a^4 + a^3 + a) = (a^4 + a^3 + a) * a^2 + (a^4 +
	 * a^3 + a) = a^6 + a^5 + a^4 + a = a^3 + 1 (mod a^5 + a^3 + 1) = {9}.
	 *
	 * During the course of the loop below, `c` contains the bitpacked
	 * coefficients of the polynomial constructed from just the values of v that
	 * were processed so far, mod g(x). In the above example, `c` initially
	 * corresponds to 1 mod (x), and after processing 2 inputs of v, it
	 * corresponds to x^2 + v0*x + v1 mod g(x). As 1 mod g(x) = 1, that is the
	 * starting value for `c`.
	 */
	c := uint64(1)
	for _, d := range v {
		/**
		 * We want to update `c` to correspond to a polynomial with one extra
		 * term. If the initial value of `c` consists of the coefficients of
		 * c(x) = f(x) mod g(x), we modify it to correspond to
		 * c'(x) = (f(x) * x + d) mod g(x), where d is the next input to
		 * process.
		 *
		 * Simplifying:
		 * c'(x) = (f(x) * x + d) mod g(x)
		 *         ((f(x) mod g(x)) * x + d) mod g(x)
		 *         (c(x) * x + d) mod g(x)
		 * If c(x) = c0*x^5 + c1*x^4 + c2*x^3 + c3*x^2 + c4*x + c5, we want to
		 * compute
		 * c'(x) = (c0*x^5 + c1*x^4 + c2*x^3 + c3*x^2 + c4*x + c5) * x + d
		 *                                                             mod g(x)
		 *       = c0*x^6 + c1*x^5 + c2*x^4 + c3*x^3 + c4*x^2 + c5*x + d
		 *                                                             mod g(x)
		 *       = c0*(x^6 mod g(x)) + c1*x^5 + c2*x^4 + c3*x^3 + c4*x^2 +
		 *                                                             c5*x + d
		 * If we call (x^6 mod g(x)) = k(x), this can be written as
		 * c'(x) = (c1*x^5 + c2*x^4 + c3*x^3 + c4*x^2 + c5*x + d) + c0*k(x)
		 */

		// First, determine the value of c0:
		c0 := byte(c >> 35)

		// Then compute c1*x^5 + c2*x^4 + c3*x^3 + c4*x^2 + c5*x + d:
		c = ((c & 0x07ffffffff) << 5) ^ uint64(d)

		// Finally, for each set bit n in c0, conditionally add {2^n}k(x):
		if c0&0x01 > 0 {
			// k(x) = {19}*x^7 + {3}*x^6 + {25}*x^5 + {11}*x^4 + {25}*x^3 +
			//        {3}*x^2 + {19}*x + {1}
			c ^= 0x98f2bc8e61
		}

		if c0&0x02 > 0 {
			// {2}k(x) = {15}*x^7 + {6}*x^6 + {27}*x^5 + {22}*x^4 + {27}*x^3 +
			//           {6}*x^2 + {15}*x + {2}
			c ^= 0x79b76d99e2
		}

		if c0&0x04 > 0 {
			// {4}k(x) = {30}*x^7 + {12}*x^6 + {31}*x^5 + {5}*x^4 + {31}*x^3 +
			//           {12}*x^2 + {30}*x + {4}
			c ^= 0xf33e5fb3c4
		}

		if c0&0x08 > 0 {
			// {8}k(x) = {21}*x^7 + {24}*x^6 + {23}*x^5 + {10}*x^4 + {23}*x^3 +
			//           {24}*x^2 + {21}*x + {8}
			c ^= 0xae2eabe2a8
		}

		if c0&0x10 > 0 {
			// {16}k(x) = {3}*x^7 + {25}*x^6 + {7}*x^5 + {20}*x^4 + {7}*x^3 +
			//            {25}*x^2 + {3}*x + {16}
			c ^= 0x1e4f43e470
		}
	}

	/**
	 * PolyMod computes what value to xor into the final values to make the
	 * checksum 0. However, if we required that the checksum was 0, it would be
	 * the case that appending a 0 to a valid list of values would result in a
	 * new valid list. For that reason, cashaddr requires the resulting checksum
	 * to be 1 instead.
	 */
	return c ^ 1
}
