package handler

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
	soccer "github.com/mwkosasih/soccer-gateway/domain"
	"github.com/mwkosasih/soccer-gateway/domain/team/client"
	pb "github.com/mwkosasih/soccer-gateway/proto/soccer"
	"github.com/mwkosasih/soccer-gateway/util"
	"google.golang.org/grpc/status"
)

type Create struct{}

func (h *Create) Handle(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r := new(soccer.CreateTeam)
	err := h.validate(r, c)
	if err != nil {
		return status.Errorf(util.InvalidArgument, "Invalid Request")
	}

	var req pb.CreateTeamRequest
	req.Name = r.Name
	grpcResp, err := client.Create(ctx, &req)
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

func (h *Create) buildResponse(c echo.Context, ctx context.Context, res *pb.NoResponse) (*util.Response, error) {
	code := util.SuccessCreated
	resp := &util.Response{
		Code:    code,
		Message: util.StatusMessage[code],
	}
	return resp, nil
}

func (h *Create) validate(r *soccer.CreateTeam, c echo.Context) error {
	if err := c.Bind(r); err != nil {
		log.Println("validate", err)
		return err
	}
	return c.Validate(r)
}

func NewCreate() *Create {
	return &Create{}
}
