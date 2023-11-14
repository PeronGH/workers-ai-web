export interface CloudflareAICredentials {
  accountId: string;
  apiToken: string;
}

export class CloudflareAI {
  readonly #accountId: string;
  readonly #apiToken: string;

  constructor({ accountId, apiToken }: CloudflareAICredentials) {
    this.#accountId = accountId;
    this.#apiToken = apiToken;
  }

  run(model: string, input: unknown) {
    const shouldKeep = typeof input === "string" ||
      input instanceof Uint8Array ||
      input instanceof ReadableStream ||
      input instanceof Blob ||
      input instanceof ArrayBuffer;

    const body = shouldKeep ? input : JSON.stringify(input);

    console.debug("RunRequest", {
      model,
      input: body,
    });

    return fetch(
      `https://api.cloudflare.com/client/v4/accounts/${this.#accountId}/ai/run/${model}`,
      {
        headers: { Authorization: `Bearer ${this.#apiToken}` },
        method: "POST",
        body,
      },
    );
  }
}
