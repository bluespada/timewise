import type { UserConfig } from 'vite';
import runtimeDev from './runtime-dev';
import react from '@vitejs/plugin-react';

export default {
    plugins: [
        react(),
        runtimeDev({
            main: "./src/main.tsx"
        }),
    ]
} satisfies UserConfig;