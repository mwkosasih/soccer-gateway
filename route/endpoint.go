package route

import (
	"github.com/labstack/echo/v4"
	player "github.com/mwkosasih/soccer-gateway/domain/player/handler"
	team "github.com/mwkosasih/soccer-gateway/domain/team/handler"
)

// Handler endpoint to use it later
type Handler interface {
	Handle(c echo.Context) (err error)
}

var endpoint = map[string]Handler{
	// team
	"create_team": team.NewCreate(),
	"get_team":    team.NewGet(),

	// player
	"create_player": player.NewCreate(),
	"get_player":    player.NewGet(),
}
