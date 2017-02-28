package ServerManager

import (
	"google.golang.org/grpc"
	DownloadServer "UserDownloadServer"
)

//import (
//
//)
//

type UserDownloadServer struct{}

func RunUserDownloadServer(s* grpc.Server) {
	DownloadServer.RegisterUserDownloadServerServer(s, &UserDownloadServer{})
}
