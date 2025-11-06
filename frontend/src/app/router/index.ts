import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/boards',
      component: () => import('@/pages/boards/boards.vue'),
    },
  ],
})

export default router
