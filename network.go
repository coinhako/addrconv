package addrconv

import (
	"strings"
)

type Network struct {
	Bech32Prefix     string // Human readable part of bech32 addresses
	PubKeyPrefix     byte   // P2PKH address prefix
	ScriptHashPrefix byte   // P2SH address prefix
	WIFPrefix        byte   // wif key prefix
	BIP32PubPrefix   []byte // extended public key prefix
	BIP32PrivPrefix  []byte // extended private key prefix
	CashAddrPrefix   string //cashaddr prefix
}

var BitcoinNetwork = Network{
	Bech32Prefix:     "bc",
	PubKeyPrefix:     0x00,
	ScriptHashPrefix: 0x05,
	WIFPrefix:        0x80,
	BIP32PubPrefix:   []byte{0x04, 0x88, 0xb2, 0x1e},
	BIP32PrivPrefix:  []byte{0x04, 0x88, 0xad, 0xe4},
}

var BitcoinCashNetwork = Network{
	PubKeyPrefix:     0x00,
	ScriptHashPrefix: 0x05,
	WIFPrefix:        0x80,
	BIP32PubPrefix:   []byte{0x04, 0x88, 0xb2, 0x1e},
	BIP32PrivPrefix:  []byte{0x04, 0x88, 0xad, 0xe4},
	CashAddrPrefix:   "bitcoincash",
}

var DigibyteNetwork = Network{
	Bech32Prefix:     "dgb",
	PubKeyPrefix:     0x1e,
	ScriptHashPrefix: 0x3f,
	WIFPrefix:        0x9e,
	BIP32PubPrefix:   []byte{0x04, 0x88, 0xb2, 0x1e},
	BIP32PrivPrefix:  []byte{0x04, 0x88, 0xad, 0xe4},
}

var LitecoinNetwork = Network{
	Bech32Prefix:     "ltc",
	PubKeyPrefix:     0x30,
	ScriptHashPrefix: 0x32,
	WIFPrefix:        0xb0,
	BIP32PubPrefix:   []byte{0x04, 0x88, 0xb2, 0x1e},
	BIP32PrivPrefix:  []byte{0x04, 0x88, 0xad, 0xe4},
}

var ZcoinNetwork = Network{
	PubKeyPrefix:     0x52,
	ScriptHashPrefix: 0x07,
	WIFPrefix:        0xd2,
	BIP32PubPrefix:   []byte{0x04, 0x88, 0xb2, 0x1e},
	BIP32PrivPrefix:  []byte{0x04, 0x88, 0xad, 0xe4},
}

var DogecoinNetwork = Network{
	PubKeyPrefix:     0x1e,
	ScriptHashPrefix: 0x16,
	WIFPrefix:        0x9e,
	BIP32PubPrefix:   []byte{0x02, 0xfa, 0xca, 0xfd},
	BIP32PrivPrefix:  []byte{0x02, 0xfa, 0xc3, 0x98},
}

// Returns the predefined network settings for common coins
// based on the provided coin name
func GetNetwork(name string) Network {
	name = strings.ToLower(name)
	if name == "bitcoin" {
		return BitcoinNetwork
	}

	if name == "digibyte" {
		return DigibyteNetwork
	}

	if name == "litecoin" {
		return LitecoinNetwork
	}

	if name == "bitcoincash" {
		return BitcoinCashNetwork
	}

	if name == "zcoin" {
		return ZcoinNetwork
	}

	if name == "dogecoin" {
		return DogecoinNetwork
	}

	return BitcoinNetwork
}

func GetNetworkByTicker(ticker string) Network {
	name := strings.ToLower(ticker)
	if name == "btc" {
		return BitcoinNetwork
	}

	if name == "dgb" {
		return DigibyteNetwork
	}

	if name == "ltc" {
		return LitecoinNetwork
	}

	if name == "bch" {
		return BitcoinCashNetwork
	}

	if name == "xzc" {
		return ZcoinNetwork
	}

	if name == "doge" {
		return DogecoinNetwork
	}

	return BitcoinNetwork
}

func (network Network) SupportsCashAddr() bool {
	return network.CashAddrPrefix != ""
}

func (network Network) SupportsBech32() bool {
	return network.Bech32Prefix != ""
}
