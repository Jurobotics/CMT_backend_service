import { Router } from "oak";
import { db } from "../utils/database.ts";
import { getRecipe } from "../utils/recipe.ts";

const router = new Router({
  prefix: "/recipe",
});

router.get("/", async (ctx) => {
  ctx.response.body = await db.queryEntries("SELECT * FROM recipe");
});

router.get("/:id", async (ctx) => {
  const { id } = ctx.params;

  const data = await getRecipe(id);

  ctx.response.body = data;
});

router.post("/", async (ctx) => {
  const body = await ctx.request.body({ type: "json" }).value;

  db.query(
    "INSERT INTO rezepte (name, beschreibung, zutaten, zubereitung) VALUES (?, ?, ?, ?)",
    [body.name, body.beschreibung, body.zutaten, body.zubereitung]
  );

  ctx.response.status = 201;
  ctx.response.body = {
    message: "Rezept wurde erstellt",
  };
});

export default router;
