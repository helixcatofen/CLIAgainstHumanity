package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"github.com/google/uuid"
	"io/ioutil"
	"github.com/gorilla/mux"
	"github.com/CardsAgainstHumanity/server/game"
)

var allGames []game.Game

func commonMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createGame(w http.ResponseWriter, r *http.Request){
	_, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Fprint(w, "unknown error")
	}
	id,_ := uuid.NewUUID()
	var newGame game.Game
	newGame.Id = int(id.ID())
	allGames = append(allGames, newGame)
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newGame)
}

// func joinGame(w http.ResponseWriter, r *http.Request){

// }

// func playCard(w http.ResponseWriter, r *http.Request){

// }

// func voteCard(w http.ResponseWriter, r *http.Request){

// }

// func getPlayedCards(w http.ResponseWriter, r *http.Request){

// }

func main() {
	fmt.Println("Started server!")
	router := mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/create_game", createGame).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
	// deck := game.New()
}