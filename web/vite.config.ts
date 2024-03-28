import { fileURLToPath, URL } from "url";
import { defineConfig } from "vite";
import solid from "vite-plugin-solid";

export default defineConfig({
  plugins: [solid()],
  resolve: {
    alias: [
      {
        find: "@",
        replacement: fileURLToPath(new URL("./src", import.meta.url)),
      },
    ],
  },
  build: {
    assetsDir: "static",
    emptyOutDir: true,
    outDir: "../server/public_dist",
  },
});