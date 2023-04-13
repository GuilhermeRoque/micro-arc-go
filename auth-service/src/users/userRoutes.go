package users

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
)

const UsersEndpoint = "/users"
const LoginEndpoint = "/login"
const LogoutEndpoint = "/logout"
const UserIdEndpoint = "userId"

var UserEndpoint = fmt.Sprintf("%s/:%s", UsersEndpoint, UserIdEndpoint)

func GetUserIdFromURL(url string) (int, error) {
	userIdString := strings.TrimPrefix(url, fmt.Sprintf("%s/", UsersEndpoint))
	userIdInt, err := strconv.Atoi(userIdString)
	return userIdInt, err
}

func SetUsersRoutes(mux *chi.Mux, userController *UserController) {
	mux.Get(UsersEndpoint, userController.GetAll)
	mux.Post(UsersEndpoint, userController.Add)
	mux.Delete(UserEndpoint, userController.Delete)
	mux.Get(UserEndpoint, userController.GetById)
	mux.Post(LoginEndpoint, userController.Login)
}
