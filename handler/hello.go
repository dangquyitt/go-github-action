package handler

import (
	"net/http"

	"github.com/dangquyitt/go-github-action/utils"

	"github.com/dangquyitt/go-github-action/model"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	message := model.Message{Text: "Hello World!"}
	utils.EncodeJSON(w, message)
}
