package ibc

import (
	"github.com/andrey-solomenniy-test/standarta/sdk/wire"
)

// Register concrete types on wire codec
func RegisterWire(cdc *wire.Codec) {
	//cdc.RegisterConcrete(IBCTransferMsg{}, "github.com/andrey-solomenniy-test/standarta/sdk/x/ibc/IBCTransferMsg", nil)
	//cdc.RegisterConcrete(IBCReceiveMsg{}, "github.com/andrey-solomenniy-test/standarta/sdk/x/ibc/IBCReceiveMsg", nil)
}
