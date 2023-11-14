import type { FC } from "hono/jsx";

/** Layout for a HTML page, importing `Alpine.js` and `Twind` */
export const Layout: FC<{ title: string }> = ({ title, children }) => (
  <html>
    <head>
      <meta charset="UTF-8" />
      <meta http-equiv="X-UA-Compatible" content="IE=edge" />
      <meta name="viewport" content="width=device-width, initial-scale=1.0" />
      <title>{title}</title>
      <script src="https://unpkg.com/alpinejs" defer />
      <script src="https://cdn.twind.style" />
    </head>
    <body>{children}</body>
  </html>
);
