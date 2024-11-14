<template>
  <div class="login-container">
    <div class="login-box">
      <h2>Login</h2>
      <form @submit.prevent="login">
        <div class="input-group">
          <input v-model="username" type="text" placeholder="Username" required />
        </div>
        <div class="input-group">
          <input v-model="password" type="password" placeholder="Password" required />
        </div>
        <button type="submit" class="login-btn">Login</button>
      </form>
      <div class="register-prompt">
        <p>Don't have an account?</p>
        <button @click="goToRegister" class="register-btn">Register</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      username: '',
      password: '',
    };
  },
  methods: {
    async login() {
      try {
        const response = await axios.post('/login', {
          username: this.username,
          password: this.password,
        });
        const token = response.data.token;
        localStorage.setItem('jwt', token);
        console.log('Login successful, token stored');
        this.$router.push('/profile');
      } catch (error) {
        if (error.response && error.response.status === 401) {
          alert('Invalid credentials');
        } else {
          alert('An error occurred. Please try again later');
        }
        console.error('Login failed:', error);
      }
    },
    goToRegister() {
      this.$router.push('/register');
    }
  },
};
</script>

<style scoped>
/* 整體頁面背景 */
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: #f4f7fc;
}

/* 登入框樣式 */
.login-box {
  background: white;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 400px;
  text-align: center;
}

/* 標題樣式 */
.login-box h2 {
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

/* 登入按鈕樣式 */
.login-btn {
  width: 100%;
  padding: 10px;
  background-color: #0056b3;
  color: white;
  font-size: 16px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.login-btn:hover {
  background-color: #00408b;
}

/* 註冊提示文字和按鈕 */
.register-prompt {
  margin-top: 20px;
  font-size: 14px;
}

.register-btn {
  margin-top: 5px;
  padding: 8px 16px;
  background-color: #28a745;
  color: white;
  font-size: 14px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.register-btn:hover {
  background-color: #218838;
}
</style>
