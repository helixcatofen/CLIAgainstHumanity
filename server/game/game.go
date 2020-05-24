package game

import(
	//"fmt"
)

type Game struct{
	Id int
	Players []Player
}

type Player struct{
	Id int
	Score int
}
 