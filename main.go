package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Tech Struct
type Tech struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Deleted bool   `json:"deleted"`
}

var techs []Tech

func getTechs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if !(len(techs) > 0) {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	json.NewEncoder(w).Encode(techs)
}
func getTech(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range techs {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(404)
}
func createTech(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tech Tech
	_ = json.NewDecoder(r.Body).Decode(&tech)
	tech.ID = strconv.Itoa(len(techs) + 1)
	techs = append(techs, tech)
	json.NewEncoder(w).Encode(tech)
}
func updateTech(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range techs {
		if item.ID == params["id"] {
			var tech Tech
			_ = json.NewDecoder(r.Body).Decode(&tech)
			tech.ID = item.ID
			techs[index] = tech
			json.NewEncoder(w).Encode(tech)
			return
		}
	}
}
func deleteTech(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range techs {
		if item.ID == params["id"] {
			var tech Tech
			_ = techs[index]
			tech.Deleted = true
			techs[index] = tech
			json.NewEncoder(w).Encode(tech)
			return
		}

	}
	json.NewEncoder(w).Encode(techs)
}
func main() {
	//Init Router
	r := mux.NewRouter()

	techs = append(techs, Tech{ID: "1", Name: "Go"})
	techs = append(techs, Tech{ID: "2", Name: "C#"})
	techs = append(techs, Tech{ID: "3", Name: "SQL"})

	//Route handlers

	r.HandleFunc("/api/techs", getTechs).Methods("GET")
	r.HandleFunc("/api/techs/{id}", getTech).Methods("GET")
	r.HandleFunc("/api/techs", createTech).Methods("POST")
	r.HandleFunc("/api/techs/{id}", updateTech).Methods("PUT")
	r.HandleFunc("/api/techs/{id}", deleteTech).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
