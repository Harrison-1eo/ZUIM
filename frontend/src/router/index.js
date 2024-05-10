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
            import ( /* webpackChunkName: "about" */ '../views/IMView.vue')
    },
    {
        path: '/login',
        name: 'login',
        component: () =>
            import ( /* webpackChunkName: "login" */ '../views/LoginView.vue')
    },
    {
        path: '/im',
        name: 'im',
        component: () =>
            import ( /* webpackChunkName: "im" */ '../views/IMView.vue'),
        children: [
            {
                path: 'home',
                name: 'home',
                component: () =>
                    import ( /* webpackChunkName: "chat" */ '../components/IM/home.vue')
            },
            {
                path: 'info',
                name: 'info',
                component: () =>
                    import ( /* webpackChunkName: "chat" */ '../components/IM/info.vue')
            },
            {
                path: 'password',
                name: 'password',
                component: () =>
                    import ( /* webpackChunkName: "contacts" */ '../components/IM/password.vue')
            },
            {
                path: 'friend',
                name: 'friend',
                component: () =>
                    import ( /* webpackChunkName: "discover" */ '../components/IM/friend.vue')
            },
            {
                path: 'chat',
                name: 'chat',
                component: () =>
                    import ( /* webpackChunkName: "me" */ '../components/IM/chat.vue')
            },
            {
                path: 'setting',
                name: 'setting',
                component: () =>
                    import ( /* webpackChunkName: "setting" */ '../components/IM/setting.vue')
            }
        ]
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
            import ( /* webpackChunkName: "not-found" */ '../views/NotFoundView.vue')
    }
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router