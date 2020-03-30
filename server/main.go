package main

import (
	"fmt"
    // "log"
	// "net/http"
	// "github.com/google/uuid"
	// "github.com/gorilla/mux"
	"github.com/CardsAgainstHumanity/server/game"
)

// type server struct{}

// func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")
//         w.WriteHeader(http.StatusOK)
//     w.Write([]byte(`{"message": "hello world"}`))
// }

// func createGame(w http.ResponseWriter, r *http.Request){

// }

// func joinGame(w http.ResponseWriter, r *http.Request){

// }

// func playCard(w http.ResponseWriter, r *http.Request){

// }

// func voteCard(w http.ResponseWriter, r *http.Request){

// }

// func getPlayedCards(w http.ResponseWriter, r *http.Request){

// }

func main() {
    // s := &server{}
    // http.Handle("/", s)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	deck := game.New()
	fmt.Println(deck)

}