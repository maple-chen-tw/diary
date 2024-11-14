<template>
    <div class="register-container">
      <div class="register-box">
        <h2>Register</h2>
        <form @submit.prevent="register">
          <div class="input-group">
            <input v-model="registerUsername" type="text" placeholder="Username" required />
          </div>
          <div class="input-group">
            <input v-model="registerEmail" type="email" placeholder="Email" required />
          </div>
          <div class="input-group">
            <input v-model="registerPassword" type="password" placeholder="Password" required />
          </div>
          <div class="input-group">
            <input v-model="registerConfirmPassword" type="password" placeholder="Confirm Password" required />
          </div>
          <button type="submit" class="register-btn">Register</button>
        </form>
        <div class="login-prompt">
          <p>Already have an account?</p>
          <button @click="goToLogin" class="login-btn">Login</button>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        registerUsername: '',
        registerEmail: '',
        registerPassword: '',
        registerConfirmPassword: '',
      };
    },
    methods: {
      async register() {
        // 基本的密碼匹配檢查
        if (this.registerPassword !== this.registerConfirmPassword) {
          alert("Passwords don't match!");
          return;
        }
  
        try {
          const response = await axios.post('/register', {
            username: this.registerUsername,
            password: this.registerPassword,
            email: this.registerEmail,
          });
          if (response.data.success) {
            alert('Registration successful! You can now log in.');
            this.$router.push('/login');
          } else {
            alert('Registration failed. Please try again.');
          }
        } catch (error) {
          alert('An error occurred during registration. Please try again later.');
          console.error('Registration failed:', error);
        }
      },
      goToLogin() {
        this.$router.push('/login');
      }
    },
  };
  </script>
  
  <style scoped>
  /* 整體頁面背景 */
  .register-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background: #f4f7fc;
  }
  
  /* 註冊框樣式 */
  .register-box {
    background: white;
    padding: 30px;
    border-radius: 8px;
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
    width: 100%;
    max-width: 400px;
    text-align: center;
  }
  
  /* 標題樣式 */
  .register-box h2 {
    font-size: 24px;
    margin-bottom: 20px;
    color: #333;
  }
  
  /* 輸入框樣式 */
  .input-group {
    margin-bottom: 15px;
  }
  
  .input-group input {
    width: 100%;
    padding: 10px;
    border: 1px solid #ddd;
    border-radius: 4px;
    font-size: 16px;
    outline: none;
    transition: border 0.3s ease;
  }
  
  .input-group input:focus {
    border-color: #0056b3;
  }
  
  /* 註冊按鈕樣式 */
  .register-btn {
    width: 100%;
    padding: 10px;
    background-color: #28a745;
    color: white;
    font-size: 16px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.3s;
  }
  
  .register-btn:hover {
    background-color: #218838;
  }
  
  /* 登入提示文字和按鈕 */
  .login-prompt {
    margin-top: 20px;
    font-size: 14px;
  }
  
  .login-btn {
    margin-top: 5px;
    padding: 8px 16px;
    background-color: #0056b3;
    color: white;
    font-size: 14px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.3s;
  }
  
  .login-btn:hover {
    background-color: #00408b;
  }
  </style>
  