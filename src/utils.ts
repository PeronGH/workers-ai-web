import { CloudflareAI } from "./ai";

const ai = new CloudflareAI({
  accountId: Bun.env.CF_ACCOUNT_ID!,
  apiToken: Bun.env.CF_API_TOKEN!,
});

export const run: typeof ai.run = ai.run.bind(ai);
