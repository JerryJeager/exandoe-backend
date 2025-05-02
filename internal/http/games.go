package http

import (
	"net/http"

	"github.com/JerryJeager/exandoe-backend/config"
	"github.com/JerryJeager/exandoe-backend/internal/models"
	"github.com/JerryJeager/exandoe-backend/internal/service/games"
	"github.com/gin-gonic/gin"
)

type GameController struct {
	serv games.GameSv
}

func NewGameController(serv games.GameSv) *GameController {
	return &GameController{serv: serv}
}

func (c *GameController) Play(ctx *gin.Context) {
	conn, err := config.WS.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	room := ctx.Query("room")
	username := ctx.Query("username")

	players, ok := config.ActiveGames[room]
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "room not found"})
		return
	}

	updated := false
	for i, c := range players {
		if c.Username == username {
			players[i].Conn = conn
			updated = true
			break
		}
	}
	if !updated {
		players = append(players, &models.Player{Username: username, Conn: conn})
	}
	config.ActiveGames[room] = players

	for {
		var move models.GameMove
		err := conn.ReadJSON(&move)
		if err != nil {
			break
		}

		// Broadcast to opponent
		for _, p := range players {
			if p.Username != username {
				_ = p.Conn.WriteJSON(move)
			}
		}
	}
}
