package addrconv

type Network struct {
	Bech32Prefix     string // Human readable part of bech32 addresses
	PubKeyPrefix     byte   // P2PKH address prefix
	ScriptHashPrefix byte   // P2SH address prefix
	WIFPrefix        byte   // wif key prefix
	BIP32PubPrefix   []byte // extended public key prefix
	BIP32PrivPrefix  []byte // extended private key prefix
}

var BitcoinNetwork = Network{
	Bech32Prefix:     "bc",
	PubKeyPrefix:     0x00,
	ScriptHashPrefix: 0x05,
	WIFPrefix:        0x80,
	BIP32PubPrefix:   []byte{0x04, 0x88, 0xb2, 0x1e},
	BIP32PrivPrefix:  []byte{0x04, 0x88, 0xad, 0xe4},
}

var DigibyteNetwork = Network{
	Bech32Prefix:     "dgb",
	PubKeyPrefix:     0x1e,
	ScriptHashPrefix: 0x3f,
	WIFPrefix:        0x9e,
	BIP32PubPrefix:   []byte{0x04, 0x88, 0xb2, 0x1e},
	BIP32PrivPrefix:  []byte{0x04, 0x88, 0xad, 0xe4},
}
