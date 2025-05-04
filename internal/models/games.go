package models

import "github.com/gorilla/websocket"

type Player struct {
	Username string
	Conn     *websocket.Conn
	Piece    string
}

type GameMove struct {
	Board3D [3][3]string `json:"board3d"`
	Board1D []string     `json:"board1d"`
	Index   int          `json:"index"`
	Status  string       `json:"status"`
	Room    string       `json:"room"`
	Turn    string       `json:"turn"`
	Type    string       `json:"type"`
}
