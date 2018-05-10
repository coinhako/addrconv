package addrconv

import (
	"errors"
	"github.com/RaghavSood/addrconv/address"
	"github.com/RaghavSood/addrconv/base58"
	"github.com/RaghavSood/addrconv/cashaddr"
	"strings"
)

func (network Network) Decode(encodedAddress string) (address.Address, error) {
	return network.decodeByGuessingEncoding(encodedAddress)
}

func (network Network) decodeByGuessingEncoding(encodedAddress string) (decodedAddress address.Address, err error) {
	// Let's try base58 first, since it's the most common, and more
	// or less all networks support it

	decodedAddress, err = base58.CheckDecode(encodedAddress)
	if err == nil { // Decoding was successful, we're done
		return decodedAddress, nil
	}

	if network.SupportsCashAddr() {
		encodedCashAddress := encodedAddress
		if !strings.HasPrefix(encodedAddress, network.CashAddrPrefix+":") {
			encodedCashAddress = network.CashAddrPrefix + ":" + encodedAddress
		}

		decodedAddress, err = cashaddr.CheckDecodeCashAddress(encodedCashAddress)
		if err == nil {
			return decodedAddress, nil
		}
	}

	return decodedAddress, errors.New("Unknown address type")
}
