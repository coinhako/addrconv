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
