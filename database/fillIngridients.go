package database

import (
	"encoding/csv"
	"io"
	"juro-go/models"
	"log"
	"os"
)

func FillIng() { // idk why this works just dont doubt it

	file, err := os.Open("data/drinksNoneAlk.csv")
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
	}
	//temp := make([]string, 1)

	for {

		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		temp := models.Ingredient{}

		for k := 18; k < 32; k++ {

			DB.Find(&temp, "name like ?", row[k])

			if temp.Name == row[k] || temp.Name != "" {
				continue
			} else {
				ingredient := models.Ingredient{
					Name: row[k],
				}
				DB.Create(&ingredient)
			}

		}

	}
}

/*func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}*/
