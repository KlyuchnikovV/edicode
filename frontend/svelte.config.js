import preprocess from "svelte-preprocess";
import mm from 'micromatch';
import adapter from '@sveltejs/adapter-auto';

/** @type {import('@sveltejs/kit').Config} */
const config = {
    preprocess: [
        preprocess({
            postcss: true,
        }),
    ],
    package: {
        source: "src",
        emitTypes: true,
    },
};

export default config;
