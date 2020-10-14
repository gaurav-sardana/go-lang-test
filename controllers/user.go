package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/halodoc/webservice/models"
)

type UserController struct {
	userIDPattern *regexp.Regexp
}

func (uc UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello from the User Controller!"))
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		userID, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		switch r.Method {
		case http.MethodGet:
			uc.get(userID, w)
		case http.MethodPut:
			uc.put(userID, w, r)
		case http.MethodDelete:
			uc.delete(userID, w)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func (uc *UserController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetUsers(), w)
}

func (uc *UserController) get(userId int, w http.ResponseWriter) {
	user, err := models.GetUserById(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(user, w)
}

func (uc *UserController) post(w http.ResponseWriter, r *http.Request) {
	user, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User Object"))
		return
	}
	user, err = models.AddUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(user, w)
}
func (uc *UserController) put(userId int, w http.ResponseWriter, r *http.Request) {
	user, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User Object"))
		return
	}
	if userId != user.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Id of submitted user must match with ID in url "))
		return
	}
	user, err = models.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(user, w)
}
func (uc *UserController) delete(userId int, w http.ResponseWriter) {
	err := models.RemoveUserById(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
}

func (uc *UserController) parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil

}

func newUserController() *UserController {
	return &UserController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}

}
