package main

import (
	"log"

	"github.com/mtslzr/pokeapi-go"
)

type Movement struct {
	Name string 
	Power, Accuracy int
	MoveType string 
	Class string 
	PP int 
	Effect string 
	EffectChance interface{} 
}

var EmptyMovement Movement
var AllMovements []Movement

func main() {

	r, err :=	pokeapi.Resource("move", 0, 1000000)

	if err != nil {
		log.Println(err)
	}

	for _, r := range r.Results {

		move , err := pokeapi.Move(r.Name)

		if err != nil {
			log.Println(err)
		}

		EmptyMovement.Name = move.Name
		EmptyMovement.Accuracy = move.Accuracy
		EmptyMovement.Class = move.DamageClass.Name
		EmptyMovement.MoveType = move.Type.Name
		EmptyMovement.PP = move.Pp
		EmptyMovement.Power = move.Power
		EmptyMovement.Effect = move.EffectEntries[0].Effect
		EmptyMovement.EffectChance = move.EffectChance

		AllMovements = append(AllMovements, EmptyMovement)

	}

	//create logic to write AllMomevents to excel


}

var excelColumnsForMoves []string = []string{"A","B","C","D","E","F","G","H","I"}
var defaultColumnsNameForStats []string = []string{"ID","Name","Type","Class","Power","Accuracy","Effect","Effect chance","Energy cost"}