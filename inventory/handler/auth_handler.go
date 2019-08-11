package handler

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/miun173/inventory-dispusibda/inventory/app"
	"github.com/miun173/inventory-dispusibda/inventory/models"
	"github.com/pkg/errors"
)

// AuthHandler handle user auth
type AuthHandler interface {
	Authorize(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	userService app.UserService
}

// NewAuthHandler get authHandler object
func NewAuthHandler(userService app.UserService) AuthHandler {
	return &authHandler{
		userService,
	}
}

func decodeAuthToken(authToken string) (int, string, error) {
	s := strings.SplitN(authToken, " ", 2)
	if len(s) != 2 {
		return 0, "", errors.New("authToken not valid")
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return 0, "", errors.WithStack(err)
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return 0, "", errors.New("authToken not valid")
	}

	id, err := strconv.Atoi(pair[0])
	if err != nil {
		return 0, "", errors.WithStack(err)
	}

	return id, pair[1], nil
}

// Authorize validate user auth
func (h *authHandler) Authorize(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	authToken := r.Header.Get("Authorization")
	id, token, err := decodeAuthToken(authToken)
	if err != nil {
		log.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err = h.userService.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if token != user.Token || user.Role != user.Role {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

// Login handle user authentication
func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	var u models.User
	_ = json.NewDecoder(r.Body).Decode(&u)

	user, err := h.userService.UserOfName(u.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "system error"})
		return
	}

	if user.ID == 0 || user.Password != user.Password {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(map[string]string{"error": "incorect username or password"})
		return
	}

	user.Password = ""
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(user)
}
