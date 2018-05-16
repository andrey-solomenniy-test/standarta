package bank

import (
	"github.com/andrey-solomenniy-test/standarta/sdk/wire"
)

// Register concrete types on wire codec
func RegisterWire(cdc *wire.Codec) {
	// TODO include option to always include prefix bytes.
	//cdc.RegisterConcrete(SendMsg{}, "github.com/andrey-solomenniy-test/standarta/sdk/bank/SendMsg", nil)
	//cdc.RegisterConcrete(IssueMsg{}, "github.com/andrey-solomenniy-test/standarta/sdk/bank/IssueMsg", nil)
}
