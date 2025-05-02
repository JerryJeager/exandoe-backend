package manualwire

import (
	"github.com/JerryJeager/exandoe-backend/config"
	"github.com/JerryJeager/exandoe-backend/internal/http"
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
