import { db } from "./database.ts";

export interface Recipe {
  id: number;
  name: string;
  beschreibung: string;
  zubereitung: string;
  bild: string;
}

export interface Ingredient {
  name: string;
  amount: number;
  uses: number;
  pumps: number;
  arduino: number;
}

export interface RecipeData {
  recipe: Array<Recipe>;
  ingredients: Array<Ingredient>;
}

export const getRecipe = async (id: number): Promise<RecipeData> => {
  const ingredients: Array<Ingredient> = await db.queryEntries(
    `SELECT ingredient.name, ingredient.amount, ingredient.uses, i.amount as pumps, s.arduino
FROM ingredient
         LEFT JOIN ingredients i on ingredient.id = i.ingredient_id
         LEFT JOIN recipe r on i.recipe_id = r.id
        LEFT OUTER JOIN servos s on s.id = ingredient.servo_id
WHERE r.id =?`,
    [id]
  );

  const recipe: Array<Recipe> = await db.queryEntries(
    `SELECT * FROM recipe WHERE id=?`,
    [id]
  );

  const data: RecipeData = { recipe, ingredients };

  return data;
};
