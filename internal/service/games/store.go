package games

import "gorm.io/gorm"

type GameStore interface {
}

type GameRepo struct {
	client *gorm.DB
}

func NewGameRepo(client *gorm.DB) *GameRepo {
	return &GameRepo{client: client}
}
