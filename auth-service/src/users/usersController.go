package users

import (
	"auth-service/src/messages"
	"database/sql"
	"log"
	"net/http"
)

type UserController struct {
	service UserService
}

func (u *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := u.service.GetAll()
	if err != nil {
		messages.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	messages.WriteJSON(w, http.StatusOK, users)
}

func (u *UserController) GetById(w http.ResponseWriter, r *http.Request) {
	userIdInt, err := GetUserIdFromURL(r.URL.Path)
	if err != nil {
		messages.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	user, err := u.service.GetByID(userIdInt)
	if err != nil {
		messages.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	messages.WriteJSON(w, http.StatusOK, user)

}

func (u *UserController) Add(w http.ResponseWriter, r *http.Request) {
	user := User{}
	err := messages.ReadJSON(w, r, &user)
	if err != nil {
		messages.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	userInserted, err := u.service.Insert(user)
	if err != nil {
		messages.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	messages.WriteJSON(w, http.StatusOK, userInserted)
}

func (u *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	userIdInt, err := GetUserIdFromURL(r.URL.Path)
	if err != nil {
		messages.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	err = u.service.DeleteByID(userIdInt)
	if err != nil {
		messages.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	messages.WriteJSON(w, http.StatusOK, nil)
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserController) Login(w http.ResponseWriter, r *http.Request) {
	payload := LoginPayload{}
	err := messages.ReadJSON(w, r, &payload)
	if err != nil {
		log.Printf("Error ReadJSON %s\n", err)
		messages.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	log.Printf("Checking email %s and password %s\n", payload.Email, payload.Password)
	user, err := u.service.GetByEmailAndCheckPassword(payload.Email, payload.Password)
	if err != nil {
		log.Printf("Error GetByEmailAndCheckPassword %s\n", err)
		messages.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	messages.WriteJSON(w, http.StatusOK, user)

}

func NewUserController(conn *sql.DB) *UserController {
	userDao := UserDAO{
		dbConn: conn,
	}
	userService := UserService{
		dao: userDao,
	}
	userController := UserController{
		service: userService,
	}
	return &userController
}
