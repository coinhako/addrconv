package addrconv

import (
	"encoding/hex"
	"github.com/RaghavSood/addrconv/address"
	"testing"
)

func TestEncodeToBase58(t *testing.T) {
	var decodedAddress address.Address

	decodedAddress.Hash, _ = hex.DecodeString("bdb2b538e6b07e93d6bafcef4bec9dc936818a19")
	decodedAddress.Type = address.P2PKH

	encodedAddress, err := BitcoinNetwork.EncodeToBase58(decodedAddress)
	if err != nil {
		t.Errorf("Error encoding address: %s", err)
	}

	if encodedAddress != "1JJ2o6iKB4UXVMHXBSzVvbAKim5su2VUfa" {
		t.Errorf("Incorrect address. Expected %s, got %s", "1JJ2o6iKB4UXVMHXBSzVvbAKim5su2VUfa", encodedAddress)
	}

	decodedAddress.Hash, _ = hex.DecodeString("4aef67ed61d391d6f3d9903ead92386c1efc9925")
	decodedAddress.Type = address.P2SH

	encodedAddress, err = BitcoinNetwork.EncodeToBase58(decodedAddress)
	if err != nil {
		t.Errorf("Error encoding address: %s", err)
	}

	if encodedAddress != "38XEixUj1QpcqxTWbxvqdbv4Mjre4imw9Z" {
		t.Errorf("Incorrect address. Expected %s, got %s", "38XEixUj1QpcqxTWbxvqdbv4Mjre4imw9Z", encodedAddress)
	}

	decodedAddress.Hash, _ = hex.DecodeString("b619de6e0a35d6d4f9ec93c77f23784dd1388971")
	decodedAddress.Type = address.P2PKH

	encodedAddress, err = ZcoinNetwork.EncodeToBase58(decodedAddress)
	if err != nil {
		t.Errorf("Error encoding address: %s", err)
	}

	if encodedAddress != "aHKKiDdEAYQjjbEgJMSUpdkapz4hVUUCHR" {
		t.Errorf("Incorrect address. Expected %s, got %s", "aHKKiDdEAYQjjbEgJMSUpdkapz4hVUUCHR", encodedAddress)
	}
}

func TestEncodeToCashAddr(t *testing.T) {
	var decodedAddress address.Address

	decodedAddress.Hash, _ = hex.DecodeString("bdb2b538e6b07e93d6bafcef4bec9dc936818a19")
	decodedAddress.Type = address.P2PKH

	encodedAddress, err := BitcoinCashNetwork.EncodeToCashAddr(decodedAddress)
	if err != nil {
		t.Errorf("Error encoding address: %s", err)
	}

	if encodedAddress != "bitcoincash:qz7m9dfcu6c8ay7kht7w7jlvnhyndqv2rya0nn54z0" {
		t.Errorf("Incorrect address. Expected %s, got %s", "bitcoincash:qz7m9dfcu6c8ay7kht7w7jlvnhyndqv2rya0nn54z0", encodedAddress)
	}

	decodedAddress.Hash, _ = hex.DecodeString("4aef67ed61d391d6f3d9903ead92386c1efc9925")
	decodedAddress.Type = address.P2SH

	encodedAddress, err = BitcoinCashNetwork.EncodeToCashAddr(decodedAddress)
	if err != nil {
		t.Errorf("Error encoding address: %s", err)
	}

	if encodedAddress != "bitcoincash:pp9w7eldv8fer4hnmxgratvj8pkpalyey5qym9j8x5" {
		t.Errorf("Incorrect address. Expected %s, got %s", "bitcoincash:pp9w7eldv8fer4hnmxgratvj8pkpalyey5qym9j8x5", encodedAddress)
	}
}
