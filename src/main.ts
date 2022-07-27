import { Application } from "oak";

import db from "./utils/database.ts";

import router from './routes/rezept.ts';

import routerQueue from './routes/queue.ts';

await db.init("test.sql");

const app = new Application();

// db.init('test.sql');

app.use(router.routes());
app.use(router.allowedMethods());



app.use(routerQueue.routes());
app.use(routerQueue.allowedMethods());


await app.listen({ port: 8000 });