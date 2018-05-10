package address

type AddressType int

const (
	UNKNOWN     AddressType = 0
	P2PKH       AddressType = 1
	P2SH        AddressType = 2
	P2SH_P2WPKH AddressType = 3
	P2SH_P2WSH  AddressType = 4
	P2WPKH      AddressType = 5
	P2WSH       AddressType = 6
	P2PK        AddressType = 7
)

type Address struct {
	Type           AddressType
	Hash           []byte
	Bech32HRP      string
	CashAddrPrefix string
}

func (address Address) IsP2SH() bool {
	return address.Type == P2SH || address.Type == P2SH_P2WPKH || address.Type == P2SH_P2WSH
}
