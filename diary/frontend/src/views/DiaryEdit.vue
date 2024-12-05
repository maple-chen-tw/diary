<template>
  <div class="diary-form-container">
    <h1>{{ isEdit ? 'Edit Diary' : 'Create New Diary' }}</h1>

    <button @click="showModal = true" class="btn btn-info btn-sm ml-2">Change a Question</button>
    
    <!-- 可選問題的浮動視窗 -->
    <div v-if="showModal" class="modal-overlay">
      <div class="modal-content">
        <h2>Question List
          <button @click="closeModal" class="btn btn-secondary">X</button>
        </h2>
        <ul>
          <li v-for="(question, index) in availableQuestions" :key="index">
            <button @click="selectQuestion(question)" class="btn btn-link">{{ question }}</button>
          </li>
        </ul>
        
      </div>
    </div>

    <form @submit.prevent="saveDiary" class="diary-form">
      <div class="form-group">
        <label for="question">Question</label>
        <input v-model="diary.question" id="question" type="text" required placeholder="Enter your question here" disabled  />
      </div>

      <div class="form-group">
        <label for="content">Content</label>
        <textarea v-model="diary.content" id="content" rows="5" required placeholder="Write your thoughts here..."></textarea>
      </div>

      <div class="form-actions">
        <button type="submit" class="btn btn-success">{{ isEdit ? 'Save Diary' : 'Create Diary' }}</button>
        <button type="button" @click="cancelEdit" class="btn btn-secondary">Cancel</button>
      </div>
    </form>
  </div>
</template>

<script>
import axios from 'axios';
import { useRoute, useRouter } from 'vue-router';

export default {
  data() {
    return {
      diary: {
        diary_id: null,
        question: '',
        content: '',
        question_id: null,
        UpdateAt: null,
      },
      isEdit: false, // Flag to check if we're in edit mode
      showModal: false,
    };
  },
  methods: {
    // Fetch a single diary by ID
    async fetchDiary(diary_id) {
      console.log('Fetching diary with ID:', diary_id);

      try {
        const response = await axios.get(`/diary/${diary_id}`);
        console.log('Diary data received:', response.data); 
        const diary = response.data;
        this.diary = {
          diary_id: diary.diary_id,
          question: diary.question,
          content: diary.content,
          question_id: diary.question_id,
          UpdateAt: diary.update_at,
        };
        this.isEdit = true;
      } catch (error) {
        console.error('Error fetching diary:', error);
      }
    },

    // Fetch a random question when creating a new diary
    async fetchRandomQuestion() {
      try {
        const response = await axios.get('/question/random'); // This gives a random question
        this.diary.title = response.data.question;
      } catch (error) {
        console.error('Error fetching random question:', error);
      }
    },

    // Save a new or existing diary
    async saveDiary() {
      try {
        if (this.isEdit) {
          // Update existing diary
          await axios.put(`/diary/${this.diary.diary_id}`, this.diary);
        } else {
          // Create new diary
          await axios.post('/diary', this.diary);
        }
        this.$router.push('/diary'); // Redirect to diary list
      } catch (error) {
        console.error('Error saving diary:', error);
      }
    },

    // Cancel editing and go back to the diary list
    cancelEdit() {
      this.$router.push('/diary'); // Navigate back to the diary list page
    },
    closeModal() {
      this.showModal = false;
    },
  },
  computed: {
    // Check if we're in edit mode by checking if an ID exists in the route params
    isEditMode() {
      return this.$route.params.id && this.$route.params.id !== 'new';
    },
  },
  mounted() {
    console.log("Route Params:", this.$route.params);
    const diary_id = this.$route.params.id;
    if (diary_id === 'new') {
    this.fetchRandomQuestion();  
  } else if (diary_id) {
    this.fetchDiary(diary_id);
  }
  },
};
</script>

<style scoped>
.diary-form-container {
  max-width: 600px;
  margin: 50px auto;
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 10px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

h1 {
  text-align: center;
  font-size: 24px;
  margin-bottom: 20px;
  color: #333;
}

.diary-form {
  display: flex;
  flex-direction: column;
}

.form-group {
  margin-bottom: 20px;
}

label {
  font-size: 16px;
  color: #333;
  margin-bottom: 5px;
  display: block;
}

input,
textarea {
  width: 100%;
  padding: 10px;
  font-size: 16px;
  border: 1px solid #ddd;
  border-radius: 8px;
  box-sizing: border-box;
  margin-top: 5px;
}

textarea {
  resize: vertical;
}

.form-actions {
  display: flex;
  justify-content: space-between;
}

button {
  padding: 10px 20px;
  font-size: 16px;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.btn-success {
  background-color: #28a745;
  color: white;
}

.btn-success:hover {
  background-color: #218838;
}

.btn-secondary {
  background-color: #6c757d;
  color: white;
}

.btn-secondary:hover {
  background-color: #5a6268;
}

button:focus {
  outline: none;
}


.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background-color: white;
  padding: 20px;
  border-radius: 5px;
  width: 300px;
  text-align: center;
}

.modal-content h2 {
  margin-bottom: 15px;
}

.modal-content button {
  margin-top: 10px;
}

button {
  cursor: pointer;
}


</style>
