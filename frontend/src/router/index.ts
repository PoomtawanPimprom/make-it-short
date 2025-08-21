import { createRouter, createWebHistory } from 'vue-router'
import Hero from '@/views/Hero.vue'
import ShortURL from '@/views/ShortURL.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: "/", name: "home", component: Hero },
    { path: "/shortURL", name: "shortURL", component: ShortURL },
  ],
})

export default router
