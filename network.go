package addrconv

type Network struct {
	Bech32Prefix     string // Human readable part of bech32 addresses
	PubKeyPrefix     byte   // P2PKH address prefix
	ScriptHashPrefix byte   // P2SH address prefix
	WIFPrefix        byte   // wif key prefix
	BIP32PubPrefix   []byte // extended public key prefix
	BIP32PrivPrefix  []byte // extended private key prefix
}
