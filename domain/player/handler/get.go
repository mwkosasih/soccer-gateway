package handler

import (
	"context"
	"encoding/json"

	"github.com/labstack/echo/v4"
	soccer "github.com/mwkosasih/soccer-gateway/domain"
	"github.com/mwkosasih/soccer-gateway/domain/player/client"
	pb "github.com/mwkosasih/soccer-gateway/proto/soccer"
	"github.com/mwkosasih/soccer-gateway/util"
	"google.golang.org/grpc/status"
)

type Get struct{}

func (h *Get) Handle(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var req pb.GetPlayerRequest
	req.Id = int32(util.StringToInteger(c.Param("id")))
	grpcResp, err := client.Get(ctx, &req)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			resp := &util.Response{
				Code:    st.Code(),
				Message: util.StatusMessage[st.Code()],
			}
			return resp.JSON(c)
		}
		return err
	}
	resp, err := h.buildResponse(c, ctx, grpcResp)
	if err != nil {
		return err
	}
	return resp.JSON(c)
}

func (h *Get) buildResponse(c echo.Context, ctx context.Context, res *pb.GetPlayerResponse) (*util.Response, error) {
	var player soccer.Player
	bytes, _ := json.Marshal(res.Player)
	json.Unmarshal(bytes, &player)

	code := util.Success
	resp := &util.Response{
		Code:    code,
		Message: util.StatusMessage[code],
		Data: map[string]interface{}{
			"player": player,
		},
	}
	return resp, nil
}

func NewGet() *Get {
	return &Get{}
}
