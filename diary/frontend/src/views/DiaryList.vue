<template>
  <div class="diary-list-container">
    <!-- Header Section -->
    <div class="header">
      <div class="logo">
        <h1>QuestionDiary</h1>
      </div>
      <div class="auth-buttons">
        <button @click="createDiary" class="auth-btn add-diary-btn">Add New Diary</button>
        <button @click="logout" class="auth-btn logout-btn">Logout</button>
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
            <p class="">{{ formatTime(diary.create_at) }}</p>
          </button>
          <button 
            class="btn delete-btn" 
            @click="deleteDiary(diary.diary_id)">
            <i class="fa fa-trash"></i> <!-- 使用 Font Awesome 垃圾桶圖標 -->
          </button>
        </div>
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
      diaryQuestion: '',
      diaryContent: '',
    };
  },
  methods: {
    logout() {
      if (confirm('Are you sure you want to logout?')) {
        localStorage.removeItem('jwt');
        this.$router.push('/');
      }
    },

    // Fetch all diaries
    async fetchDiaries() {
      try {
        const response = await axios.get('/diary');
        this.diaries = response.data;
      } catch (error) {
        console.error('Error fetching diaries:', error);
      }
    },


    createDiary() {
      
      axios.get('/question/random')
      .then(response => {
        const randomQuestionID = response.data.question.question_id;
        const randomQuestion = response.data.question.question;
        const newDiary = {
          question_id: randomQuestionID,
          question: randomQuestion,
          content: '',
        };
        console.log(newDiary)
        axios.post('/diary/new', newDiary, {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('jwt')}`
          }
        })
        .then(response => {
          console.log("Create Diary Response:", response.data); 
          if (response.data.diary_id !== 0) {
            this.diaries.push(response.data);
            console.log("New Diary Added:", this.diaries);
            this.isCreatingNewDiary = false;
          } else {
            console.error("Diary creation failed. Invalid diary_id:", response.data);
          }
        })
        .catch(error => {
          console.error("Failed to create diary:", error);
        });
      })
      .catch(error => {
        console.error("Failed to fetch random question:", error);
      });
      
    },

    // View a specific diary entry
    viewDiary(diary_id) {
      this.$router.push(`/diary/${diary_id}`);
    },

    async deleteDiary(diary_id) {
      if (confirm('Are you sure you want to delete this diary?')) {
        try {
          // Make the DELETE request to the server
          await axios.delete(`/diary/${diary_id}`);
        
          // Remove the deleted diary from the local diaries list
          this.diaries = this.diaries.filter(diary => diary.diary_id !== diary_id);
        
          // Optionally, navigate back to the diary list or show a success message
          this.$router.push('/diary');
        } catch (error) {
          console.error('Error deleting diary:', error);
        }
      }
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
      },
    formatTime(dateString){
      if (!dateString) return "Invaild Date";
      const date = new Date(dateString);
      const options = {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit',
        hour12: true, // Enable 12-hour format (AM/PM)
    };

    // Return a formatted string in a readable format (e.g., 'MM/DD/YYYY 下午12:59')
    return date.toLocaleString('zh-TW', options).replace(/上午|下午/g, (match) => match === '上午' ? 'AM' : 'PM');
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

.diary-item {
  position: relative; /* 使垃圾桶按鈕能夠相對定位 */
}

.delete-btn {
  position: absolute;
  bottom: 10px; /* 距離底部 10px */
  right: 10px;  /* 距離右邊 10px */
  background-color: white;  /* 背景顏色 */
  color: #808080; /* 淺灰色文字或圖標 */
  border: none;
  border-radius: 5px;
  padding: 10px;
  cursor: pointer;
  transition: background-color 0.3s, color 0.3s;
}

.delete-btn:hover {
  background-color: #f5f5f5; /* 當鼠標懸停時，背景顏色變為輕微灰色 */
  color: #333; /* 當鼠標懸停時，顏色變為深灰 */
}

.delete-btn i {
  font-size: 16px;
}


</style>
