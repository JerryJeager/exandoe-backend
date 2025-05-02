package manualwire

import (
	"github.com/JerryJeager/exandoe-backend/config"
	"github.com/JerryJeager/exandoe-backend/internal/http"
	"github.com/JerryJeager/exandoe-backend/internal/service/games"
	"github.com/JerryJeager/exandoe-backend/internal/service/users"
)

func GetUserRepository() *users.UserRepo {
	repo := config.GetSession()
	return users.NewUserRepo(repo)
}

func GetUserService(repo users.UserStore) *users.UserServ {
	return users.NewUserService(repo)
}

func GetUserController() *http.UserController {
	repo := GetUserRepository()
	service := GetUserService(repo)
	return http.NewUserController(service)
}

func GetGameRepository() *games.GameRepo {
	repo := config.GetSession()
	return games.NewGameRepo(repo)
}

func GetGameService(repo games.GameStore) *games.GameServ {
	return games.NewGameService(repo)
}

func GetGameController() *http.GameController {
	repo := GetGameRepository()
	service := GetGameService(repo)
	return http.NewGameController(service)
}
