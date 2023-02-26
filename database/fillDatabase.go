package database

import (
	"encoding/csv"
	"fmt"
	"io"
	"juro-go/models"
	"log"
	"os"
	"strconv"
)

const (
	image = 16
	prep  = 11
	name  = 1
	desc  = 5
)

func FillDB() {

	file, err := os.Open("data/drinksAlk.csv")
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		ingredients := make([]models.Ingredients, 0)

		for j := 0; j < 15; j++ {

			amount, _ := strconv.Atoi(row[32+j])
			amount *= 30
			temp := models.Ingredient{}
			DB.Find(&temp, "name like ?", row[17+j])
			//fmt.Println(1, temp.ID)

			if row[17+j] == "" {
				break

			} else {
				fmt.Println(2, row[17+j])
				fmt.Println(3, temp.ID)

				ingredients = append(ingredients, models.Ingredients{

					Amount:       amount,
					IngredientID: temp.ID,
					Ingredient:   models.Ingredient{},
					RecipeID:     0,
					Recipe:       models.Recipe{},
				})
			}

		}

		test := models.Recipe{

			Name:        row[name],
			Description: row[desc],
			Preparation: row[prep],
			Image:       row[image],
			Ingredients: ingredients,
		}

		DB.Create(&test)

	}

}
