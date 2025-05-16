import { defineConfig } from "vite";
import viteReact from "@vitejs/plugin-react";
import tailwindcss from "@tailwindcss/vite";

import { TanStackRouterVite } from "@tanstack/router-plugin/vite";
import { resolve } from "node:path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [TanStackRouterVite({ autoCodeSplitting: true }), viteReact(), tailwindcss()],
  test: {
    globals: true,
    environment: "jsdom",
  },
  preview: {
    port: 3000,
    strictPort: true,
  },
  server: {
    host: "0.0.0.0",
    port: 3000,
    strictPort: true,
    hmr: {
      host: "0.0.0.0",
      clientPort: 3000
    },
    watch: {
      usePolling: true,
    },
  },
  resolve: {
    alias: {
      '@': resolve(__dirname, './src'),
    },
  }
});
