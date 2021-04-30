package client

import (
	"context"
	"os"

	pb "github.com/mwkosasih/soccer-gateway/proto/soccer"
	"github.com/mwkosasih/soccer-gateway/util"
)

func Get(ctx context.Context, req *pb.GetPlayerRequest) (*pb.GetPlayerResponse, error) {
	conn := util.Dial(os.Getenv("soccer_grpc"))
	defer conn.Close()
	client := pb.NewSoccerServiceClient(conn)
	return client.GetPlayer(ctx, req)
}

func Create(ctx context.Context, req *pb.CreatePlayerRequest) (*pb.NoResponse, error) {
	conn := util.Dial(os.Getenv("soccer_grpc"))
	defer conn.Close()
	client := pb.NewSoccerServiceClient(conn)
	return client.CreatePlayer(ctx, req)
}
