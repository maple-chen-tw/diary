import { createRouter, createWebHistory } from 'vue-router';
import Login from '@/views/Login.vue';
import Register from '@/views/Register.vue';
import Profile from '@/views/Profile.vue';

const routes = [
  { path: '/login', 
    name: 'Login',
    component: Login },
  {
    path: '/register',
    name: 'Register',
    component: Register,
  },
  {
    path: '/profile',
    name: 'Profile',
    component: Profile,
    beforeEnter: (to, from, next) => {
      const token = localStorage.getItem('jwt');
      if (!token) {
        next('/login');  // 如果沒有 JWT，重定向到 /login
      } else {
        next();  // 如果有 token，繼續進入 /profile
      }
    }
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
