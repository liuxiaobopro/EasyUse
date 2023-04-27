import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import WindiCSS from 'vite-plugin-windicss'

import path from "path";

export default defineConfig({
    resolve: {
        alias: {
            "@": path.resolve(__dirname, "src"),
        }
    },
    server: {
        host: '0.0.0.0',
        port: 50002,
        proxy: {
            '/api': {
                target: 'url',
                changeOrigin: true,
                rewrite: (path) => path.replace(/^\/api/, '')
            },
        }
    },
    plugins: [vue(), WindiCSS()],
})
