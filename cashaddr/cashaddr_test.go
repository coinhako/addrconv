package cashaddr

import (
	"encoding/hex"
	"testing"

	"github.com/coinhako/addrconv/address"
	"github.com/coinhako/blockutils"
)

func TestCheckEncodeCashAddress(t *testing.T) {
	var script blockutils.Script
	var err error
	script, err = hex.DecodeString("6fd39d9bbb63afe8ad3c0ec984bc101a0b5a88d6")
	if err != nil {
		t.Errorf("Error decoding hex: %s", err)
	}
	encodedAddress := CheckEncodeCashAddress(script, "bitcoincash", address.P2PKH)

	if encodedAddress != "bitcoincash:qpha88vmhd36l69d8s8vnp9uzqdqkk5g6cnfvrsf5l" {
		t.Errorf("Incorrect address. Expected %s, got %s", "bitcoincash:qpha88vmhd36l69d8s8vnp9uzqdqkk5g6cnfvrsf5l", encodedAddress)
	}

	script, err = hex.DecodeString("4aef67ed61d391d6f3d9903ead92386c1efc9925")
	if err != nil {
		t.Errorf("Error decoding hex: %s", err)
	}
	encodedAddress = CheckEncodeCashAddress(script, "bitcoincash", address.P2SH)

	if encodedAddress != "bitcoincash:pp9w7eldv8fer4hnmxgratvj8pkpalyey5qym9j8x5" {
		t.Errorf("Incorrect address. Expected %s, got %s", "bitcoincash:pp9w7eldv8fer4hnmxgratvj8pkpalyey5qym9j8x5", encodedAddress)
	}
}
