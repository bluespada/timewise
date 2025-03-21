import type { UserConfig } from 'vite';
import runtimeDev from './runtime-dev';

export default {
    plugins: [
        runtimeDev({
            main: "./src/main.ts"
        }),
    ]
} satisfies UserConfig;