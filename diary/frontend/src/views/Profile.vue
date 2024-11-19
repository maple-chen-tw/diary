<template>
    <div class="profile">
      <h1>Welcome to your profile!</h1>
      <p v-if="loading">Loading...</p>
      <p v-if="error" class="error">{{ error }}</p>
      <p v-if="username">Username: {{ username }}</p>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        username: '',  // 用來存儲從 API 獲取的 username
        error: null,   // 用來存儲錯誤信息
        loading: true, // 用來顯示加載狀態
      };
    },
    created() {
      this.fetchProfile();
    },
    methods: {
      async fetchProfile() {
        try {
          // 從 localStorage 獲取 JWT token
          const token = localStorage.getItem('jwt');
          
          if (!token) {
            this.error = 'No token found';
            this.loading = false;
            return;
          }
  
          // 發送 GET 請求到後端 API，並將 JWT token 放入 Authorization 頭部
          const response = await axios.get('/profile', {
            headers: {
              Authorization: `Bearer ${token}`,  // 傳送 JWT token
            },
          });
  
          // 當成功獲取資料後，將資料賦值給 username
          this.username = response.data.username;
          this.loading = false;
        } catch (err) {
          // 處理錯誤
          this.error = 'Failed to fetch profile: ' + (err.response?.data?.message || err.message);
          this.loading = false;
        }
      },
    },
  };
  </script>
  
  <style scoped>
  .profile {
    text-align: center;
  }
  
  .error {
    color: red;
  }
  
  .loading {
    font-style: italic;
  }
  </style>
  