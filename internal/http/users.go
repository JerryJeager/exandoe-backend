package http

import (
	"net/http"

	"github.com/JerryJeager/exandoe-backend/config"
	"github.com/JerryJeager/exandoe-backend/internal/models"
	"github.com/JerryJeager/exandoe-backend/internal/service/users"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	serv users.UserSv
}

func NewUserController(serv users.UserSv) *UserController {
	return &UserController{serv: serv}
}

func (c *UserController) Signin(ctx *gin.Context) {
	userNameParam := ctx.Query("username")
	if userNameParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	if _, ok := config.WS.Clients[userNameParam]; ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user already connected"})
		return
	}

	conn, err := config.WS.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	config.WS.Clients[userNameParam] = conn
	defer delete(config.WS.Clients, userNameParam)
	defer func() {
		delete(config.WS.Clients, userNameParam)

		onlineUsers := make([]string, 0, len(config.WS.Clients))
		for username := range config.WS.Clients {
			onlineUsers = append(onlineUsers, username)
		}

		for _, client := range config.WS.Clients {
			_ = client.WriteJSON(gin.H{
				"type":  "online_users",
				"users": onlineUsers,
			})
		}
	}()

	onlineUsers := make([]string, 0, len(config.WS.Clients))
	for username := range config.WS.Clients {
		onlineUsers = append(onlineUsers, username)
	}

	for _, client := range config.WS.Clients {
		err := client.WriteJSON(gin.H{
			"type":  "online_users",
			"users": onlineUsers,
		})
		if err != nil {
			continue
		}
	}

	// Keep connection alive
	for {
		var msg models.ChallengeMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			break // client disconnected or sent invalid data
		}

		switch msg.Type {
		case "challenge_request":
			if receiverConn, ok := config.WS.Clients[msg.To]; ok {
				receiverConn.WriteJSON(msg)
			}
		case "challenge_response":
			if receiverConn, ok := config.WS.Clients[msg.From]; ok {
				receiverConn.WriteJSON(msg)
			}
			if msg.Accepted != nil && *msg.Accepted {
				roomID := "game-" + msg.From + "-" + msg.To
				config.ActiveGames[roomID] = []*models.Player{{Username: msg.From, Conn: config.WS.Clients[msg.From]}, {Username: msg.To, Conn: config.WS.Clients[msg.To]}}

				startMsg := map[string]interface{}{
					"type":    "start_game",
					"room_id": roomID,
				}
				config.WS.Clients[msg.From].WriteJSON(startMsg)
				config.WS.Clients[msg.To].WriteJSON(startMsg)
			}
		default:
		}
	}
}
