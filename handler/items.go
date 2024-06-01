package handler

import (
	"net/http"

	"github.com/dangquyitt/go-github-action/model"
	"github.com/dangquyitt/go-github-action/utils"
)

func ItemsHandler(w http.ResponseWriter, r *http.Request) {
	items := []model.Item{
		{ID: 1, Name: "Book", Price: 10.99},
		{ID: 2, Name: "Pen", Price: 2.50},
		{ID: 3, Name: "Headphones", Price: 49.99},
	}
	utils.EncodeJSON(w, items)
}
