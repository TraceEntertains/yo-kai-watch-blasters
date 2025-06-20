package nex

import (
	//"fmt"

	"context"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	commonnattraversal "github.com/PretendoNetwork/nex-protocols-common-go/v2/nat-traversal"
	commonsecure "github.com/PretendoNetwork/nex-protocols-common-go/v2/secure-connection"
	nattraversal "github.com/PretendoNetwork/nex-protocols-go/v2/nat-traversal"
	secure "github.com/PretendoNetwork/nex-protocols-go/v2/secure-connection"
	"github.com/PretendoNetwork/yo-kai-watch-blasters/globals"

	commonmatchmaking "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making"
	commonmatchmakingext "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making-ext"
	commonmatchmakeextension "github.com/PretendoNetwork/nex-protocols-common-go/v2/matchmake-extension"
	commonmessagedelivery "github.com/PretendoNetwork/nex-protocols-common-go/v2/message-delivery"
	matchmaking "github.com/PretendoNetwork/nex-protocols-go/v2/match-making"
	matchmakingext "github.com/PretendoNetwork/nex-protocols-go/v2/match-making-ext"
	matchmakeextension "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension"

	matchmakingtypes "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
	messagedelivery "github.com/PretendoNetwork/nex-protocols-go/v2/message-delivery"
	ranking "github.com/PretendoNetwork/nex-protocols-go/v2/ranking"

	local_match_making "github.com/PretendoNetwork/yo-kai-watch-blasters/nex/match_making"
)

// Is this needed? -Ash
func cleanupSearchMatchmakeSessionHandler(matchmakeSession *matchmakingtypes.MatchmakeSession) {
	matchmakeSession.Attributes[2] = types.NewUInt32(0)
	matchmakeSession.MatchmakeParam = matchmakingtypes.NewMatchmakeParam()
	matchmakeSession.ApplicationBuffer = types.NewBuffer(make([]byte, 0))
	globals.Logger.Info(matchmakeSession.String())
}

func CreateReportDBRecord(_ types.PID, _ types.UInt32, _ types.QBuffer) error {
	return nil
}

func registerCommonSecureServerProtocols() {
	secureProtocol := secure.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(secureProtocol)
	commonSecureProtocol := commonsecure.NewCommonProtocol(secureProtocol)
	commonSecureProtocol.EnableInsecureRegister()

	globals.MatchmakingManager.GetUserFriendPIDs = globals.GetUserFriendPIDs

	commonSecureProtocol.CreateReportDBRecord = CreateReportDBRecord

	natTraversalProtocol := nattraversal.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(natTraversalProtocol)
	commonnattraversal.NewCommonProtocol(natTraversalProtocol)

	matchMakingProtocol := matchmaking.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchMakingProtocol)
	matchMakingProtocol.FindByOwner = local_match_making.FindByOwner
	commonmatchmaking.NewCommonProtocol(matchMakingProtocol).SetManager(globals.MatchmakingManager)

	matchMakingExtProtocol := matchmakingext.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchMakingExtProtocol)
	commonmatchmakingext.NewCommonProtocol(matchMakingExtProtocol).SetManager(globals.MatchmakingManager)

	rankingProtocol := ranking.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(rankingProtocol)

	matchmakeExtensionProtocol := matchmakeextension.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchmakeExtensionProtocol)
	commonMatchmakeExtensionProtocol := commonmatchmakeextension.NewCommonProtocol(matchmakeExtensionProtocol)
	commonMatchmakeExtensionProtocol.SetManager(globals.MatchmakingManager)

	messageDeliveryProtocol := messagedelivery.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(messageDeliveryProtocol)
	commonmessagedelivery.NewCommonProtocol(messageDeliveryProtocol).SetManager(globals.MessagingManager)

	commonMatchmakeExtensionProtocol.CleanupSearchMatchmakeSession = cleanupSearchMatchmakeSessionHandler
	commonMatchmakeExtensionProtocol.CleanupMatchmakeSessionSearchCriterias = func(searchCriterias types.List[matchmakingtypes.MatchmakeSessionSearchCriteria]) {
		for _, searchCriteria := range searchCriterias {
			searchCriteria.Attribs[2] = types.NewString("")
		}
	}

	commonMatchmakeExtensionProtocol.OnAfterCreateMatchmakeSession = func(packet nex.PacketInterface, anyGathering matchmakingtypes.GatheringHolder, message types.String, participationCount types.UInt16) {
		globals.MatchmakingManager.Database.ExecContext(context.Background(), "UPDATE matchmaking.matchmake_sessions SET user_password_enabled=false WHERE user_password=''")
	}

}
