<template>
  <div>
    <h1>Diary List</h1>
    <button @click="createDiary" class="btn btn-primary">Add New Diary</button>

    <div v-if="diaries.length === 0">No diaries found.</div>

    <div v-else>
      <div v-for="diary in diaries" :key="diary.DiaryID" class="diary-item">
        <button 
          class="btn btn-info diary-btn" 
          @click="editDiary(diary.diary_id)">
          <h3>{{ diary.question }}</h3>
          <p>{{ diary.content }}</p>
        </button>
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
      diaries: []
    };
  },
  methods: {
    // Fetch all diaries
    async fetchDiaries() {
      try {
        const response = await axios.get('/diary');
        console.log(response.data);
        this.diaries = response.data;
      } catch (error) {
        console.error('Error fetching diaries:', error);
      }
    },

    // Edit diary: Redirect to the edit page with the diary ID
    editDiary(diary_id) {
      this.$router.push(`/diary/edit/${diary_id}`);
    },

    // Create new diary: Redirect to the new diary page
    createDiary() {
      this.$router.push('/diary/new');
    }
  },
  mounted() {
    this.fetchDiaries();
  }
};
</script>

<style scoped>
/* Custom styles */
.diary-item {
  margin-bottom: 10px;
}

.diary-btn {
  display: block;
  width: 100%;
  padding: 15px;
  text-align: left;
  background-color: #f8f9fa;
  border: 1px solid #ddd;
  border-radius: 8px;
  cursor: pointer;
  transition: background-color 0.3s;
  margin-bottom: 10px;
}

.diary-btn:hover {
  background-color: #e2e6ea;
}

.diary-btn h3 {
  margin: 0;
  font-size: 18px;
  font-weight: bold;
}

.diary-btn p {
  margin: 5px 0 0;
  font-size: 14px;
}
</style>
