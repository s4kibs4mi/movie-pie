package rest

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Status: http.StatusOK,
		Data:   "OK",
	}
	resp.ServerJSON(w)
}
