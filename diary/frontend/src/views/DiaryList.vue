<template>
  <div class="diary-list-container">
    <!-- Header Section -->
    <div class="header">
      <div class="logo">
        <h1>QuestionDiary</h1>
      </div>
      <div class="auth-buttons">
        <button @click="createDiary" class="auth-btn add-diary-btn">Add New Diary</button>
      </div>
    </div>

    <!-- Main Content -->
    <div class="main-content">

      <!-- Diary List -->
      <div v-if="diaries.length === 0" class="no-diaries-message">
        No diaries found.
      </div>

      <div v-else>
        <div v-for="diary in diaries" :key="diary.diary_id" class="diary-item">
          <button 
            class="btn diary-btn" 
            @click="viewDiary(diary.diary_id)">
            <h3 class="diary-question">{{formatDate(diary.create_at)}} - {{ diary.question }}</h3>
            <p class="diary-content">{{ diary.content }}</p>
          </button>
        </div>
      </div>
    </div>

    <!-- Create Diary Modal -->
    <div v-if="showCreateDiaryModal" class="modal-overlay">
      <div class="modal">
        <button @click="showCreateDiaryModal = false" class="close-btn">X</button>
        <h2>Create New Diary</h2>
        <form @submit.prevent="createDiaryEntry">
          <div class="input-group">
            <input v-model="diaryQuestion" type="text" placeholder="Enter your question" required />
          </div>
          <div class="input-group">
            <textarea v-model="diaryContent" placeholder="Write your thoughts..." required></textarea>
          </div>
          <button type="submit" class="create-diary-btn">Create Diary</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import { useRouter } from 'vue-router';

export default {
  data() {
    return {
      diaries: [],
      showCreateDiaryModal: false,
      diaryQuestion: '',
      diaryContent: '',
    };
  },
  methods: {
    // Fetch all diaries
    async fetchDiaries() {
      try {
        const response = await axios.get('/diary');
        this.diaries = response.data;
      } catch (error) {
        console.error('Error fetching diaries:', error);
      }
    },

    // Create a new diary entry: Display the modal
    createDiary() {
      this.showCreateDiaryModal = true;
    },

    // Create diary entry action: Save to the server
    async createDiaryEntry() {
      try {
        const response = await axios.post('/diary', {
          question: this.diaryQuestion,
          content: this.diaryContent,
        });
        this.diaries.push(response.data);
        this.showCreateDiaryModal = false;
        this.diaryQuestion = '';
        this.diaryContent = '';
        alert('Diary entry created!');
      } catch (error) {
        console.error('Error creating diary:', error);
        alert('An error occurred. Please try again.');
      }
    },

    // View a specific diary entry
    viewDiary(diary_id) {
      this.$router.push(`/diary/${diary_id}`);
    },
    formatDate(dateString) {
        if (!dateString) return "Invaild Date";
          
        // Create a new Date object using the ISO 8601 format string
        const date = new Date(dateString);
          
        // Check if the Date object is valid
        if (isNaN(date.getTime())) {
          console.error("Invalid date:", dateString);
          return 'Invalid Date'; // Handle invalid date format
        }
      
        // Return a formatted string in a readable format (e.g., 'MM/DD/YYYY')
        return date.toLocaleDateString(); 
      }
  },
  mounted() {
    this.fetchDiaries();
  }
};
</script>

<style scoped>
/* 整體頁面背景 */
.diary-list-container {
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

/* Create New Diary Button */
.add-diary-btn {
  background-color: #4caf50;
  color: white;
}

.add-diary-btn:hover {
  background-color: #45a049;
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

.no-diaries-message {
  font-size: 18px;
  color: #666;
  margin-top: 20px;
}

/* 日記項目按鈕 */
.diary-item {
  margin-bottom: 15px;
}

.diary-btn {
  display: block;
  width: 100%;
  padding: 15px;
  text-align: left;
  background-color: #ffffff;
  border: 1px solid #ddd;
  border-radius: 8px;
  cursor: pointer;
  transition: background-color 0.3s ease, transform 0.2s ease;
}

.diary-btn:hover {
  background-color: #f1f1f1;
  transform: translateY(-5px);
}

.diary-question {
  font-size: 18px;
  font-weight: bold;
  color: #333;
  margin: 0;
}

.diary-content {
  font-size: 14px;
  color: #666;
  margin-top: 5px;
}

/* 創建日記模態框 */
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

/* Form Inputs */
.input-group {
  margin-bottom: 20px;
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

.create-diary-btn {
  padding: 12px 25px;
  font-size: 16px;
  font-weight: bold;
  background-color: #4caf50;
  color: white;
  border: none;
  border-radius: 50px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.create-diary-btn:hover {
  background-color: #45a049;
}

/* Close Button */
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
</style>
