import { createRouter, createWebHistory } from 'vue-router'
// import UserLogin from "@/components/UserLogin.vue";

const routes = [
    {
        path: '/',
        name: 'home',
        component: () =>
            import ( /* webpackChunkName: "home" */ '../views/HomeView.vue')
    },
    {
        path: '/about',
        name: 'about',
        component: () =>
            import ( /* webpackChunkName: "about" */ '../views/UserView.vue')
    },
    {
        path: '/login',
        name: 'login',
        component: () =>
            import ( /* webpackChunkName: "login" */ '../views/LoginView.vue')
    },
    {
        path: '/hello',
        name: 'hello',
        component: () =>
            import ( /* webpackChunkName: "hello" */ '../components/HelloWorld.vue')
    },
    {
        path: "/:catchAll(.*)",
        name: 'not-found',
        component: () =>
            import ( /* webpackChunkName: "not-found" */ '../views/NotFound.vue')
    }
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router