package order

import (
	"juro-go/api/recipe"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Pump struct {
	ServoId int `json:"servo_id"`
	Pumps   int `json:"pumps"`
}

func Order(ctx *fiber.Ctx) error {
	recipeId := ctx.Params("recipeId")
	id, _ := strconv.Atoi(recipeId)

	recipe := recipe.GetRecipeById(id)

	var pumps []Pump

	for _, element := range recipe.Ingredients {
		pumps = append(pumps, Pump{ServoId: int(element.ID), Pumps: (element.AmountNeeded / int(element.DoserSize))})
	}

	return ctx.JSON(pumps)
}
