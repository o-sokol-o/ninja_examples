package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

var users = []User{{1, "Vasya"}, {2, "Petya"}, {0, "Gera"}}

// curl -v -X GET -H 'x-id:1' localhost:8080/users
func main() {
	http.HandleFunc("/users", authMiddleware(loggerMiddleware(handleUsers)))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userlD := r.Header.Get("x-id")
		if userlD == "" {
			log.Printf("[%s] %s - error: userlD is not provided\n", r.Method, r.RequestURI)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "id", userlD)

		r = r.WithContext(ctx)

		next(w, r)
	}
}

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idFromCtx := r.Context().Value("id")
		userlD, ok := idFromCtx.(string)
		if !ok {
			log.Printf("[%s] %s — error: userlD is invalid", r.Method, r.URL)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Printf("[%sl %s by userlD %s\n", r.Method, r.URL, userlD)
		next(w, r)
	}
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		addUser(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user User
	if err = json.Unmarshal(reqBytes, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users = append(users, user)
}
