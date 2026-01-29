import { defineConfig } from 'vite'
import electron from 'vite-plugin-electron'
import { resolve } from 'path'

export default defineConfig({
    plugins: [
        electron({
            entry: 'src/main/main.js',
            onstart(options) {
                options.startup()
            },
        }),
    ],
    root: resolve(__dirname, 'src/renderer'),
    build: {
        outDir: resolve(__dirname, 'dist'),
        emptyOutDir: true,
    },
    css: {
        postcss: './postcss.config.js'
    }
})