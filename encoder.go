package addrconv

import (
	"errors"
	"fmt"

	"github.com/coinhako/addrconv/address"
	"github.com/coinhako/addrconv/base58"
	"github.com/coinhako/addrconv/bech32"
	"github.com/coinhako/addrconv/cashaddr"
	"github.com/coinhako/blockutils"
)

func (network Network) Encode(script blockutils.Script) (string, error) {
	if script.IsOpReturn() {
		return script.String(), nil
	}

	if script.IsP2PK() {
		hash160, err := script.P2PKHash160()
		if err != nil {
			return script.String(), err
		}
		return base58.CheckEncode(hash160, network.PubKeyPrefix), nil
	}

	if script.IsP2PKH() {
		hash160, err := script.P2PKHHash160()
		if err != nil {
			return script.String(), err
		}
		return base58.CheckEncode(hash160, network.PubKeyPrefix), nil
	}

	if script.IsP2SH() {
		hash160, err := script.P2SHHash160()
		if err != nil {
			return script.String(), err
		}
		return base58.CheckEncode(hash160, network.ScriptHashPrefix), nil
	}

	if script.IsWitnessScript() {
		witnessVersion, err := script.WitnessVersion()
		if err != nil {
			return script.String(), err
		}
		intWitnessVersion := int(witnessVersion)
		witnessProgram, err := script.WitnessProgram()
		if err != nil {
			return script.String(), err
		}
		intWitnessProgram, err := toIntSlice(witnessProgram)
		if err != nil {
			return script.String(), err
		}
		return bech32.SegwitAddrEncode(network.Bech32Prefix, intWitnessVersion, intWitnessProgram)
	}

	return script.String(), nil
}

func toIntSlice(buf []byte) ([]int, error) {
	vals := make([]int, len(buf))
	for i := 0; i < len(vals); i++ {
		vals[i] = int(buf[i])
	}
	return vals, nil
}

func (network Network) EncodeToBase58(decodedAddress address.Address) (string, error) {

	if decodedAddress.Type == address.P2PKH {
		return base58.CheckEncode(decodedAddress.Hash, network.PubKeyPrefix), nil
	}

	if decodedAddress.IsP2SH() {
		return base58.CheckEncode(decodedAddress.Hash, network.ScriptHashPrefix), nil
	}

	return "", fmt.Errorf("Unknown address %d type for base58", decodedAddress.Type)

}

func (network Network) EncodeToCashAddr(decodedAddress address.Address) (encodedAddress string, err error) {
	if !network.SupportsCashAddr() {
		err = errors.New("Network does not support cashaddr")
		return encodedAddress, err
	}

	if decodedAddress.Type != address.P2SH && decodedAddress.Type != address.P2PKH {
		err = errors.New("cashaddr only supports P2SH and P2PKH addresses")
		return encodedAddress, err
	}

	decodedAddress.CashAddrPrefix = network.CashAddrPrefix
	encodedAddress = cashaddr.CheckEncodeCashAddress(decodedAddress.Hash, decodedAddress.CashAddrPrefix, decodedAddress.Type)
	return encodedAddress, nil
}
