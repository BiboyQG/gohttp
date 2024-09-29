package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

var users = []User{
	{
		FirstName: "John",
		LastName:  "Doe",
	},
	{
		FirstName: "Jane",
		LastName:  "Doe",
	},
}

func (api *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (api *api) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u := User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}

	err = insertUser(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func insertUser(user User) error {
	// input validation
	if user.FirstName == "" || user.LastName == "" {
		return errors.New("first name and last name are required to be non-empty!")
	}
	
	//storage validation
	for _, u := range users {
		if u.FirstName == user.FirstName && u.LastName == user.LastName {
			return errors.New("user already exists")
		}
	}
	
	users = append(users, user)
	return nil
}
