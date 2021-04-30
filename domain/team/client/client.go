package client

import (
	"context"
	"os"

	pb "github.com/mwkosasih/soccer-gateway/proto/soccer"
	"github.com/mwkosasih/soccer-gateway/util"
)

func Get(ctx context.Context, req *pb.GetTeamRequest) (*pb.GetTeamResponse, error) {
	conn := util.Dial(os.Getenv("soccer_grpc"))
	defer conn.Close()
	client := pb.NewSoccerServiceClient(conn)
	return client.GetTeam(ctx, req)
}

func Create(ctx context.Context, req *pb.CreateTeamRequest) (*pb.NoResponse, error) {
	conn := util.Dial(os.Getenv("soccer_grpc"))
	defer conn.Close()
	client := pb.NewSoccerServiceClient(conn)
	return client.CreateTeam(ctx, req)
}
