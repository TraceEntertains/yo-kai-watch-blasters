package globals

import (
	"database/sql"

	pbaccount "github.com/PretendoNetwork/grpc/go/account"
	pbfriends "github.com/PretendoNetwork/grpc/go/friends"
	"github.com/PretendoNetwork/nex-go/v2"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	"github.com/PretendoNetwork/plogger-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var Postgres *sql.DB
var MatchmakingManager *common_globals.MatchmakingManager
var MessagingManager *common_globals.MessagingManager
var Logger *plogger.Logger
var KerberosPassword = "password" // * Default password

var AuthenticationServer *nex.PRUDPServer
var AuthenticationEndpoint *nex.PRUDPEndPoint

var SecureServer *nex.PRUDPServer
var SecureEndpoint *nex.PRUDPEndPoint
var AESKey []byte

var GRPCAccountClientConnection *grpc.ClientConn
var GRPCAccountClient pbaccount.AccountClient
var GRPCAccountCommonMetadata metadata.MD

var GRPCFriendsClientConnection *grpc.ClientConn
var GRPCFriendsClient pbfriends.FriendsClient
var GRPCFriendsCommonMetadata metadata.MD
