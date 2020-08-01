package handlers

import (
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
			//fmt.Fprint(w, "You made a GET request 🖖")
			usersPostOne(w, r)
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

	id := bson.ObjectIdHex(path)

	switch r.Method {
	case http.MethodGet:
		userGetOne(w, r, id)
	case http.MethodPut:
		usersPutOne(w, r, id)
		return
	case http.MethodPatch:
		usersPatchOne(w, r, id)
		return
	case http.MethodDelete:
		usersDeleteOne(w, r, id)
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}

}
