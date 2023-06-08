package database

import (
	"github.com/xuri/excelize/v2"
	"juro-go/models"
	"log"
	"strconv"
)

func FillIng() { // I don't know why this works just don't doubt it

	file, err := excelize.OpenFile("data/Drinks (2).xlsx")
	if err != nil {
		log.Fatal(err)
	}
	i := 1
	rows, err := file.GetRows("Sheet1")
	if err != nil {
		log.Fatal(err)
	}

	for {
		if i == len(rows) {
			break
		}

		temp := models.Ingredient{}

		for k := fIng; k < 27; k++ {

			j := k + 12

			DB.Find(&temp, "name like ?", rows[i][k])

			if temp.Name == rows[i][k] || temp.Name != "" {
				continue
			} else {

				ingredient := models.Ingredient{
					Name:  rows[i][k],
					Solid: isSolid(rows[i][j]),
				}
				DB.Create(&ingredient)
			}

		}
		i++
	}
}
func isSolid(string2 string) bool {
	if alc1, err := strconv.ParseBool(string2); err == nil {
		return alc1
	}
	return false
}

/*func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}*/
