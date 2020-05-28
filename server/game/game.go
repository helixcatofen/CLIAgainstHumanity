package game

import(
	//"fmt"
)

type Game struct{
	Id int
	Players []Player
	Judge int
}

type Player struct{
	Id int
	Score int
	Name string
}
 