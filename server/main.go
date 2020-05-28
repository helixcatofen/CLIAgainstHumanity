package main

import (
	"fmt"
	"log"
	"encoding/json"
	"strconv"
	"net/http"
	"github.com/google/uuid"
	"io/ioutil"
	"github.com/gorilla/mux"
	"github.com/CardsAgainstHumanity/server/game"
)

var allGames map[int] *game.Game

type NewPlayer struct{
	Name string
}

type Error struct{
	Error string 
}
 

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
	allGames[int(id.ID())] = &newGame
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newGame)
}

func joinGame(w http.ResponseWriter, r *http.Request){
	var newPlayer NewPlayer
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Fprint(w, "unknown error")
	}
	temp,_ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	gameId := int(temp)
	json.Unmarshal(body, &newPlayer)
	if((allGames[gameId]) != nil){
		w.WriteHeader(http.StatusOK)
		id,_ := uuid.NewUUID()
		var player game.Player = game.Player{int(id.ID()), 0, newPlayer.Name}
		allGames[gameId].Players = append(allGames[gameId].Players, player)
		json.NewEncoder(w).Encode(allGames[gameId])
	}	else{
		w.WriteHeader(http.StatusNotFound)
		errMsg := Error{"This game does not exist"}
		json.NewEncoder(w).Encode(errMsg)
	}
}

// func playCard(w http.ResponseWriter, r *http.Request){

// }

// func voteCard(w http.ResponseWriter, r *http.Request){

// }

// func getPlayedCards(w http.ResponseWriter, r *http.Request){

// }

func main() {
	allGames = make(map[int] *game.Game)
	fmt.Println("Started server!")
	router := mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/create_game", createGame).Methods("POST")
	router.HandleFunc("/join_game/{id}", joinGame).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
	// deck := game.New()
}

// func gameIndex(gameId int) int{
// 	for i:=0; i <len(allGames); i++{
// 		if(allGames[i].Id == gameId) {return i}
// 	}
// 	return -1
// }