import { Router } from "oak";
import db from "../utils/database.ts";

interface Recipe {
    id?: number;
    name: string;
    beschreibung: string;
    zutaten: string;
    zubereitung: string;
}
  
const router = new Router({
    prefix: "/recipe",
})

router.get('/', async (ctx) => {
    ctx.response.body = await db.query('SELECT * FROM rezepte');
});
  
router.get('/:id', async (ctx) => {
    const { id } = ctx.params;
    const rezept = await db.query('SELECT * FROM rezepte WHERE id=?', [id]);
  
    ctx.response.body = rezept;
});
  
  
router.post('/', async (ctx) => {
    const body: Recipe = await ctx.request.body({ type: 'json' }).value;
  
    db.query('INSERT INTO rezepte (name, beschreibung, zutaten, zubereitung) VALUES (?, ?, ?, ?)', [body.name, body.beschreibung, body.zutaten, body.zubereitung]);
    
    ctx.response.status = 201;
    ctx.response.body = {
      message: 'Rezept wurde erstellt'
    }

});
  

export default router;