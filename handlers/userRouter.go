package handlers

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strings"
)

// UserRouter handles the users route
func UserRouter(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")

	if path == "/users" {
		switch r.Method {
		case http.MethodGet:
			//fmt.Fprint(w, "You made a GET request 🚀")
			userGetAll(w, r)
			return
		case http.MethodPost:
			fmt.Fprint(w, "You made a GET request 🖖")
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
		}
	}

	path = strings.TrimPrefix(path, "/users/")
	if !bson.IsObjectIdHex(path) {
		postError(w, http.StatusNotFound)
		return
	}

	//id := bson.ObjectIdHex(path)

	switch r.Method {
	case http.MethodGet:
		return
	case http.MethodPut:
		return
	case http.MethodPatch:
		return
	case http.MethodDelete:
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}

}
