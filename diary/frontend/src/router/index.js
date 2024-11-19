import { createRouter, createWebHistory } from 'vue-router';
import Login from '@/views/Login.vue';
import Register from '@/views/Register.vue';
import Profile from '@/views/Profile.vue';
import DiaryList from '@/views/DiaryList.vue';
import DiaryForm from '@/views/DiaryForm.vue';


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
        next('/login');
      } else {
        next();
      }
    }
  },
  {
    path: '/diary',
    name: 'DiaryList',
    component: DiaryList,
    beforeEnter: (to, from, next) => {
      const token = localStorage.getItem('jwt');
      if (!token) {
        next('/login');
      } else {
        next();
      }
    }
  },
  {
    path: '/diary/new',
    name: 'CreateDiary',
    component: DiaryForm,
    beforeEnter: (to, from, next) => {
      const token = localStorage.getItem('jwt');
      if (!token) {
        next('/login');
      } else {
        next();
      }
    }
  },
  {
    path: '/diary/edit/:id',
    name: 'EditDiary',
    component: DiaryForm,
    beforeEnter: (to, from, next) => {
      const token = localStorage.getItem('jwt');
      if (!token) {
        next('/login');
      } else {
        next();
      }
    }
  },
  {
    path: '/diary/:id',
    name: 'ViewDiary',
    component: DiaryForm,  // You might want to create a separate ViewDiary component if necessary
    beforeEnter: (to, from, next) => {
      const token = localStorage.getItem('jwt');
      if (!token) {
        next('/login');
      } else {
        next();
      }
    }
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
