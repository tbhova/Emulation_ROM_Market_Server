package ServerManager

import (
	"google.golang.org/grpc"
	GameServer "AvailableGamesServer"
	"golang.org/x/net/context"
)

// Data type used to implement the AvailableGameServer
type AvailableGameServer struct{}

func RunAvailableGameServer(s* grpc.Server) {
	GameServer.RegisterAvialableGameServerServer(s, &AvailableGameServer{})
}

func (s *AvailableGameServer) getAvailableGamesList(ctx context.Context, in GameServer.GameFilters) (GameServer.GameIdList, error) {
	var ids []*(GameServer.GameId) = make([]*(GameServer.GameId), 1)
	ids[0] = &GameServer.GameId{"id 1"}
	
	return GameServer.GameIdList{ids}, nil
}

func (s *AvailableGameServer) getGameDetails(ctx context.Context, in GameServer.GameId) (GameServer.GameDetails, error) {
	var details GameServer.GameDetails
	details.Console = "nes"
	
}