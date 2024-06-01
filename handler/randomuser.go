package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dangquyitt/go-github-action/model"
	"github.com/dangquyitt/go-github-action/utils"
)

const randomUserAPIURL = "https://randomuser.me/api/"

func GetRandomUser(w http.ResponseWriter, r *http.Request) {
	userData, err := fetchRandomUser()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error fetching user data: %v", err)
		return
	}

	utils.EncodeJSON(w, userData.Results[0]) // Encode and return the first user data
}

func fetchRandomUser() (model.UserData, error) {
	resp, err := http.Get(randomUserAPIURL)
	if err != nil {
		return model.UserData{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.UserData{}, err
	}

	var userData model.UserData
	err = json.Unmarshal(body, &userData)
	if err != nil {
		return model.UserData{}, err
	}

	return userData, nil
}
