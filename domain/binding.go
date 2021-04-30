package soccer

type (
	CreateTeam struct {
		Name string `json:"name"`
	}

	Team struct {
		ID      int32    `json:"id"`
		Name    string   `json:"name"`
		Players []Player `json:"players"`
	}

	CreatePlayer struct {
		TeamID   int32  `json:"team_id"`
		Name     string `json:"name"`
		Position string `json:"position"`
		Weight   int32  `json:"weight"`
		Height   int32  `json:"height"`
	}
	Player struct {
		ID       int32  `json:"id"`
		Name     string `json:"name"`
		Position string `json:"position"`
		Weight   int32  `json:"weight"`
		Height   int32  `json:"height"`
	}
)
