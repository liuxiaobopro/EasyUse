import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import WindiCSS from 'vite-plugin-windicss'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
// import { createStyleImportPlugin, ElementPlusResolve } from 'vite-plugin-style-import';
import path from "path";

export default defineConfig({
    resolve: {
        alias: {
            "@": path.resolve(__dirname, "src"),
            "_C": path.resolve(__dirname, "src/components"),
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
    css: {
        preprocessorOptions: {
            scss: {
                // 自定义的主题色
                additionalData: `@use "@/styles/element/index.scss" as *;`,
            },
        },
    },
    plugins: [
        vue(),
        WindiCSS(),
        //配置自动导入element start
        // createStyleImportPlugin({
        //     resolves: [ElementPlusResolve()],
        //     libs: [
        //         {
        //             libraryName: 'element-plus',
        //             esModule: true,
        //             resolveStyle: (name) => {
        //                 return `element-plus/theme-chalk/${name}.css`
        //             },
        //         },
        //     ]
        // }),
        // 自动引入
        AutoImport({
            resolvers: [ElementPlusResolver({
                // 自动引入修改主题色添加这一行，使用预处理样式，不添加将会导致使用ElMessage，ElNotification等组件时默认的主题色会覆盖自定义的主题色
                importStyle: "sass",
            })],
        }),
        Components({
            resolvers: [ElementPlusResolver({
                // 自动引入修改主题色添加这一行，使用预处理样式
                importStyle: "sass",
            })],
        })
    ],
})
