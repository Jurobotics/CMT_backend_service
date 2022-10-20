import { Router } from "oak";
import { getRecipe, RecipeData } from "../utils/recipe.ts";

const router = new Router({
  prefix: "/order",
});

interface pump {
  servo_id: number;
  pumps: number; // Dosierer Volume in ml
}

const url = "http://127.0.0.1:5015";

router.get("/:recipeId", async (ctx) => {
  const { recipeId } = ctx.params;

  if (!isNaN(+recipeId)) {
    const data: RecipeData = await getRecipe(+recipeId);

    const pumps: Array<pump> = [];

    data.ingredients.forEach((ingredient) => {
      pumps.push({
        servo_id: ingredient.arduino,
        pumps: ingredient.pumps / ingredient.amount,
      });
    });

    await fetch(url, {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify(pumps),
    });

    ctx.response.body = pumps;
  } else {
    ctx.response.body = { message: "please enter a valid number" };
    ctx.response.status = 400;
  }
});

export default router;
