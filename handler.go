package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hola")
}
func (s *Server) addMiddleware(f http.HandlerFunc, middleware ...Middleware) http.HandlerFunc {
	for _, m := range middleware {
		f = m(f)
	}
	return f

}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the API Endpoint")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	var metadata MetaData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "Error: %v\n", err)
		return
	}
	fmt.Fprintf(w, "Payload %v\n", metadata)
}

func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	var user User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error: %v\n", err)
		return
	}

	response, err := user.toJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// fmt.Fprintf(w, "Payload %v\n", user)
	fmt.Println(user.Name)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
