package globals

import (
	"context"

	"github.com/PretendoNetwork/nex-go/v2/types"

	"time"

	pb "github.com/PretendoNetwork/grpc/go/account"
	"github.com/PretendoNetwork/nex-go/v2"
	"google.golang.org/grpc/metadata"
)

type UserDataCacheEntry struct {
	userData     *pb.GetUserDataResponse
	creationTime time.Time
	updatedTime  time.Time
}

var UserDataCache map[types.PID]UserDataCacheEntry = map[types.PID]UserDataCacheEntry{}

func UserDataFromPID(pid types.PID) (*pb.GetUserDataResponse, uint32) {

	data, exists := UserDataCache[pid]
	if !exists || data.updatedTime.Add(time.Hour*24).Before(time.Now().UTC()) {
		ctx := metadata.NewOutgoingContext(context.Background(), GRPCAccountCommonMetadata)

		response, err := GRPCAccountClient.GetUserData(ctx, &pb.GetUserDataRequest{Pid: uint32(pid)})
		if err != nil {
			Logger.Error(err.Error())
			if !exists {
				UserDataCache[pid] = UserDataCacheEntry{userData: nil, creationTime: time.Now().UTC(), updatedTime: time.Now().UTC()}
			} else {
				data.userData = nil
				data.updatedTime = time.Now().UTC()

				UserDataCache[pid] = data
			}

			return &pb.GetUserDataResponse{}, nex.ResultCodes.RendezVous.InvalidUsername
		}

		data.userData = response
		data.updatedTime = time.Now().UTC()

		UserDataCache[pid] = data
	}

	if data.userData != nil {
		return data.userData, 0
	}

	// this should only happen if there is some inexplicable failure in fetching the user data (or its a 3ds server)
	return nil, nex.ResultCodes.Core.Unknown
}
