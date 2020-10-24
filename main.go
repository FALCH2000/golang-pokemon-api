package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"pokemon-api/database"

	"github.com/gorilla/mux"
)

func getAllPokemons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.PokemonDb)
}

func insertPokemon(w http.ResponseWriter, r *http.Request) {
	var newPokemon database.Pokemon

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newPokemon)

	database.PokemonDb[newPokemon.ID] = newPokemon
}

func handleRequests() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "88"
	}
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(commonMiddleware)
	myRouter.HandleFunc("/pokemons", getAllPokemons).Methods("GET")
	myRouter.HandleFunc("/pokemons", insertPokemon).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("Pokemon Rest API")
	handleRequests()
}
