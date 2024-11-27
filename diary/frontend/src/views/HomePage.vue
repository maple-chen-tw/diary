<template>
  <div class="login-container">
    <div class="header">
      <div class="logo">
        <h1>QuestionDiary</h1>
      </div>
      <div class="auth-buttons">
        <!--  Sign In  -->
        <button @click="showLoginModal = true" class="auth-btn login-btn">Sign In</button>
        <!--  Register -->
        <button @click="showRegisterModal = true" class="auth-btn register-btn">Register</button>
      </div>
    </div>


    <div class="main-content">
      <p class="site-info">
        QuestionDiary is a personal space to reflect, write, and grow. It's your diary for capturing thoughts, feelings, and experiences.
      </p>

      <!-- Start Diary  -->
      <button @click="showRegisterModal = true" class="start-diary-btn">Start Diary</button>
    </div>

    <!-- Register -->
    <div v-if="showRegisterModal" class="modal-overlay">
      <div class="modal">
        <button @click="showRegisterModal = false" class="close-btn">X</button>
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
            <input v-model="confirmPassword" type="password" placeholder="Confirm Password" required />
          </div>
          <button type="submit" class="register-btn">Create Account</button>
          
        </form>
      </div>
    </div>


    <!-- login -->
    <div v-if="showLoginModal" class="modal-overlay">
      <div class="modal">
        <button @click="showLoginModal = false" class="close-btn">X</button>
        <h2>Sign In</h2>
        <form @submit.prevent="login">
          <div class="input-group">
            <input v-model="loginUsername" type="text" placeholder="Username" required />
          </div>
          <div class="input-group">
            <input v-model="loginPassword" type="password" placeholder="Password" required />
          </div>
          <button type="submit" class="login-btn">Log In</button>
        </form>
      </div>
    </div>

  </div>
</template>

<script>
import axios from 'axios';
export default {
  data() {
    return {
      showRegisterModal: false,
      showLoginModal: false,
      
      // register
      registerUsername: '',
      registerEmail: '',
      registerPassword: '',
      confirmPassword: '', 
      
      // login
      loginUsername: '',
      loginPassword: '',
      

    };
  },
  methods: {
    async login() {
      try {
        const response = await axios.post('/login', {
          username: this.loginUsername,
          password: this.loginPassword,
        });
        const token = response.data.token;
        localStorage.setItem('jwt', token);
        console.log('Login successful, token stored');
        this.$router.push('/diary');
      } catch (error) {
        if (error.response && error.response.status === 401) {
          alert('Invalid credentials');
        } else {
          alert('An error occurred. Please try again later');
        }
        console.error('Login failed:', error);
      }
    },

    // register
    register() {
    if (this.registerPassword !== this.confirmPassword) {
      alert("Passwords do not match!");
      return;
    }
    
    console.log('Registering with:', this.registerUsername, this.registerEmail, this.registerPassword);
    this.showRegisterModal = false;
    alert('Registration successful!');
    }

  }
};
</script>

<style scoped>
/* 整體頁面背景 */
.login-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  background: #f7f7f7;
  font-family: 'Merriweather', sans-serif;
  text-align: center;
  min-height: 100vh;
}

/* 頂部區域 */
.header {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px 30px;
  background-color: #ffffff;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  z-index: 10;
}

.logo h1 {
  font-size: 36px;
  margin: 0;
  color: #000;
}

.auth-buttons {
  display: flex;
  gap: 15px;
}

.auth-btn {
  padding: 12px 25px;
  font-size: 16px;
  font-weight: bold;
  border: none;
  border-radius: 50px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.login-btn {
  background-color: #4caf50;
  color: white;
}

.login-btn:hover {
  background-color: #45a049;
}

.register-btn {
  background-color: #ffffff;
  color: #4caf50;
  border: 2px solid #4caf50;
}

.register-btn:hover {
  background-color: #f1f1f1;
}

/* 主體區域 */
.main-content {
  padding: 40px 20px;
  width: 100%;
  max-width: 600px;
  background-color: transparent;
  margin-top: 100px;
}

.site-info {
  font-size: 18px;
  font-weight: 300;
  color: #666;
  margin-bottom: 30px;
}

/* 開始日記按鈕 */
.start-diary-btn {
  padding: 15px 30px;
  font-size: 18px;
  background-color: #4caf50;
  color: white;
  border: none;
  border-radius: 50px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.start-diary-btn:hover {
  background-color: #45a049;
}

/* 註冊、登入和創建日記的模態框 */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal {
  background-color: #ffffff;
  padding: 30px;
  border-radius: 10px;
  width: 400px;
  max-width: 400px;
  position: relative;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  text-align: center;
}

.modal h2 {
  font-size: 24px;
  margin-bottom: 20px;
}


.input-group {
  margin-bottom: 20px;
}
.close-btn {
  position: absolute;
  top: 10px;
  right: 10px;
  font-size: 24px;
  font-weight: bold;
  background: none;
  border: none;
  color: #666;
  cursor: pointer;
}

.close-btn:hover {
  color: #333;
}

.input-group input,
.input-group textarea {
  width: 100%;
  padding: 12px;
  border: 1px solid #ccc;
  border-radius: 8px;
  font-size: 16px;
  outline: none;
}

.input-group input:focus,
.input-group textarea:focus {
  border-color: #4caf50;
}

</style>
