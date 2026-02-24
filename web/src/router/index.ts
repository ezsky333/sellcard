/**
 * router/index.ts
 *
 * Manual routes for ./src/pages/*.vue
 */

// Composables
import { createRouter, createWebHistory } from 'vue-router'
import Index from '@/pages/index.vue'
import Backend from '@/pages/backend.vue'
import Login from '@/pages/login.vue'
import { useAppStore } from '@/stores/app'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: Index,
    },
    {
      path: '/backend',
      component: Backend,
    },
    {
      path: '/login',
      component: Login,
    },
  ],
})

router.beforeEach((to) => {
  const store = useAppStore()
  if (to.path === '/backend' && !store.token) {
    store.setRedirect(to.fullPath)
    return { path: '/login' }
  }
})

export default router
