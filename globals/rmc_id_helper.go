package globals

import (
	globals_rmc "github.com/PretendoNetwork/yo-kai-watch-blasters/globals/rmc"
)

func GetProtocolByID(protocolId uint16) globals_rmc.ProtocolInfo {
	switch protocolId {
	case 3:
		return globals_rmc.NewNATTraversal()
	case 10:
		return globals_rmc.NewTicketGranting()
	case 11:
		return globals_rmc.NewSecureConnection()
	case 21:
		return globals_rmc.NewMatchMaking()
	case 27:
		return globals_rmc.NewMessageDelivery()
	case 50:
		return globals_rmc.NewMatchMakingExtension()
	case 109:
		return globals_rmc.NewMatchmakeExtension()
	}

	return globals_rmc.NewProtocolInfo()
}
