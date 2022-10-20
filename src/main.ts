import { Application } from "oak";
import { oakCors } from "https://deno.land/x/cors/mod.ts";

import { db } from "./utils/database.ts";

import * as path from "path";

const PORT = 8000;

await db.init("app.sql");

const app = new Application();
app.use(oakCors());
// db.init('test.sql');

const routesPath = path.join(Deno.cwd(), "src", "routes");

for await (const file of Deno.readDir(routesPath)) {
  const { default: router } = await import(path.join(routesPath, file.name));
  app.use(router.routes());
  app.use(router.allowedMethods());
}

Deno.addSignalListener("SIGTERM", () => {
  db.close();
  Deno.exit(0);
});

console.log(`Server is running on http://localhost:${PORT}`);

await app.listen({ port: PORT });
