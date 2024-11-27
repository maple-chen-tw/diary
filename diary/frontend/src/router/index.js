import { createRouter, createWebHistory } from 'vue-router';
import HomePage from '@/views/HomePage.vue';
import Register from '@/views/Register.vue';
import Profile from '@/views/Profile.vue';
import DiaryList from '@/views/DiaryList.vue';
import DiaryEdit from '@/views/DiaryEdit.vue';
import DiaryView from '@/views/DiaryView.vue';

const routes = [
  { path: '/', 
    name: 'HomePage',
    component: HomePage },

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
    component: DiaryEdit,
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
    component: DiaryEdit,
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
    component: DiaryView,
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
    path: '/question/random',
    name: 'RandomQuestion',
    component: DiaryEdit,
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
