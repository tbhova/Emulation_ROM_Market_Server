package ServerManager

import (
	GameServer "AvailableGamesServer"
	"golang.org/x/net/context"
	"log"
)

// Data type used to implement the AvailableGameServer
type AvailableGameServer struct{}

func RunAvailableGameServer() {
	GameServer.RegisterAvialableGameServerServer(s, &AvailableGameServer{})
}

func (s *AvailableGameServer) GetAvailableGamesList(ctx context.Context, in *GameServer.GameFilters) (*GameServer.GameIdList, error) {
	var ids []*(GameServer.GameId) = make([]*(GameServer.GameId), 1)
	ids[0] = &GameServer.GameId{"id 1"}
	
	return &GameServer.GameIdList{ids}, nil
}

func (s *AvailableGameServer) GetGameDetails(ctx context.Context, id *GameServer.GameId) (*GameServer.GameDetails, error) {
	var details *GameServer.GameDetails
	
	err := db.QueryRow("SELECT * FROM GAME WHERE ID = ?", id.GetId()).Scan(details)
	
	if err != nil {
		log.Fatalf("Get game id: %s Error: %s", id.GetId(), err)
		return details, err
	}
	
	return details, nil
	
	//details.Console = "nes"
	//details.Id = "0"
	//details.Title = "The Italian Plumber"
	//details.ReleaseDate = "01/12/1990"
	//details.Developer = "Horgan"
	//details.Publisher = "Horgan"
	//details.Genre = "Platformer"
	//details.Players = 2
	//return details, nil
}