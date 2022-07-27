import { Router } from "oak";
import db from "../utils/database.ts";

interface Modul {
    id?: number;
    name: string;
    dosierer: string;
} 

const routerModul = new Router({
    prefix: "/modul",
});

routerModul.get('/', async (ctx) => {
    ctx.response.body = await db.query('SELECT * FROM module');
});

routerModul.post('/', async (ctx) => {
    const body: Modul = await ctx.request.body({ type: 'json' }).value;

    db.query('INSERT INTO module (name, dosierer) VALUES (?, ?)', [body.name, body.dosierer]);

    ctx.response.status = 201;
    ctx.response.body = {
        message: 'Modul wurde erstellt'
    }
})

export default routerModul;