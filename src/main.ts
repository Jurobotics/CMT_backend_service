import { Application } from "oak";

import db from "./utils/database.ts";

import router from './routes/rezept.ts';
import routerQueue from './routes/queue.ts';
import routerModul from './routes/modul.ts';

await db.init("test.sql");

const app = new Application();

// db.init('test.sql');

app.use(router.routes());
app.use(router.allowedMethods());

app.use(routerModul.routes());
app.use(routerModul.allowedMethods());

app.use(routerQueue.routes());
app.use(routerQueue.allowedMethods());


await app.listen({ port: 8000 });