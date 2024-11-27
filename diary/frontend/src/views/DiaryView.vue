<template>
    <div class="diary-view-container">
      <div v-if="diary">
        <h1>{{ diary.question }}</h1>
        <p class="date">{{ formatDate(diary.create_at) }}</p>
        <div class="diary-content">
          <p>{{ diary.content }}</p>
        </div>
  
        <div class="diary-actions">
          <button @click="editDiary" class="btn btn-success">Edit</button>
          <button @click="deleteDiary" class="btn btn-danger">Delete</button>
        </div>
      </div>
      <div v-else>
        <p>Loading...</p>
      </div>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        diary: null, // Object to hold the diary data
      };
    },
    methods: {
      // Fetch a single diary entry by ID
      async fetchDiary(diary_id) {
        try {
          const response = await axios.get(`/diary/${diary_id}`);
          this.diary = response.data;
        } catch (error) {
          console.error('Error fetching diary:', error);
        }
      },
  
      // Redirect to the diary edit page
      editDiary() {
        this.$router.push(`/diary/edit/${this.diary.diary_id}`);
      },
  
      // Delete the diary entry
      async deleteDiary() {
        if (confirm('Are you sure you want to delete this diary?')) {
          try {
            await axios.delete(`/diary/${this.diary.diary_id}`);
            this.$router.push('/diary'); // Navigate back to the diary list after deletion
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
      }

    },
    mounted() {
      const diary_id = this.$route.params.id; // Get the diary ID from route params
      this.fetchDiary(diary_id); // Fetch the diary details when the component is mounted
    },
  };
  </script>
  
  <style scoped>
  .diary-view-container {
    max-width: 800px;
    margin: 50px auto;
    padding: 20px;
    background-color: #f9f9f9;
    border-radius: 10px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  }
  
  h1 {
    font-size: 24px;
    color: #333;
  }
  
  .date {
    font-size: 14px;
    color: #999;
    margin-bottom: 10px;
  }
  
  .diary-content {
    font-size: 16px;
    color: #555;
    margin-bottom: 20px;
  }
  
  .diary-actions {
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
  
  .btn-danger {
    background-color: #dc3545;
    color: white;
  }
  
  button:hover {
    opacity: 0.9;
  }
  
  button:focus {
    outline: none;
  }
  </style>
  