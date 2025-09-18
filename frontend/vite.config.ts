import { defineConfig } from 'vite';
import viteReact from '@vitejs/plugin-react';
import tailwindcss from "@tailwindcss/vite";

import { TanStackRouterVite } from '@tanstack/router-plugin/vite'
import { resolve } from 'node:path'

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [
		TanStackRouterVite({ autoCodeSplitting: true }),
		viteReact(),
		tailwindcss(),
	],
	base: import.meta.env.PROD ? "https://knnkt.dk" : "https://localhost",
	server: {
		host: "0.0.0.0",
		port: 3000,
		strictPort: true,
		hmr: {
			clientPort: 443,
		},
		watch: {
			usePolling: true
		}
	},
	test: {
		globals: true,
		environment: 'jsdom',
	},
	resolve: {
		alias: {
			'@': resolve(__dirname, './src'),
		},
	},
})
