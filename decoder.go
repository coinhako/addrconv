package addrconv

import (
	"github.com/RaghavSood/addrconv/base58"
	// "github.com/RaghavSood/blockutils"
)

func (network Network) Decode(encodedAddress string) ([]byte, byte, error) {
	addressScript, version, err := base58.CheckDecode(encodedAddress)
	if err != nil {
		return []byte{}, 0xFF, err
	}

	return addressScript, version, nil
}
