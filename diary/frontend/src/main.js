import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import axios from 'axios';

axios.defaults.baseURL = 'http://localhost:8080';
axios.interceptors.request.use((config) => {
  const token = localStorage.getItem('jwt');
  if (token) {
    config.headers['Authorization'] = `Bearer ${token}`;
  }
  return config;
});

axios.interceptors.response.use(
    (response) => response,
    (error) => {
      console.error('API Request Error:', error.response || error);
      return Promise.reject(error);
    }
  );
  

createApp(App).use(router).mount('#app');
