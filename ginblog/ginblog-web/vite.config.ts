import {defineConfig} from "vite";
import vue from "@vitejs/plugin-vue";
import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import {ElementPlusResolver} from "unplugin-vue-components/resolvers";
import * as path from 'path'

const resolve = path.resolve;
export default defineConfig({
    base: "/slim-admin/",
    plugins: [
        vue({reactivityTransform: true}),
        AutoImport({
            resolvers: [ElementPlusResolver()],
            imports: [
                'vue',
                "vue-router",
                "pinia",
                'vue/macros',
                '@vueuse/head',
                '@vueuse/core',
            ], dts: 'src/auto-imports.d.ts',
        }),

        Components({
            resolvers: [ElementPlusResolver()], dts: 'src/components.d.ts',
        }),
    ], resolve: {
        //导入时想要省略的扩展名列表。注意，不 建议忽略自定义导入类型的扩展名（例如：.vue），因为它会干扰 IDE 和类型支持。
        alias: [
            {find: "@", replacement: resolve(__dirname, "./src")},

        ],
    },
});
