package addrconv

import (
	"github.com/RaghavSood/blockutils"
)

func ToAddress(script blockutils.Script) (string, error) {
	return ToNetworkAddress(script, BitcoinNetwork)
}

func ToNetworkAddress(script blockutils.Script, network Network) (string, error) {
	return network.Encode(script)
}

func FromAddress(address string) ([]byte, byte, error) {
	return FromNetworkAddress(address, BitcoinNetwork)
}

func FromNetworkAddress(address string, network Network) ([]byte, byte, error) {
	return network.Decode(address)
}
