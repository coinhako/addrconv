package addrconv

import (
	"github.com/RaghavSood/blockutils"
)

func (network Network) Encode(script blockutils.Script) (string, error) {
	if script.IsP2PKH() {
		hash160, err := script.P2PKHHash160()
		if err != nil {
			return "", err
		}

		return CheckEncode(hash160, network.PubKeyPrefix), nil
	}
	return "test", nil
}
