package addrconv

import (
	"encoding/hex"
	"github.com/RaghavSood/blockutils"
	"testing"
)

func TestToAddress(t *testing.T) {
	var script blockutils.Script
	var err error
	script, err = hex.DecodeString("76a914bdb2b538e6b07e93d6bafcef4bec9dc936818a1988ac")
	if err != nil {
		t.Errorf("Error decoding hex: %s", err)
	}
	address, err := ToAddress(script)
	if err != nil {
		t.Errorf("Error encoding address: %s", err)
	}

	if address != "1JJ2o6iKB4UXVMHXBSzVvbAKim5su2VUfa" {
		t.Errorf("Incorrect address. Expected %s, got %s", "1JJ2o6iKB4UXVMHXBSzVvbAKim5su2VUfa", address)
	}
}

func TestDigibyteAddress(t *testing.T) {
	var script blockutils.Script
	var err error
	script, err = hex.DecodeString("76a914510fffca0668d410aea742e95a2fefa7952f695e88ac")
	if err != nil {
		t.Errorf("Error decoding hex: %s", err)
	}
	address, err := ToNetworkAddress(script, DigibyteNetwork)
	if err != nil {
		t.Errorf("Error encoding address: %s", err)
	}

	if address != "DCXiSSQwi7gw9YXrMY4mxt2i4hQZEBb5Yv" {
		t.Errorf("Incorrect address. Expected %s, got %s", "DCXiSSQwi7gw9YXrMY4mxt2i4hQZEBb5Yv", address)
	}
}

func TestBTCP2SHAddress(t *testing.T) {
	var script blockutils.Script
	var err error
	script, err = hex.DecodeString("a9144aef67ed61d391d6f3d9903ead92386c1efc992587")
	if err != nil {
		t.Errorf("Error decoding hex: %s", err)
	}
	address, err := ToNetworkAddress(script, BitcoinNetwork)
	if err != nil {
		t.Errorf("Error encoding address: %s", err)
	}

	if address != "38XEixUj1QpcqxTWbxvqdbv4Mjre4imw9Z" {
		t.Errorf("Incorrect address. Expected %s, got %s", "38XEixUj1QpcqxTWbxvqdbv4Mjre4imw9Z", address)
	}
}

func TestOmniOpReturnScript(t *testing.T) {
	var script blockutils.Script
	var err error
	script, err = hex.DecodeString("6a146f6d6e69000000000000000300000000000000c8")
	if err != nil {
		t.Errorf("Error decoding hex: %s", err)
	}
	address, err := ToNetworkAddress(script, BitcoinNetwork)
	if err != nil {
		t.Errorf("Error encoding address: %s", err)
	}

	if address != "6a146f6d6e69000000000000000300000000000000c8" {
		t.Errorf("Incorrect address. Expected %s, got %s", "6a146f6d6e69000000000000000300000000000000c8", address)
	}
}

func TestBech32Addresses(t *testing.T) {
	var script blockutils.Script
	var err error

	var scripts = []string{"0014751e76e8199196d454941c45d1b3a323f1433bd6", "0020701a8d401c84fb13e6baf169d59684e17abd9fa216c8cc5b9fc63d622ff8c58d", "00203eb5062a0b0850b23a599425289a091c374ca934101d03144f060c5b46a979be"}
	var addresses = []string{"bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4", "bc1qwqdg6squsna38e46795at95yu9atm8azzmyvckulcc7kytlcckxswvvzej", "bc1q866sv2stppgtywjejsjj3xsfrsm5e2f5zqwsx9z0qcx9k34f0xlqaslnwr"}

	for i, v := range scripts {

		script, err = hex.DecodeString(v)
		if err != nil {
			t.Errorf("Error decoding hex: %s", err)
		}
		address, err := ToNetworkAddress(script, BitcoinNetwork)
		if err != nil {
			t.Errorf("Error encoding address: %s", err)
		}

		if address != addresses[i] {
			t.Errorf("Incorrect address. Expected %s, got %s", addresses[i], address)
		}
	}
}

func TestBTCP2PK(t *testing.T) {
	var script blockutils.Script
	var err error

	var scripts = []string{"4104f601d3111e0f502f8d5927fd4077e8723f9b156138a776afa5ceae4d6da7370d4f1effb4171cd8ce7b40d54e3a0b45b528575ce63465986810085babdef06f01ac", "2103f601d3111e0f502f8d5927fd4077e8723f9b156138a776afa5ceae4d6da7370dac", "410411db93e1dcdb8a016b49840f8c53bc1eb68a382e97b1482ecad7b148a6909a5cb2e0eaddfb84ccf9744464f82e160bfa9b8b64f9d4c03f999b8643f656b412a3ac"}
	var addresses = []string{"18AuG6THx5V9JuZbtWtrabYWfxY28YcXhN", "1ABTTFaiYdTNX4yzQ1vn5QP1URrHwCDbFc", "12cbQLTFMXRnSzktFkuoG3eHoMeFtpTu3S"}

	for i, v := range scripts {

		script, err = hex.DecodeString(v)
		if err != nil {
			t.Errorf("Error decoding hex: %s", err)
		}
		address, err := ToNetworkAddress(script, BitcoinNetwork)
		if err != nil {
			t.Errorf("Error encoding address: %s", err)
		}

		if address != addresses[i] {
			t.Errorf("Incorrect address. Expected %s, got %s", addresses[i], address)
		}
	}
}

func TestOPReturn(t *testing.T) {
	var script blockutils.Script
	var err error

	var scripts = []string{"6a21e5b88ce69c9be4bda0e698afe4b880e4b8aae69c89e59381e591b3e79a84e4baba"}
	var addresses = []string{"6a21e5b88ce69c9be4bda0e698afe4b880e4b8aae69c89e59381e591b3e79a84e4baba"}

	for i, v := range scripts {

		script, err = hex.DecodeString(v)
		if err != nil {
			t.Errorf("Error decoding hex: %s", err)
		}
		address, err := ToNetworkAddress(script, BitcoinNetwork)
		if err != nil {
			t.Errorf("Error encoding address: %s", err)
		}

		if address != addresses[i] {
			t.Errorf("Incorrect address. Expected %s, got %s", addresses[i], address)
		}
	}
}

func TestDecodeAddress(t *testing.T) {
	var scripts = []string{"f7b11fa0d7cad927d47183e14580fd63418d77e5", "c91b3f1306c9b7b4c0fcfc4ac9a3b22d20145b19"}
	var addresses = []string{"1Pag69EdPN95wGYbKUDb2YxPz9QJbskeTh", "3L2NPDLsrqCycMt1Q7t9fLWqEmPKrnVWT1"}
	var versions = []byte{0x00, 0x05}

	for i, v := range addresses {
		addressScript, version, err := FromAddress(v)
		if err != nil {
			t.Errorf("Error encoding address: %s", err)
		}
		script := blockutils.Script(addressScript)
		if script.String() != scripts[i] {
			t.Errorf("Incorrect address. Expected %s, got %s", scripts[i], addressScript)
		}

		if version != versions[i] {
			t.Errorf("Incorrect address version. Expected %#x, got %#x", versions[i], version)
		}
	}
}
