package addrconv

import (
	"github.com/RaghavSood/addrconv/base58"
	"github.com/RaghavSood/addrconv/bech32"
	"github.com/RaghavSood/blockutils"
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
