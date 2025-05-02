package models

import "github.com/gorilla/websocket"

type Player struct {
	Username string
	Conn     *websocket.Conn
}


type GameMove struct {
	
}