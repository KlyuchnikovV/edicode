import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import preprocess from 'svelte-preprocess'
import typescript from '@rollup/plugin-typescript';

// https://vitejs.dev/config/
/** @type {import('vite').UserConfig} */
export default defineConfig({
  plugins: [
    svelte({
      include: [
        "src/**/*.svelte",
        "node_modules/**/*.svelte",
        "../plugins/**/*.svelte",
      ],
      compilerOptions: {
      },
      preprocess: preprocess({}),
    }),
    typescript({
      sourceMap: true
    }),

  ],
  build: {
    outDir: "./dist",
    assetsDir: ".",
    sourcemap: true,
    minify: false,
  },

})
