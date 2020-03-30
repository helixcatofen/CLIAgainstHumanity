package game

import(
	//"fmt"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

type CardDeck struct{
	BlackCards []BlackCard
	WhiteCards []string
}

type BlackCard struct{
	Text string
	Pick int
}

func New() CardDeck{
	var deck CardDeck
	dat, err := ioutil.ReadFile("/home/helixcatofen/go/src/github.com/CardsAgainstHumanity/server/game/cards.json")
	if err != nil {
        panic(err)
	}
	json.Unmarshal(dat, &deck)
	return deck
}

func GetBlackCard(deck CardDeck) BlackCard{
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(deck.BlackCards))
	card := deck.BlackCards[index]
	deck.BlackCards[index] = deck.BlackCards[len(deck.BlackCards) - 1]
	deck.BlackCards = deck.BlackCards[:len(deck.BlackCards) - 1]
	return card
} 

func GetWhiteCard(deck CardDeck) string{
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(deck.BlackCards))
	card := deck.WhiteCards[index]
	deck.WhiteCards[index] = deck.WhiteCards[len(deck.WhiteCards) - 1]
	deck.WhiteCards = deck.WhiteCards[:len(deck.WhiteCards) - 1]
	return card
}