package ServerManager

import (
	"google.golang.org/grpc"
	DownloadServer "UserDownloadServer"
	"golang.org/x/net/context"
)

type UserDownloadServer struct{}

func RunUserDownloadServer(s* grpc.Server) {
	DownloadServer.RegisterUserDownloadServerServer(s, &UserDownloadServer{})
}

func (s *UserDownloadServer) DownloadGame(ctx context.Context, in *DownloadServer.DownloadRequest) (*DownloadServer.DownloadResponse, error) {
	var response DownloadServer.DownloadResponse
	response.Available = true
	response.S3DownloadLink = "somewhere on S3"
	return &response, nil
}

func (s *UserDownloadServer) AvailableDownloads(ctx context.Context, in *DownloadServer.UserDownloadListRequest) (*DownloadServer.DownloadsList, error) {
	var downloads []*DownloadServer.Game = make([]*DownloadServer.Game, 1)
	downloads[0].Id = "some uuid"
	return &DownloadServer.DownloadsList{downloads}, nil
}