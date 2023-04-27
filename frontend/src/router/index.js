import { createRouter, createWebHashHistory } from "vue-router";

const routes = [
    {
        path: '/',
        name: '/',
        component: import('../layouts/index.vue'),
        meta: {
            title: '首页',
        },
        children: [
            {
                path: '/',
                name: '/',
                component: import('../pages/index.vue'),
                meta: {
                    title: '首页',
                }
            },
            {
                path: '/demo',
                name: '/demo',
                component: import('../pages/demo.vue'),
                meta: {
                    title: 'demo',
                }
            },
        ]
    },
    {
        path: '/:pathMatch(.*)*',
        name: 'NotFound',
        component: import('@/pages/404.vue'),
        meta: {
            title: '404',
        }
    }
]

export const router = createRouter({
    history: createWebHashHistory(),
    routes,
})
