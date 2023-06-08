package database

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"juro-go/models"
	"log"
	"strconv"
)

const (
	image = 2
	alc   = 1
	name  = 0
	fIng  = 15
	aIng  = 12
)

func FillDB() {
	rows := readFile("data/Drinks (2).xlsx")
	i := 1

	for {
		fmt.Println(i)
		if i == len(rows) {
			break
		}
		ingredients := make([]models.Ingredients, 0)

		for j := 0; j < aIng; j++ {

			temp := models.Ingredient{}
			DB.Find(&temp, "name like ?", rows[i][fIng+j])

			def := models.Default{}

			var amount float32 = 0
			var amountSolid string

			if temp.Name == "" {
				break
			}
			if rows[i][27+j] == "1" {
				amountSolid = rows[i][3+j]
			} else {

				a, err := strconv.ParseFloat(rows[i][3+j], 32)
				if err != nil {
					log.Fatal(err)
				}
				amount = float32(a)
			}
			ingredients = append(ingredients, models.Ingredients{
				Default:      def,
				Amount:       amount,
				AmountSolid:  amountSolid,
				IngredientID: temp.ID,
				Ingredient:   models.Ingredient{},
				RecipeID:     def.ID,
				Recipe:       models.Recipe{},
			})

		}
		var alc1 bool
		alc1, _ = strconv.ParseBool(rows[i][alc])

		test := models.Recipe{

			Name:        rows[i][name],
			Image:       rows[i][image],
			Alc:         alc1,
			Ingredients: ingredients,
		}
		DB.Create(&test)

		i++
	}

}
func readFile(path string) [][]string {
	file, err := excelize.OpenFile(path)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := file.GetRows("Sheet1")
	if err != nil {
		log.Fatal(err)
	}
	return rows
}
