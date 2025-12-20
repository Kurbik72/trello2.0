import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/boards',
      component: () => import('@/pages/boards/boards.vue'),
    },
    {
      path: '/auth',
      component: () => import('@/pages/auth/auth.vue'),
      redirect: '/auth/login',
      children: [
        {
          path: 'login',
          component: () => import('@/pages/auth/login.vue'),
        },
        {
          path: 'register',
          component: () => import('@/pages/auth/register.vue'),
        },
      ],
    },
  ],
})

export default router
