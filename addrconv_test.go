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
