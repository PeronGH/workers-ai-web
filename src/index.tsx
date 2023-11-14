import { Hono } from "hono";
import { Layout } from "./layout";
import { run } from "./utils";
import { ImageGeneration } from "./image-generation";

const app = new Hono();

app.get("/", (c) =>
  c.html(
    <Layout title="Workers AI Web UI">
      <ImageGeneration />
    </Layout>,
  ));

app.post("/api/run/*", async (c) => {
  const model = c.req.path.slice("/api/run/".length);
  return run(model, await c.req.arrayBuffer());
});

const server = Bun.serve({
  fetch: app.fetch,
  port: Bun.env.PORT ?? 3000,
});

console.info(`Listening on http://${server.hostname}:${server.port}`);
