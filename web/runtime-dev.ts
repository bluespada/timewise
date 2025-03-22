import { Plugin as VitePlugin, PluginOption, build } from 'vite';

export default function RuntimeDev(input: { [key: string]: string }) : VitePlugin {
    return {
        name: "runtime-dev",
        config(config, env) {
            config.base = "/assets",
            config.build = {
                lib: {
                    entry: input.main,
                    name: "timewise",
                    fileName: (format) => `timewise.${format}.js`,
                },
                rollupOptions: {
                    input: {
                        ...input,
                    },
                    output: {
                        entryFileNames: 'js/[name].js',
                        assetFileNames: 'css/[name].[ext]'
                    },
                    ...config.build?.rollupOptions,
                },
                sourcemap: env.mode === "development",
                manifest: env.mode === "production",
                minify: env.mode === "production",
                cssMinify: env.mode === "production",
                emptyOutDir: true,
                ...config.build,
            }
            return config;
        },
    }
}