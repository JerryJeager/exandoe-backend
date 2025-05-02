package games

type GameSv interface {
}

type GameServ struct {
	repo GameStore
}

func NewGameService(repo GameStore) *GameServ {
	return &GameServ{repo: repo}
}
