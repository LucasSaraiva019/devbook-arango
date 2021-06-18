package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/LucasSaraiva019/devbook-arango/src/models"
	"github.com/LucasSaraiva019/devbook-arango/src/repositories"
	"github.com/LucasSaraiva019/devbook-arango/src/responses"
)

// CreateUser insert a user into the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("register"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	err = repositories.Create(&user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}
