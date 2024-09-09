package main

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/mtslzr/pokeapi-go"
	"github.com/xuri/excelize/v2"
)

type Movement struct {
	Name string 
	Power, Accuracy int
	MoveType string 
	Class string 
	PP int 
	Effect string 
	EffectChance interface{} 
	EnergyCost int
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

		EmptyMovement.Name = capitalizeFirstLetter(move.Name)
		EmptyMovement.Accuracy = move.Accuracy 
		EmptyMovement.Class =  strings.Title(move.DamageClass.Name)
		EmptyMovement.MoveType = strings.Title(move.Type.Name)
		EmptyMovement.PP = move.Pp
		EmptyMovement.Power = move.Power
		
		if move.EffectChance != nil {
			EmptyMovement.EffectChance = move.EffectChance
		} else {
			EmptyMovement.EffectChance = "-"
		}

		if move.EffectEntries != nil && len(move.EffectEntries) > 0 {
			EmptyMovement.Effect = move.EffectEntries[0].Effect
		}

		AllMovements = append(AllMovements, EmptyMovement)

	}

	//create logic to write AllMomevents to excel
	f := excelize.NewFile()

	var sheet = "sheet1"

	for move_index, move := range AllMovements {

		for i := 0; i < 9; i++{

			switch i {
					
			case 0: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForMoves[i], move_index+2) , move_index+1)
			
			case 1: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForMoves[i], move_index+2) , move.Name)

			case 2:  f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForMoves[i], move_index+2) , move.MoveType)
			
			case 3: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForMoves[i], move_index+2) , move.Class)
			
			case 4: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForMoves[i], move_index+2) ,CalculatePower(move.Power))

			case 5: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForMoves[i], move_index+2) , CalculateAccuracy(move.Accuracy))
			
			case 6: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForMoves[i], move_index+2) , CalculateEnergyCost(move.PP))
			
			case 7:  f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForMoves[i], move_index+2) , move.Effect)
			
			case 8: f.SetCellValue(sheet, fmt.Sprintf("%v%v", excelColumnsForMoves[i], move_index+2) , CreateEffectChance(move.EffectChance))
				
			}

		}
	
	}

	f.SaveAs("AllMovements.xlsx")


}

func capitalizeFirstLetter(s string) string {
    if len(s) == 0 {
        return s
    }
    // Dividir la cadena en partes usando el guion como delimitador
    parts := strings.Split(s, "-")
    // Convertir la primera letra de la primera parte a mayúscula
    parts[0] = strings.Title(parts[0])
    // Unir las partes de nuevo con guiones
    return strings.Join(parts, " ")
}

var CreateEffectChance = func(chanceEffect interface{})  string {

	if chanceEffect == "-" {
		return "-"
	} else {
		return fmt.Sprintf("%v%%", chanceEffect)
	}

}

var CalculatePower = func(power int) (finalPower int) {
	result := float64(power) / 10
	if result-math.Floor(result) == 0.5 {
		finalPower = int(math.Floor(result))
	} else {
		finalPower = int(math.Round(result))
	}
	return finalPower
}

var CalculateAccuracy = func(accuracy_as_Int int)(accuracy_as_d10 string){

	if accuracy_as_Int == 0 {
		return "-"
	}

	result := 10 - math.Round(float64(accuracy_as_Int)/10)

	if result == 0 {
		return "D10"
	} else {
		return fmt.Sprintf("D10 - %v", result)
	}

}

var CalculateEnergyCost = func(pp int)(energycost int){

	energycost = pp/5

	switch energycost {

	case 1: return 9
	case 2: return 8
	case 3: return 7
	case 4: return 6
	case 5: return 5
	case 6: return 4
	case 7: return 3
	case 8: return 2
	case 9: return 1
	case 10: return 1
	default: return 0
	}

}

var excelColumnsForMoves []string = []string{"A","B","C","D","E","F","G","H","I"}
var defaultColumnsNameForStats []string = []string{"ID","Name","Type","Class","Power","Accuracy","Effect","Effect chance","Energy cost"}