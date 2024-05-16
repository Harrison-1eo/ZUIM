import { createRouter, createWebHashHistory } from 'vue-router'
// import UserLogin from "@/components/UserLogin.vue";

const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/about',
    name: 'about',
    component: () => import(/* webpackChunkName: "about" */ '../views/IMView.vue')
  },
  {
    path: '/login',
    name: 'login',
    component: () => import(/* webpackChunkName: "login" */ '../views/LoginView.vue')
  },
  {
    path: '/im',
    name: 'im',
    component: () => import(/* webpackChunkName: "im" */ '../views/IMView.vue'),
    children: [
      {
        path: 'home',
        name: 'home',
        component: () =>
          import(/* webpackChunkName: "chat" */ '../components/IM/OtherFunc/home.vue')
      },
      {
        path: 'info',
        name: 'info',
        component: () =>
          import(/* webpackChunkName: "chat" */ '../components/IM/OtherFunc/UserInfo.vue')
      },
      {
        path: 'updateInfo',
        name: 'updateInfo',
        component: () =>
          import(/* webpackChunkName: "chat" */ '../components/IM/OtherFunc/UpdateUserInfo.vue')
      },
      {
        path: 'password',
        name: 'password',
        component: () =>
          import(/* webpackChunkName: "contacts" */ '../components/IM/OtherFunc/UpdatePassword.vue')
      },
      {
        path: 'friend',
        name: 'friend',
        component: () =>
          import(/* webpackChunkName: "discover" */ '../components/IM/OtherFunc/friend.vue')
      },
      {
        path: 'chat',
        name: 'chat',
        component: () => import(/* webpackChunkName: "me" */ '../components/IM/Room/RoomList.vue')
      },
      {
        path: 'setting',
        name: 'setting',
        component: () =>
          import(/* webpackChunkName: "setting" */ '../components/IM/OtherFunc/setting.vue')
      }
    ]
  },
  {
    path: '/:catchAll(.*)',
    name: 'not-found',
    component: () => import(/* webpackChunkName: "not-found" */ '../views/NotFoundView.vue')
  }
]

// const router = createRouter({
//   history: createWebHashHistory(process.env.BASE_URL),
//   routes
// })

const router = createRouter({
  history: createWebHashHistory(import.meta.env.VITE_BASE_URL),
  routes
})

export default router